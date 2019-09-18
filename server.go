package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const StatusEndpoint = "https://webhooks.mongodb-stitch.com/api/client/v2.0/app/ctagsvis-bymac/service/ctagsStatus/incoming_webhook/status"
const ClearDataEndpoint = "https://webhooks.mongodb-stitch.com/api/client/v2.0/app/ctagsvis-bymac/service/ctagsStatus/incoming_webhook/clearData"
const LoadDataEndpoint = "https://webhooks.mongodb-stitch.com/api/client/v2.0/app/ctagsvis-bymac/service/ctagsStatus/incoming_webhook/loadData"

type TagData struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Pattern    string `json:"pattern"`
	Kind       string `json:"kind"`
	Scope      string `json:"scope"`
	Package    string `json:"package"`
	LineNum    int    `json:"lineNum"`
	GitProject string `json:"gp"`
}

func getHtmlTemplateBinData() []byte {
	data, _ := htmlIndexHtmlBytes()
	return data
}

func getOutput(w http.ResponseWriter) {
	tmpl := template.Must(template.New("index").Parse(string(getHtmlTemplateBinData())))
	tmpl.Execute(w, nil)
}

func makeHandleRootFn() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		getOutput(w)
	}
}

func updateStatus(project, status string) {
	fmt.Println(project, status)
	reqb, _ := json.Marshal(map[string]string{
		"project": project,
		"status":  status,
	})
	resp, err := http.Post(StatusEndpoint, "application/json", bytes.NewBuffer(reqb))
	if err != nil {
		fmt.Println("error posting status", err)
	}
	resb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error getting request body", err)
	}
	fmt.Println("updateStatus response", string(resb))
}

func enrichCtags(inputFile, gitproject string) error {
	updateStatus(gitproject, "enriching ctags data")
	tagsFile, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}
	raw := string(tagsFile)
	lines := strings.Split(raw, "\n")
	tags := make([]TagData, 0)
	for j, line := range lines {
		var tag TagData
		json.Unmarshal([]byte(line), &tag)
		if tag.Path == "" {
			continue
		}
		if j%100 == 0 && j > 0 {
			updateStatus(gitproject, "enriching ctags data - processed "+fmt.Sprintf("%d", 100*j/len(lines))+"%")
		}
		// read the first two lines to extract package
		contents, err := ioutil.ReadFile(tag.Path)
		if err != nil {
			return fmt.Errorf("err reading %s: %v", tag.Path, err)
		}
		tag.GitProject = gitproject
		rawCont := string(contents)
		contLines := strings.Split(rawCont, "\n")
	INNER:
		for i, contLine := range contLines {
			if i < 2 {
				if strings.Contains(contLine, "package") {
					tokens := strings.Fields(contLine)
					tag.Package = tokens[1]
				}
			}
			if tag.Pattern != "" {
				pattern := tag.Pattern[2 : len(tag.Pattern)-2]
				pattern = fmt.Sprintf("^%s", regexp.QuoteMeta(pattern))
				matched, err := regexp.MatchString(pattern, contLine)
				if err != nil {
					return fmt.Errorf("failed to check regex: %s. err: %v", pattern, err)
				}
				if matched {
					tag.LineNum = i + 1
					break INNER
				}
			}
		}
		tags = append(tags, tag)
	}
	newJson, err := json.Marshal(tags)
	if err != nil {
		return fmt.Errorf("failed to unmarshal. err: %v", err)
	}

	fo, err := os.Create(inputFile)
	if err != nil {
		return err
	}

	defer func() {
		if err := fo.Close(); err != nil {
			return
		}
	}()

	if _, err := fo.Write(newJson); err != nil {
		return err
	}
	return nil
}

func getRepoShortName(gp string) string {
	a := strings.Split(gp, "/")
	return a[len(a)-1]
}

func fetchRepo(gp string) error {
	updateStatus(gp, "fetching repository")
	sn := getRepoShortName(gp)
	if _, err := os.Stat(sn); !os.IsNotExist(err) {
		os.RemoveAll(sn)
	}
	cmd := exec.Command("git", "clone", gp)
	fmt.Println("running command", *cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run git clone for package %s.\nout: %v\nerr: %v", gp, string(out), err)
	}
	return nil
}

func runCtags(gp string) error {
	// ctags -R  --output-format=json --languages=go ~/atlasproxy/src/github.com/10gen/atlasproxy > atlasproxytags.json
	updateStatus(gp, "running ctags")
	sn := getRepoShortName(gp)
	os.Remove(fmt.Sprintf("%s.json", sn))
	cmd := exec.Command("ctags", "-R", "--output-format=json", "--languages=go", sn)
	fmt.Println("running command", *cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run ctags for package %s.\nout: %v\nerr: %v", sn, string(out), err)
	}
	fo, err := os.Create(fmt.Sprintf("%s.json", sn))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	fo.Write(out)
	return nil
}

func loadDataToAtlas(gp string) error {
	updateStatus(gp, "loading data to Atlas")
	sn := getRepoShortName(gp)
	filename := fmt.Sprintf("%s.json", sn)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	cb, _ := json.Marshal(map[string]interface{}{
		"project": gp,
	})
	resp, err := http.Post(ClearDataEndpoint, "application/json", bytes.NewBuffer(cb))
	if err != nil {
		return fmt.Errorf("error clearing data", err)
	}

	var da interface{}
	json.Unmarshal(data, &da)
	reqb, _ := json.Marshal(map[string]interface{}{
		"data": da,
	})
	resp, err = http.Post(LoadDataEndpoint, "application/json", bytes.NewBuffer(reqb))
	if err != nil {
		fmt.Errorf("error posting data", err)
	}
	resb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error getting request body", err)
	}
	fmt.Println("loadData response", string(resb))
	return nil
}

func makeHandleProcessFn() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			gp := r.Form.Get("gitproject")
			fmt.Println("process", gp, "in")
			updateStatus(gp, "processing started")
			err := fetchRepo(gp)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("Err fetching: %s, %v", gp, err)))
				return
			}
			err = runCtags(gp)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("Err ctags: %s, %v", gp, err)))
				return
			}
			err = enrichCtags(fmt.Sprintf("%s.json", getRepoShortName(gp)), gp)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("Err enriching: %s, %v", gp, err)))
				return
			}
			err = loadDataToAtlas(gp)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("Err loading data to atlas: %s, %v", gp, err)))
				return
			}
			updateStatus(gp, "done")
			fmt.Println("process", gp, "done")
			w.Write([]byte(fmt.Sprintf("OK: %s", gp)))
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", makeHandleRootFn())
	mux.HandleFunc("/process", makeHandleProcessFn())
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", mux)

}
