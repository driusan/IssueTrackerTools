// Code generated by go-bindata.
// sources:
// js/BugApp.js
// js/BugList.js
// js/BugPage.js
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

var _jsBugappJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x55\x4b\x4f\x1b\x31\x10\x3e\x6f\x7e\x85\xbb\x07\xb4\x88\x68\xa1\x55\x4f\x41\x39\xf0\x12\x50\xf1\x12\xe4\x50\xa9\xea\xc1\xec\x4e\x12\x17\xc7\xbb\xb5\xc7\x14\x84\xf2\xdf\x3b\x7e\x84\x75\x80\x6d\xe1\x12\x25\x33\xdf\x37\xfe\xe6\x99\x7b\xae\xd9\xbe\x9d\xed\xb5\x2d\x1b\xb3\x6b\xe0\x15\x96\x95\x06\x8e\x70\x20\xb9\x31\xc5\xd3\x20\xab\x85\x69\x25\x7f\xbc\xe0\x0b\x18\xb1\x3c\x60\xf3\xe1\x60\x90\x55\xcd\xa2\x6d\x14\x28\x3c\x14\xf5\x79\x63\x15\x8e\xd8\xd4\xaa\x0a\x45\xa3\x58\xb1\xc9\x88\x9a\xb9\xe8\x38\xe7\x48\xb1\x71\x2e\xcc\x2e\x99\xf6\x7e\xf1\x87\x63\xc0\x22\xdf\x16\xc6\x58\x30\xdb\xf9\x30\xa1\x69\x30\x14\xd3\x40\xa0\x67\x8e\x5b\x1a\xc0\x1b\x24\x45\x85\x37\x65\x4e\x82\xc9\x47\xec\xdb\xcd\xe5\x45\xd9\x72\x6d\xa0\x63\x39\xc0\x72\xd3\x3d\x13\x3e\xbb\xc7\x28\x08\x0a\x45\xc4\x8f\xbe\x76\xb3\x22\xbe\xeb\xc5\xe5\x70\x90\xcd\x00\x4f\x95\x40\xc1\xa5\x0f\xf4\xba\x2a\x1a\xd0\x6a\x15\xde\x4c\xe3\x3f\x39\x36\x99\x26\x02\x25\xd0\xef\xfc\xb2\x05\xc5\x4e\x7d\x99\xf2\xe0\x8a\xb9\xff\xf8\x39\x8c\x64\x09\x15\x42\x4d\x66\x27\x8e\x3c\xca\x4a\xe9\xc4\x44\x2d\xc6\x03\xc8\x7d\xc2\x55\x2d\x41\xa7\x62\x62\xda\x50\xb6\x1a\xee\x5d\x1b\x61\xca\xad\xc4\xc2\x67\xe3\x3a\x77\x6b\x67\xd4\x38\x28\x2b\xab\x35\xf9\x27\x5c\x53\x6a\x25\xc2\x03\x1e\x34\x0a\xc9\xb2\xfb\xde\x16\xb3\x2d\x1f\x6c\x8b\xe5\x1f\x69\x37\x4b\xd2\x7b\xbb\xfa\xec\x45\xe9\xc9\x41\xec\xc8\x7a\x5d\x78\x27\x30\x8d\x9f\xd6\x2f\xd6\x8e\x25\xb1\x54\xbd\x5e\xb0\x6e\xa6\xab\x2e\x7f\x31\x65\x45\x88\xeb\x82\x96\x49\x44\xf6\x69\xec\x43\xc6\xe4\x22\xe7\xc5\x9a\x1d\x49\x58\x90\xb5\x20\xfc\x15\x9f\xc1\x30\x60\x33\x3f\x02\x23\xf6\x76\xe0\xd2\x7b\xfd\x04\x64\x87\x60\x2a\x2d\x5a\x27\xb0\x17\x9e\x60\x02\xe9\x5c\x48\x30\x48\xbb\xdb\x4b\x79\x46\x04\x82\x2b\x98\x35\xbd\xe8\xe0\x0e\xd0\x2b\x2d\x1a\x2d\xf0\xb1\x17\xbc\x02\x04\xf8\x69\x4d\xd9\x8b\xa9\x70\x95\xee\x21\x74\x90\x40\x99\xf0\x59\xbf\x14\xe7\x0c\xb0\x46\xed\xf3\xea\x2e\x02\xd7\x06\x23\xf8\xf7\xa4\x74\xdb\xb4\x16\xc9\x19\x56\xec\x4b\x9c\x83\x26\xc3\x81\x14\xd5\x9d\x9b\xa6\x38\x3d\xeb\x0b\x95\xec\x3f\x03\x69\xe0\x9d\xbd\x3e\x13\x06\xfb\x7b\x9d\xf4\xf7\x9f\x12\x3f\xa2\xae\xbb\x3b\x6f\x49\xf2\xd7\xa4\x16\xf7\xe1\xcc\xb8\xa9\xf5\x5f\xfa\xa0\x59\x3e\xff\x1c\xa0\x1d\x36\xcb\xc3\xa5\x62\xd3\x86\x7a\x19\xbd\x6b\x5d\x0a\x87\x2e\x64\xe7\xbc\x9b\xff\x79\xe3\x59\x4f\xf2\x48\xac\xac\xa7\xd3\x87\x5f\xd6\x81\x4b\x71\xe0\xf6\x32\x5e\x1e\xaa\x7b\xb7\xb5\x56\xcb\x21\xab\xb8\x94\xb7\x34\x0e\x7e\x17\x1d\xf2\x61\x21\xe7\x88\xee\x4f\x4f\xc1\x1f\xf6\xfd\xfc\xec\x84\x7e\x5d\xc3\x6f\x4a\x20\x5c\xc0\x08\x28\x1b\x45\xca\xea\x47\x9f\x43\x35\xe7\x6a\x06\x6b\xd1\xc3\x72\x3f\x1f\x01\x8f\xf5\xe7\x85\x8d\xc7\x63\xf6\x95\x6d\x6c\x74\xcd\xb3\x86\x8c\xec\xcb\xce\xce\xea\x22\x44\x51\x2b\x6a\x38\x6a\x13\xba\xaf\xab\x96\x2d\x53\x21\xf4\x67\x50\xe4\xc7\x47\x13\xba\xa0\x3e\x27\xd4\x16\x52\xa5\x86\x0e\x96\x93\xbe\xdc\xfd\x1b\x00\x00\xff\xff\xe7\xbb\xf6\x88\xd5\x07\x00\x00")

func jsBugappJsBytes() ([]byte, error) {
	return bindataRead(
		_jsBugappJs,
		"js/BugApp.js",
	)
}

func jsBugappJs() (*asset, error) {
	bytes, err := jsBugappJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/BugApp.js", size: 2005, mode: os.FileMode(436), modTime: time.Unix(1453335009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsBuglistJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x3f\x6b\xf3\x30\x10\xc6\x67\xf9\x53\x1c\x7a\x17\x9b\x37\x78\xe8\xe8\xd0\xa5\x69\x87\x42\x29\xa5\x74\x2b\x1d\x84\x7d\x89\x45\x64\xd9\x48\xb2\x21\x04\x7f\xf7\x9e\xfe\xb8\x31\x84\x90\xe5\x38\x9d\x4e\xbf\x7b\x9e\xd3\x24\x0c\x3c\x8d\x87\x37\x69\x1d\x3c\xc2\x27\x8a\xda\x95\xb5\x41\xe1\x70\xa7\x84\xb5\xf9\x39\x63\x8d\xb4\x83\x12\xa7\x77\xd1\x61\x05\x3c\x35\xf3\x4d\x96\xb1\x03\xba\x67\xdc\x8b\x51\xb9\x0f\xd3\x0f\xb6\x82\xfd\xa8\x6b\x27\x7b\x0d\x79\x01\xf4\x92\x19\x74\xa3\xd1\x21\x65\xfc\x4b\x3a\x85\x3c\x22\x2c\xbd\xf7\xb5\x90\x56\xf0\xfd\x13\x8e\xbd\xa6\xf3\x4e\xc9\xfa\x88\xcd\x1a\x86\x91\xc6\x18\x96\x83\xc1\x09\xf5\x32\x36\x2f\xb6\x71\xc4\xd6\x5f\xcf\x14\x66\xca\x66\x82\x19\xd4\x0d\x9a\x6b\x45\xde\xaf\x6b\x85\x37\xeb\x5a\x69\xb7\xa9\x84\x0a\x3b\xc2\xda\x54\xa6\x31\xe4\xa7\xf4\xea\xca\x4e\x0c\xf9\x85\x32\x09\x95\xc4\x24\x6f\xeb\x95\xbd\x44\x4a\x1e\xb4\x72\x25\xa3\x49\x76\x86\x23\x9e\x2e\xab\x7b\x75\xd8\x71\xf8\x0f\x84\x82\x39\x76\xdc\x84\x30\x2e\x12\x84\x28\xad\xc1\x3d\x61\xfe\xf1\x0d\xf4\x3a\xac\xa9\x0a\x5e\x92\xda\xf5\xf6\x16\x30\x99\x53\x21\x29\x7c\x2c\xbc\xdd\x39\xc4\x3b\xea\x79\x23\xa7\x38\x58\x8f\x4a\x85\xe4\xb6\xd1\xf6\x21\x69\xfc\xeb\x65\xab\x25\x86\x6f\x0f\xd3\xef\x60\x7a\x75\x85\x59\xbe\x25\x4b\x0e\xbc\xf4\x39\x23\x07\xbf\x01\x00\x00\xff\xff\xc0\x85\x8c\x9c\xb7\x02\x00\x00")

func jsBuglistJsBytes() ([]byte, error) {
	return bindataRead(
		_jsBuglistJs,
		"js/BugList.js",
	)
}

func jsBuglistJs() (*asset, error) {
	bytes, err := jsBuglistJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/BugList.js", size: 695, mode: os.FileMode(436), modTime: time.Unix(1453335009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsBugpageJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x4d\x8e\x1a\x3d\x10\xdd\x73\x0a\xcb\xdf\x06\x24\x40\x1a\x34\x9f\x14\x31\xca\x22\x33\xc9\x22\x52\x32\x41\x24\x17\x30\xdd\x05\x38\xe3\xb6\x5b\xfe\x21\x41\x23\xb6\x59\x47\x39\x41\xce\x92\xa3\xe4\x24\x71\xbb\x9b\x86\xa6\xdd\x3f\x4c\xc8\x80\x17\x48\x54\x95\xeb\x3d\xbb\xea\x55\x7b\x45\x24\xba\x35\x8b\x09\x59\x00\x7a\x89\xa6\x40\x02\x3d\x0c\x24\x10\x0d\x77\x8c\x28\xd5\x7d\xec\x20\xbb\x42\xaa\x62\x46\xd6\xf7\x24\x82\x31\xc2\x59\x3c\xee\x77\x9c\x53\x02\x0f\x41\x8e\xd1\xdc\xf0\x40\x53\xc1\x51\xb7\x87\xd2\x6d\xc9\x4a\x00\xe6\x14\x58\x38\x15\x5f\x2c\xc2\x2e\x88\xdb\x64\x7d\xeb\x66\x06\xf6\xe3\x93\x45\xe7\xa8\xeb\x75\xa4\x70\xda\x48\x5e\xa0\xfa\x86\x41\x04\x5c\x77\x4b\xb1\xc9\xc2\x21\x5d\x59\xaa\x3e\xd7\x23\x0a\x92\x43\x66\xc7\x92\x96\xa0\x63\x8a\xd1\xc6\x1f\xdf\x1a\xb3\x01\xb7\x8c\x1d\x08\x36\x88\xc2\xc1\x55\x25\xb4\xcb\xf8\xeb\x27\xf6\x3a\x7b\x67\xa2\x3b\x42\x8c\xcc\x80\xa5\xbf\x03\xca\xe7\xa2\x96\x7f\x52\xf1\x8b\xe2\x7f\x8d\x66\x24\xb4\x8d\x5c\x47\xda\xf5\xe1\x99\x59\xf3\x02\xeb\xff\x9f\xd8\x24\x25\x6b\xef\xa6\x60\xda\x14\xfe\xa5\x32\xdb\x45\x1c\x40\xc6\x92\x0a\x49\xf5\x3a\x91\x74\xa6\xee\x2e\x9e\x64\x46\xdc\x47\x7a\x49\xd5\x30\x96\x22\x56\xc3\xad\xf5\xe0\xbe\x94\x26\xda\xa8\x6c\x2a\xe4\x29\x3e\x3a\x6b\x31\x41\x6a\x3b\xd8\x1e\x51\x06\x4a\x0b\x0e\x85\xed\xef\xb7\xd6\x62\x86\xdc\xbc\x77\xe4\xb6\x73\xc4\x53\x22\x6e\x18\x2b\x5a\x5a\x15\xbd\xa2\xd8\xde\xd6\x7c\x81\x02\xc1\x35\xa1\x1c\xa4\xb7\xdc\xa7\x1f\x7f\x9f\x4d\x34\x13\x5a\xda\xd1\x3c\x33\x0b\x65\xa2\x88\xc8\xf5\x89\x26\xe1\x72\x54\xd3\xe3\xe5\xcb\xdc\x5f\x7b\x45\xfc\x44\x35\x3b\x9b\x12\x5b\xb3\x7c\x0d\x2a\x90\x34\x4e\xbe\x71\x6d\x65\xf8\x1c\xd5\x75\x12\x51\x95\xf5\xdc\xca\xd9\xef\xcd\xa5\xea\x77\xe7\x52\x2c\x1f\xad\x53\x73\xd0\x93\x6b\xe6\xba\xb5\x52\xec\x0b\xe6\x1d\x55\xba\x6f\xf3\xb8\xae\xb2\x39\x3e\xe8\x25\x48\xf4\x56\x29\x03\xaa\xe2\x46\xed\x2e\x35\xde\xaf\xf5\x2b\xc6\x12\x9b\x3f\x5a\x70\xeb\xbb\x63\x34\x78\x80\xb0\xb0\x4b\x70\x87\xb5\xf3\x96\xb6\x6f\x9e\xf7\xde\xae\x46\x7f\x39\x62\x0c\x6b\xd5\x83\xb1\x7d\x34\xfa\xa7\xd9\x51\x70\x0e\x92\xd1\xd6\x1f\xfb\x58\xc2\x8a\x0a\x53\xdd\xfc\x47\xa3\x3b\x06\xa4\x86\x40\x4a\x62\x29\x61\x6e\xf1\xff\xab\x05\x7e\x12\xb8\x23\xa0\x62\xc2\x1b\x38\xa4\x3c\x30\x91\x94\x0c\x96\x34\x0c\x81\x63\x4b\x48\x4b\x53\xff\xe6\xc9\x21\x7e\x7f\xfb\xee\x7f\x4c\x6c\x57\xc5\xe4\xcd\x13\xa0\xc9\xf6\xf2\x2b\xe3\xca\x03\xb1\x26\xf1\x3f\x6b\x92\xf4\x35\x70\xbe\x16\xe9\xdb\x79\xe1\xc6\xc1\xc1\xa8\xb8\x25\xc1\x43\x53\xa9\xf0\x34\x7d\xca\x68\x81\x98\x9d\x6a\x97\x7e\xd5\x1c\xbe\xea\xcb\xd5\x22\xbe\xb7\xf4\x50\x43\xbe\x0b\x16\xec\x8f\x06\xc1\x1e\xdb\x1c\x4d\xdf\xf4\xfc\x5f\xf6\xb4\xde\x74\x36\xbd\x9b\x3f\x01\x00\x00\xff\xff\xca\x4d\xd7\xd1\x56\x10\x00\x00")

func jsBugpageJsBytes() ([]byte, error) {
	return bindataRead(
		_jsBugpageJs,
		"js/BugPage.js",
	)
}

func jsBugpageJs() (*asset, error) {
	bytes, err := jsBugpageJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/BugPage.js", size: 4182, mode: os.FileMode(436), modTime: time.Unix(1453335009, 0)}
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
	"js/BugApp.js": jsBugappJs,
	"js/BugList.js": jsBuglistJs,
	"js/BugPage.js": jsBugpageJs,
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
	"js": &bintree{nil, map[string]*bintree{
		"BugApp.js": &bintree{jsBugappJs, map[string]*bintree{}},
		"BugList.js": &bintree{jsBuglistJs, map[string]*bintree{}},
		"BugPage.js": &bintree{jsBugpageJs, map[string]*bintree{}},
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
