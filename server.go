package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

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

func enrich(inputFile, gitproject string) {
	tagsFile, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
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
			fmt.Println("processed", j, "lines")
		}
		// read the first two lines to extract package
		contents, err := ioutil.ReadFile(tag.Path)
		if err != nil {
			panic("err reading:" + tag.Path + ":" + err.Error())
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
					panic("failed to check regex: " + pattern + ". err: " + err.Error())
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
		panic("failed to unmarshal. err: " + err.Error())
	}

	fo, err := os.Create(inputFile)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := fo.Write(newJson); err != nil {
		panic(err)
	}

}

func getHtmlTemplateBinData() []byte {
	data, _ := HtmlIndexHtmlBytes()
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", makeHandleRootFn())
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", mux)

}
