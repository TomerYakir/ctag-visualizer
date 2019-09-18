// Code generated by go-bindata.
// sources:
// html/index.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _htmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\xeb\x72\xdb\x36\x16\xfe\xdf\xa7\x40\x91\x4e\x24\x4d\x45\x50\xb6\x15\x5b\x71\x44\xed\xa4\x69\xd3\xee\x4e\xba\xd3\xa9\xb3\xd3\x4b\x26\xd3\x01\x81\x43\x12\x16\x08\x70\x01\x50\xb6\x92\xf1\xbb\xef\x80\x17\x49\x94\x28\x3b\xce\x74\x67\x2f\xfc\x43\x12\x38\x38\x97\xef\x80\xc0\x77\xc0\x79\xe6\x72\xb9\xf8\x02\x21\x84\xe6\x19\x50\x5e\x3f\x56\xaf\x96\x19\x51\x38\x64\x0d\x8b\x70\xe6\x5c\x61\x2f\xc3\xd0\x9e\x11\x9a\xd3\x0f\x5a\xd1\x1b\x4b\x98\xce\x43\xeb\x84\x63\x59\x60\xf9\xd2\x86\xd7\x36\x8c\x4b\xc5\x25\xd8\x70\x4a\xa6\x64\xd2\x74\x92\x6b\x8b\x17\xf3\xb0\x56\x77\xa0\x7f\xf3\xee\xaf\x8e\x2d\xa6\x39\x90\xeb\x7f\x96\x60\xd6\x95\xa9\xfa\x31\x38\x23\x53\x72\x42\x72\xa1\xbc\xde\xce\x68\xa1\x1c\xa4\x46\xb8\x75\x84\x6d\x46\x4f\x9f\x9d\x07\xaf\xae\x7e\xd5\xe6\xd7\xd5\xef\xec\xed\x92\x8a\xdb\xf3\xdf\x56\xfa\xfc\x87\xa2\x60\xbf\x7f\x0f\x2e\xfe\xed\xc7\xef\x7f\xb9\x7a\x2d\xbf\xb9\x99\xfd\x90\xbc\xfa\x9b\x8e\xba\xba\x98\xd1\xd6\x6a\x23\x52\xa1\x22\x4c\x95\x56\xeb\x5c\x97\xfd\x71\x48\xa1\x96\xc8\x80\x8c\xb0\x75\x6b\x09\x36\x03\x70\x18\x65\x06\x92\x9d\x58\xb8\xba\xb6\x84\x49\x5d\xf2\x44\x52\x03\x55\x40\xf4\x9a\xde\x86\x52\xc4\x36\xcc\xa9\x03\x23\xa8\x14\x1f\x20\x3c\x21\x13\x32\x09\x99\xed\xb4\x56\xf1\x32\x6b\xf1\x03\xf9\x79\xb4\x9d\xeb\x43\x33\x9d\x74\xa1\x7d\x7b\x8b\x2e\x4c\x5a\x59\x87\x98\x14\xa0\x1c\x8a\x50\x93\xf0\xab\xfa\x26\x94\x70\xb5\xde\x6f\x21\xa1\xa5\x74\x2f\x8b\xe2\x55\x25\x3a\x1c\x30\x47\x53\xbb\x12\x36\x88\xd7\x39\x65\x83\xd1\x8b\x1e\xb5\x3c\x46\x51\xa3\x9b\xa4\xe0\xae\xc0\xac\x04\x83\x46\x43\x63\xea\x67\xc8\xb5\x83\x1f\xb5\x4a\x75\xdd\x41\x12\xca\x9c\x36\xeb\x31\x1a\xe4\xbe\x95\xc7\x01\x75\x92\xda\xc1\x88\xf0\xb8\xb1\xeb\xcd\xf5\xd8\xf3\x36\x1c\x75\xa5\x45\x11\x4a\x4a\xc5\x9c\xd0\x6a\x58\x18\x7d\x0d\xcc\x8d\xd0\xc7\xce\x88\x76\x94\x96\x40\xa4\x4e\x87\x78\x33\xf8\x52\x28\xbc\x17\x8e\xbf\x78\x4c\x98\x96\x12\x6a\xad\xd8\x56\xb2\x78\x44\x12\xa1\xf8\xf0\x23\x6e\xcc\xe0\x4b\xd4\x3c\xdd\x8d\x08\xb5\x2f\x8d\xa1\xeb\xe1\x88\xb8\x0c\xd4\x90\x6b\x66\x51\xb4\xe8\x71\xa4\x9a\xfc\x09\xaa\x24\x88\x04\x95\xba\x0c\x2d\xd0\x64\x84\x3e\xa2\x5e\x59\x7f\x7d\x35\xc4\xa4\x75\xe2\xdd\xe4\x3d\x11\x4a\x81\x79\x0b\xb7\x3e\x8d\x8d\x0b\xe8\x6b\x84\xfd\x33\x03\x6b\x85\x4a\x2f\x11\x46\x5f\x23\x6f\xc3\xcb\xd7\x63\x0f\xe3\xf4\xd7\xdd\xa1\xd9\xbb\x51\xa7\xe9\xae\x0f\x7f\xca\xf9\x4f\x94\x2d\x69\x0a\x9d\x04\x2c\xd3\x3e\xf0\xbd\xff\x45\x2d\xed\x61\xa4\x45\x01\x8a\x0f\xf1\x5c\x0a\xc4\x24\xb5\x36\x1a\x78\xb8\x69\x61\x45\x2c\x21\xf0\x4b\x1b\x18\x74\xd8\x14\x28\x1d\xe4\xd4\xa4\x42\x0d\x16\x8d\x3a\x34\x8f\x17\x3e\xd2\x62\x99\x7a\x04\xe6\x61\xbc\x98\x97\xb2\xd5\x5a\xe1\x14\xb4\xa2\xdb\x94\x76\x74\xc7\x9a\xaf\x51\x47\x32\xd8\xd1\x38\x58\xcc\xc3\x52\x2e\xe6\xa1\x14\x8b\xfd\xa9\x72\x0c\x98\xb7\xeb\x02\x7a\xa6\xe5\xd8\xeb\x1c\x23\x27\x0a\x18\xa3\x82\xba\x6c\x8c\xa4\x50\xf0\xf7\x32\x3f\x86\x59\xaf\x57\x47\xf1\xab\x63\x0b\x84\x83\x7c\xb0\x98\x53\xe4\xa8\x49\xc1\x45\x83\x3f\x62\x49\xd5\x72\x50\xaf\x72\x03\xaf\xc5\x19\xaa\x6c\xa2\x4d\xfe\x5a\x48\x78\xab\xdf\x08\xb5\xdc\x71\xb2\xeb\x59\x85\x81\xf3\x11\x35\x50\x7b\xff\x6b\xac\xe9\xa2\x17\xef\x4a\x78\xcf\x21\xb4\xed\x0a\x76\x95\x74\xe1\xfd\x24\x74\x5f\x97\x8a\xf5\xa2\x5b\x03\xeb\xdb\x1f\x05\x6f\xc7\xa7\xff\x04\xb8\xde\xe3\x16\xdc\xea\xb9\x9d\xc8\x1e\x60\x9b\x53\x29\x1f\x65\xf0\x67\x48\xc0\x80\x62\x60\x1f\x30\xed\x6d\x23\xb3\x91\xae\xf3\xd9\xd8\x7b\x44\x3e\x9a\x65\xe0\x68\x5a\xaa\x49\xff\xe8\xac\xfc\x17\x4c\xfa\xff\xe5\xbc\xdc\x6b\xbb\x93\xa6\x31\x4a\x8a\x7b\x73\x12\x86\x68\xc3\x28\x75\x69\x18\xa4\x86\x16\x59\x45\x56\x52\xe1\xb2\x32\xae\x1f\x75\xe0\x40\xfa\xbe\x3c\x88\xb5\x0b\x68\x21\xc2\x83\x86\x20\x8c\xa5\x8e\xc3\x58\x3b\x92\xea\x27\x6f\xce\x66\x97\xe7\x4f\x1d\x8d\xa3\x6d\xb0\x87\x5b\x31\xc4\x65\x9a\x82\x39\xdc\xbc\x0a\xbf\xf9\x11\x03\x85\xa4\x0c\x86\x1b\x56\x85\xc7\x08\xf7\xed\xe9\x12\x1c\xf2\x1c\x25\x29\x88\x2d\xa4\x70\x43\x1c\xf6\x89\xc5\xc4\x66\x22\x71\xc3\x9e\x2e\x03\xae\x34\x0a\xe1\x63\x68\x54\x53\xd5\x4f\x92\x36\x4e\xdf\x10\x93\x6b\x2d\x54\x65\xcc\x77\x3d\x79\xe3\x1b\x1b\xb4\x7d\xc3\x01\x02\xf8\x13\xf6\x98\x9e\xa9\xfc\xb8\x9c\xfe\x69\x60\xd4\xf1\x56\xd1\xe6\xd4\x3a\x30\x0f\x06\xbd\x1f\xde\xf6\xb9\x6f\x79\x51\x54\xae\x3f\xc0\x4f\x0d\xc7\xf9\x34\xaa\xb7\xcb\x96\x88\xcd\xf4\x4d\x9f\xff\x9f\x43\xa9\xac\xa3\xc6\x01\xc7\x87\xda\x2c\xb8\xbf\x2a\x07\x66\x45\xe5\x70\xe3\xa2\xa7\x74\x1b\xa6\xb9\x75\xf8\x6e\x8c\xce\x26\x93\x49\x9f\x4f\xa4\xd0\xd6\x0d\x11\x0e\x1b\xab\x78\xec\x55\x08\xd7\x8c\xdd\xf0\x4d\x74\x87\x1a\x9e\x69\xe0\x1e\x9a\xc9\x24\x50\xb3\xf1\xab\xc7\xe0\x63\x80\x48\xc0\xb1\x4c\xa8\x14\x71\xea\x28\x4a\x8c\xce\xd1\x4b\xcf\xd5\x7b\xd0\x40\x87\x14\xba\xa2\xf2\x5b\x06\x9d\x16\x9f\x4f\x9e\x3b\x4c\xfe\xb5\x2e\x15\xaf\x78\x2e\x1e\x57\xb7\xd1\xc3\x63\xde\x55\x15\xc8\xb7\xdf\xa0\xba\xf4\x79\x8f\x5e\x69\xa5\x80\x39\xe0\xc8\xe9\xa6\x11\xf7\xeb\xa9\x3d\xab\x6e\x24\x11\xd2\x81\x19\x56\xec\x26\x5a\x20\x7f\x6f\x49\x2e\xfa\x32\xea\x5f\x84\xd0\x66\x66\xb7\x74\x18\x45\xe8\x1d\x21\x44\xc1\x0d\xba\x02\x57\x57\x05\x39\x2d\x7a\xd5\x8e\x46\xef\xfb\x55\xb6\xca\x88\xd5\xa6\xf7\x6b\xed\x08\x25\xda\x7c\x47\x59\xe6\xc9\x7a\x85\xf2\xf1\xba\x23\x0c\xdb\x51\x47\x45\xb6\x04\xa0\x22\xff\xfd\xa6\x5b\xe8\xee\xc5\x2c\x8a\xa2\x8a\x74\x3f\x7d\x5a\xb7\x2f\x85\xe2\x55\x23\xf6\x5f\x14\xde\xb4\x5b\xa6\x8b\x8d\xf4\x68\x13\x8d\x3e\x3e\x63\x0e\xc3\x09\x24\xac\x40\x56\xab\xc9\xe1\x96\xd3\x1f\x9e\xe7\x37\x7b\xa4\x46\x13\x45\x73\xf0\xf7\x7a\x0f\xd7\xa4\x5d\x72\x8f\xc3\x70\xf7\xef\x81\xc8\x3a\x53\x32\x87\xb7\x70\xb4\xe3\x1f\x44\xc4\xd3\xdf\x87\x20\xf0\x05\xcd\x5e\xec\x95\xfd\x3a\xfc\xc6\x45\x8f\x40\xf5\xf8\x30\x08\xfe\xaa\x2a\x80\x68\xab\xe7\x7e\xe9\x5d\x68\x92\x36\xb6\xe4\x18\x38\xc9\xb1\x09\x94\x1c\xce\x20\xbf\xb8\x91\x2d\xfd\x7f\xc4\x74\x6a\x00\x6c\xd7\xfb\x07\x65\x9b\xda\x65\xbf\x60\xf9\x9c\x59\x84\x0e\xab\xf4\x6e\xd7\x91\x5a\xff\x81\x3d\xc0\x06\x4c\x2b\x47\x85\x02\x73\xcf\xd6\x89\xf6\x2a\xfa\x47\x0c\x1a\x90\x9d\xca\x7b\x30\xda\x7d\xfb\x94\xfd\x89\x64\x82\xf7\x0a\xde\x8d\x08\xa3\x8e\x65\x43\x30\xe6\xe1\x9d\x03\x8c\xd1\xc6\x8b\x1e\x22\xd8\x87\xd0\x17\xc7\x04\xee\xb6\x8f\xdb\x33\xb8\xee\x90\x79\x75\xcc\xd8\x3d\x87\x23\xf7\x9d\x6c\xf4\xb8\xce\x85\x2d\x24\x5d\x5f\xa2\x58\x6a\xb6\x44\x5f\x8a\xbc\xd0\xc6\x51\xe5\xee\x65\x8a\x0d\x68\x3d\xfa\x62\xca\x96\xa9\xf1\x1b\x67\xc0\xb4\xd4\xe6\x12\x3d\x39\x87\xf8\x79\x12\x1f\xa7\x66\xbb\x11\x85\x3b\x21\xcd\xc3\xed\x39\xf4\x3c\xd6\x7c\xbd\x73\xe4\xc9\xc5\xaa\x29\xd4\xda\x03\xb4\xc5\xde\x09\x70\x73\xe2\x28\x54\x4a\x08\xd9\x31\xc0\xc5\x6a\x47\x8f\xe7\xb9\xad\x22\x3f\x26\xa8\x1b\xb4\x44\xf6\xe4\x14\xa3\xca\x9b\x08\x6f\x40\x52\x5a\xc1\x0b\x8c\xba\xb6\x76\x9d\x31\xfa\x66\xcf\x93\x7d\x09\xa1\x8a\xd2\x05\x89\x00\xc9\x6b\x3b\xe7\x3d\x03\xaa\x41\x95\x24\x12\x3c\xc2\xa9\x70\x7f\xb4\xe7\x82\xd5\x9a\x1a\x61\x07\xb7\x0e\xb7\x3a\x57\x54\x0a\x4e\x1d\x60\x54\xd5\x2b\x99\x96\x1c\x4c\x84\xbf\xf3\xec\x0c\xb9\x0c\x3c\xcd\xdb\xf0\xad\x7f\xfc\xfc\x66\x3f\x82\x8d\x49\x49\x63\xbf\x73\x69\xd3\xb5\xb9\xf8\x5e\x38\xd4\xf0\xe4\x79\x58\x09\x1d\x51\x40\x5b\x97\x6e\xe8\x0a\x6c\x00\x49\xe2\x4d\xd6\x2f\x52\xa4\x99\x43\xb1\x53\x41\x55\x6f\x62\xa4\x15\x93\x82\x2d\x23\xdc\x65\xe2\xc3\xaf\x86\x83\x27\x3b\xe6\x07\x15\x75\x5c\x51\x59\xc2\x08\x2f\x5e\xd6\xb2\xbe\x70\x3d\xc4\xb9\x9b\xdd\xbe\x84\x87\x3e\xc1\xc7\x27\xd2\xee\x7a\x73\x24\xfb\x7b\xfa\xb3\xd3\x43\x3f\xae\x74\x0e\xc8\x50\xc5\x75\x8e\x60\x0d\x01\xa3\x8a\xaf\x11\xcb\xa8\x71\x76\xcf\xbb\xfd\xe1\xf3\x83\x08\xb6\xf3\x21\x31\x34\x87\xd6\xab\x58\x1b\x0e\xa6\x71\xaa\x7e\x09\x0c\xe5\xa2\xb4\x97\xe8\xb4\xb8\x7d\x11\xeb\xdb\xc0\x66\x94\xeb\x9b\x4b\x34\xf1\x2d\xe8\x64\x52\xdc\xa2\x09\x32\x69\x4c\x87\x17\x93\x31\xba\x38\x1f\xa3\x8b\xe7\x63\x44\x4e\x47\x2f\x30\xba\x11\xdc\x65\x11\x3e\x9f\x4e\x30\xca\xc0\xa7\x2a\xc2\xd3\xd9\x04\xef\xfd\x56\xa8\x62\x20\xcd\x81\x7a\x55\x99\xd6\x4d\x81\xd3\x39\x18\x07\xd6\x05\xc2\xa4\x85\x09\x21\x8f\x81\x37\x9d\x7f\x11\x3c\x3a\x3d\x3b\x9b\xce\x9e\x5f\xc4\xc1\x19\xe5\x67\xc1\x74\xca\x9f\x05\x74\x0a\xb3\x80\x43\x72\x3a\x79\x76\x72\xc6\xf9\x34\x79\xea\x40\x51\xe5\xa2\x73\x3e\x3d\x49\xce\xe2\x24\x98\x5d\xcc\x4e\x82\xe9\xf3\xe4\x34\xa0\x27\xb3\x38\x48\x9e\x4d\xe0\x19\x9d\x5d\x50\xe0\x27\x78\x31\x0f\x6b\x44\xfe\x2f\xe1\x9a\xce\xe8\xec\x7c\xca\x92\x80\x4d\x27\x10\x4c\xd9\xc5\x2c\x98\x25\xd3\x49\x70\x42\x2f\x66\xf4\x64\x32\x99\x4e\xd9\xe9\x9f\x01\x57\xcf\x27\x73\x74\xb1\xdc\xf9\x56\x7a\xb6\xe7\x4f\xfe\x5c\x1a\xce\x6b\x7b\x26\xff\xe6\xb4\x77\x63\x60\xf7\x34\x7d\x5f\x57\x58\xca\xc5\x51\x67\x7b\x7f\x52\xd5\x3f\x86\x68\xe9\x32\x5f\xad\x09\xf5\x8b\x70\xd9\x2b\x03\x1c\x94\xdf\x2d\x86\xbe\x50\x6a\x7e\x27\xbd\x6c\xff\xf4\xed\x74\x8f\x9a\x12\xb2\xb4\x70\x84\x08\xd4\x47\x8c\xcd\x36\xf2\x99\x07\x04\x78\xff\x90\x66\xd4\xc3\x04\x9a\x0d\xb2\xde\x15\xe7\x61\xf5\xef\xf6\x5f\x01\x00\x00\xff\xff\x0f\x0f\x6b\x28\xc2\x1d\x00\x00")

func htmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexHtml,
		"html/index.html",
	)
}

func htmlIndexHtml() (*asset, error) {
	bytes, err := htmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.html", size: 7618, mode: os.FileMode(420), modTime: time.Unix(1568838546, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"html/index.html": htmlIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"html": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{htmlIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

