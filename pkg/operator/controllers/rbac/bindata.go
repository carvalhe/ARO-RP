// Code generated for package rbac by go-bindata DO NOT EDIT. (@generated)
// sources:
// staticresources/clusterrole.yaml
// staticresources/clusterrolebinding.yaml
package rbac

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clusterroleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x5a\x4f\xaf\xe3\xb8\x0d\xbf\xe7\x53\x18\xdb\xc3\x02\x05\x5e\x06\x45\x2f\xc5\xeb\x71\xb7\x28\x0a\x14\x5d\x60\x30\xed\x9d\x91\x19\x9b\x1b\x59\xd4\x48\x54\xde\xa4\x9f\xbe\x90\x2c\x39\x76\xfe\xd9\x71\x3a\x73\x7a\x36\x45\xf3\x47\x51\x14\xff\xe5\xfd\xa1\xfa\x85\x6b\xac\x1a\x34\xe8\x40\xb0\xae\x76\xa7\xaa\x05\x75\xf8\xd4\xa0\xa9\xc9\x2b\x3e\xa2\x3b\x29\x50\x2d\xfe\xb5\xfa\xf5\xb7\xea\x5f\xbf\x7d\xa9\xfe\xf6\xeb\x3f\xbe\x6c\x37\x60\xe9\x3f\xe8\x3c\xb1\x79\xaf\xdc\x0e\xd4\x16\x82\xb4\xec\xe8\xbf\x20\xc4\x66\x7b\xf8\x8b\xdf\x12\x7f\x3a\xfe\x69\x73\x20\x53\xbf\x57\xbf\xe8\xe0\x05\xdd\x67\xd6\xb8\xe9\x50\xa0\x06\x81\xf7\x4d\x55\x29\x87\xe9\x83\x2f\xd4\xa1\x17\xe8\xec\x7b\x65\x82\xd6\x9b\xaa\x32\xd0\xe1\x7b\xe5\x4f\x5e\xb0\x7b\x07\xc7\x6f\xde\xe1\xc6\x05\x8d\xfe\x7d\xf3\x56\x81\xa5\xbf\x3b\x0e\xd6\x47\x21\x6f\xd5\x4f\x3f\x6d\xaa\xca\xa1\xe7\xe0\x14\x66\x9a\xe2\xce\xb2\x41\x23\x5e\x40\x82\x47\xbf\xa9\xaa\x23\xba\x5d\x5e\x6e\x50\xd2\x5f\x4d\x5e\x96\x0a\x34\x7b\x6a\x3a\xb0\x3e\xbd\xa2\xa9\x2d\x93\x91\xfc\x76\xc4\xf2\xa8\xa9\x23\x71\x60\x1a\xec\xdf\xe3\x4e\xbc\x05\x55\x5e\xb9\xce\x4f\x36\x1a\xd0\x0b\x1a\x39\xb2\x0e\x1d\x2a\x0d\xd4\xdd\x5e\xca\x54\xae\x87\x07\xc1\xce\x6a\x90\xbc\xe2\xd0\x6a\x52\xc9\x94\x8a\x8d\x38\xd6\x1a\x5d\x59\xea\x77\xf1\x35\xb0\x40\x4f\xf2\xe8\x8e\xa4\x10\x94\xe2\x50\xb4\xce\xb4\x47\x56\x8a\x0f\x1f\x20\xaa\x5d\x66\xaf\xa8\xed\x27\xcd\xcd\xb5\xc4\xab\xcf\xa1\xee\xc8\x47\x67\x72\xd8\x90\x17\x37\x76\xa2\x6b\xc1\x5d\x10\x10\x32\xcd\x07\xee\x5a\xe6\x43\x7f\x2e\xa1\xff\xa8\xdf\xcc\x11\x34\xd5\x0f\x79\x56\xec\x11\x2c\xe1\x37\x41\x13\xf5\xf4\x77\x95\x53\xc1\x0b\x77\x85\x58\xe3\x9e\x0c\xbd\x06\xba\xc8\x26\x60\xe9\xb5\x13\xcc\x02\xd0\x6d\xd9\xa2\xf1\x2d\xed\xe5\x1e\x90\xc3\xaf\x01\xbd\x0c\xce\xb3\x0a\x2d\xdd\xa2\xeb\x1b\x96\x5d\xd7\xe1\x91\xfc\x70\x9c\x35\x60\xc7\xc6\x63\x76\xd5\x1a\xad\xe6\x53\x37\x5c\xb8\xec\xfc\xc3\x7a\xbc\xf0\xb8\x0f\x3a\x13\x56\xaa\x37\x63\x87\xb3\x12\xbd\x6f\xfd\x40\xa4\xa5\x97\xca\xf1\x8c\x64\xd5\x47\xe5\xb5\xaa\x07\x69\xd1\x48\x0e\x3b\x77\x3d\x53\xf8\x80\x26\x9e\x27\x7e\x5c\x00\xa5\xe0\x8f\xb7\x05\x5f\xa6\x92\x6b\xb9\x1e\xf5\xde\x87\xdd\xef\xa8\x04\x94\x42\xef\xcf\x18\x93\xc5\x94\x33\x26\x6b\xb7\x3f\x7a\x5a\xb1\x45\xb6\x75\xac\x71\x47\xa6\x26\xd3\xf8\x4b\x7a\xf6\xde\x4b\x8e\xb2\xb4\x38\x59\x3d\xa3\x56\x79\xbd\x61\xb2\x1f\x62\x96\xd1\x6e\x1d\x7a\x71\xa4\x5e\x09\x8e\x41\xd8\x2b\xd0\x64\x9a\x6b\xa4\xa4\x12\x1b\x01\x6d\xb9\x2e\x9c\xaf\x38\x7b\x81\x5a\x76\xf0\x53\xc4\xb7\xaa\x03\xd5\x92\xc1\x97\x15\xd9\x25\xf2\x35\xaa\x63\xf3\x3b\xef\x7a\xac\xfc\xb0\x46\x7a\x20\x5d\xcf\x6c\x30\xf1\x9c\x83\x5e\x26\x7c\x6f\xc0\xa5\x51\x4f\xa1\x13\xda\xc7\xa0\x84\x0f\x92\xf4\x88\x89\x1a\x93\x9c\x31\xe5\xb4\x95\xbb\x50\x9a\x43\xad\x1c\xd6\x31\x1e\x82\x9e\xf3\x90\x81\xd1\xbf\x08\x9b\x0e\x61\x3e\x61\xf7\x89\xdd\xdf\x88\xda\x17\x07\x38\x78\x2f\xdb\xd8\x0f\xb0\x9b\x10\x8f\x7d\xc1\xef\x0b\xb6\x1f\x62\x58\x6d\x7c\x7e\xda\x23\x48\x70\xd8\x0c\x95\x29\x75\x50\x0a\x61\x32\x7b\x07\x5e\x5c\x50\x91\xa5\xd0\x62\x1c\x28\x5f\x1b\x94\x0f\x76\x87\xfe\x85\xa3\xaa\xf9\x31\xab\xd3\x86\xec\xe1\xd6\x71\x0c\x56\xc3\xcb\x37\xca\x12\xbc\x6a\xb1\x0e\xeb\xaf\x57\xde\xd6\xdc\x09\xf6\x5c\x4a\x53\xcd\x1f\x46\x33\xd4\x13\xa3\xc4\x32\xd1\x19\xd0\x9a\x1b\x4d\xe6\x30\x59\xbb\x22\x18\xce\xae\x78\x69\x5a\xab\x43\x43\x53\xd2\xd7\x40\xea\xe0\x05\x9c\x4c\xc8\x27\xe8\xb4\x87\xce\x3e\xce\x1b\x8f\x77\x1d\x0b\x2f\xab\xc1\xa4\xad\x27\x63\xcf\xd8\xc0\x72\x9d\x8f\x4b\xb1\x31\xa8\x84\x8e\x24\x27\xd5\xa2\x3a\xac\xd6\x82\x5d\x4d\xe6\x71\xda\xd7\x08\x8f\x7b\xb9\x07\x00\x43\x47\x7b\x57\xfa\xd0\xd4\xe9\xf5\x95\x74\xdf\x0a\xde\x87\x28\x9d\xe2\x0a\xd1\x7b\xcd\x1f\xf9\xac\xb6\xe7\x82\xfd\x1e\x52\xe4\x8e\xf7\xa1\x83\x72\x4f\x88\x1d\xc9\x49\xe3\x11\xf5\xff\xa3\x23\x6a\x51\x77\x33\x5e\x12\x59\x54\x0b\x4e\x1c\x5a\xf6\x24\xec\x68\xad\x5d\x53\x24\x99\x81\x1b\x47\x9b\xf4\x28\x0e\xa1\xfb\xee\x80\x09\x65\xc0\x9e\x4b\x51\x4f\xca\x15\x68\x46\x3b\xca\x6f\x8b\x8b\xc3\xf4\x51\x6e\x21\x4f\x4b\xef\xf6\x38\xbb\x27\x01\xd6\x05\xb3\x3a\xa6\xe6\x18\xbf\x14\xbc\x36\xde\xa1\x62\xb7\xb6\xa4\x88\xd7\x41\x19\xda\x2a\xa3\xf6\x37\x01\x72\xd8\x7a\x03\x11\x50\x6d\x6c\xac\xde\x5e\xee\xd3\x73\x79\x37\xb3\xb5\xcc\xd5\x22\x68\x69\x87\x50\x39\xd0\xa7\x2f\xab\xbb\xd7\x2c\x60\x72\xc3\xe7\x0f\x5c\x80\x0c\x3a\x17\x8c\x50\x87\x63\x07\x38\x77\xe5\x63\xea\x21\xec\x50\xa3\x8c\x49\x13\x5c\xcb\xac\x6f\x90\xd7\x6e\x09\x05\xf4\x9f\x6f\x57\x87\xe0\x30\x2d\xb7\xec\xcf\x05\x41\x3f\x3e\xc8\x3d\xd5\x3a\x40\x47\xea\x7e\x14\x1f\x4d\xf1\xf8\xa1\xa3\x5e\x4b\xa6\x66\x66\x96\xe3\x85\x5d\xba\xfd\x43\x11\x95\x29\xb9\x00\x1b\x24\xac\xdd\x1b\x9b\x14\x88\x4d\xb3\x55\xec\x90\xfd\x56\x71\x77\xa3\x70\xd4\xe8\xa4\x03\x13\xe3\xc7\xf8\x98\xc7\xf4\xc1\x04\x59\xe6\x60\xff\x1d\x0e\x8f\x1d\x4a\x8b\xc1\x5f\x11\x52\x63\xde\x6f\xaf\x9f\x5c\x4d\x64\x48\x0b\x86\x13\xcf\xda\xb0\x93\xaf\xf9\xb2\x76\x6d\x52\x79\x62\x8a\x57\x99\x64\x59\x93\x2a\xe5\x65\xf2\xb1\xb0\x33\x65\xd4\x64\x50\x26\xe3\xdd\xd7\xd4\x5c\x12\x1d\x7b\xdd\x1c\x07\x29\xd6\x2f\x1f\xda\x03\xbd\xa6\x42\x74\x88\x7b\x3e\x99\x63\xb8\xd2\x30\x14\xeb\x37\x4b\xf7\x91\xb5\xd6\x28\xc2\x35\xde\x55\xa1\x84\xa6\x41\x85\x15\x00\x0b\x8d\x7c\xab\x49\xba\xe8\xf1\x26\x9d\x91\xf2\x54\x3b\x1a\x7a\xac\x8b\xe8\x79\xee\x93\x94\x27\x6f\xc0\xfa\x96\xe5\x72\x4c\x7f\x6e\xa1\x50\x54\x3d\x4a\xbd\x91\x31\x56\x8c\x49\xbf\x89\x33\x96\x33\xb9\x90\x14\x23\xf3\x45\xd3\x17\x49\x67\xb6\xc9\xdd\x8d\x4b\x93\xa6\x29\x93\x6e\x44\x9c\x72\x35\xa7\x5d\x5a\x31\xe4\x05\xe4\x40\xbf\x83\x9b\xef\xbc\x82\x49\x84\x7b\xf1\x58\xfd\xc3\x80\xa6\x40\x40\x73\x93\x69\xe3\xf3\xcb\xca\x4c\x1a\x5c\x32\x5e\x40\xa7\x9e\x68\x7a\xcf\x14\x9b\x9a\xce\x7e\x51\xc8\x4d\xd2\x67\xaa\x4e\xbf\xb5\xb0\xf3\xca\x91\x7d\x21\x62\x5b\x50\x87\x68\x9d\xed\xb2\x8d\x66\xf6\x0e\x0c\xed\x67\xa6\x0b\xd7\x50\xd1\xc3\x4e\x37\xdb\xbd\x9a\xbc\x0b\x69\x1b\xbb\x50\x37\x25\x04\xc6\xfc\x87\x2a\xc4\xc6\xe2\xb5\xcb\x6f\xfb\xa1\xe3\x76\x7e\x6e\x9d\x39\xf3\xac\x61\x35\x5c\x1a\x21\xcc\x42\x25\xae\xdb\x73\x9a\x3b\x16\x7c\x46\xf0\x4a\xe5\xd3\x8f\x79\xb3\x83\x1f\xab\x09\xeb\x32\x70\xbe\xfc\x19\x70\xb1\x47\x2c\xc1\x7a\x16\xe4\xc1\xd6\xee\xfe\x9a\xfc\x23\x67\xec\x8f\xf4\x8b\xa9\x77\x76\xca\x1d\x64\xad\xfc\x3c\xc8\x79\xbe\x55\x5b\x83\xd5\x07\xfe\x47\x89\xbf\x0c\x0d\x5e\x4a\xbb\x25\x40\x6c\xc9\xf4\x93\xb1\x39\xfb\x81\x69\x10\xb4\x66\xf5\x4a\x9d\x3b\xa0\x3e\x0d\x76\xfe\x36\x65\xdf\x6f\x31\x89\x79\x71\x40\xab\x27\x37\x25\xed\x6f\x73\xa2\xbb\x6b\xef\xfc\x7b\x7f\xa9\x12\x46\xf5\xd6\xc5\x4a\x2e\x0b\x6e\x2c\xad\x55\x71\x46\xb3\xcb\x22\xc7\xd3\xb9\x09\x8a\x85\x4d\xff\xb9\x02\x0b\x8a\x84\xa6\x9d\xcb\xf5\x3e\xce\xcd\xf7\x4a\x75\xcb\x3f\x40\xcc\xfd\x78\xe0\xf8\x80\xae\x30\xa7\xa4\x6e\x4a\xf6\x7f\x4c\x5d\xab\x57\x30\x38\xf7\x8b\x86\x75\xbc\xa7\x12\x9c\xd2\x07\x2b\xc1\x82\x6d\x1c\xd4\xb8\xed\x4b\xab\x39\xd8\xcc\xfd\x52\xc8\x08\x7e\xf6\x5f\x04\x46\x95\x10\xa5\x6a\x79\xf0\x86\xf8\xf1\xf7\xc3\x8d\x2c\x19\xf0\xd4\x81\xb5\xb7\xbb\xff\x2b\xc9\x1f\x2d\x3a\x84\x1d\x07\x99\x99\x1b\x91\x3d\x0f\x35\xf8\x88\x4e\xf7\x18\x29\x7e\x90\x75\x18\xab\xc8\xa7\xe2\x95\x61\xf3\x39\x43\xfc\xfb\xf3\x3f\x33\xf7\xcf\x7f\xfc\xf9\xfa\xf3\xff\x05\x00\x00\xff\xff\x16\x19\x5d\x78\xa8\x25\x00\x00")

func clusterroleYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterroleYaml,
		"clusterrole.yaml",
	)
}

func clusterroleYaml() (*asset, error) {
	bytes, err := clusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterrolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcd\xbd\x0a\xc2\x40\x0c\x07\xf0\xfd\x9e\x22\x2f\xd0\x8a\x9b\xdc\xa8\x83\x7b\x41\xf7\xb4\x8d\x1a\xdb\x26\x47\x92\x13\xf4\xe9\x45\x70\x93\x3a\xff\x3f\x7e\x58\xf8\x4c\xe6\xac\x92\xc1\x7a\x1c\x5a\xac\x71\x53\xe3\x17\x06\xab\xb4\xd3\xce\x5b\xd6\xcd\x63\x9b\x26\x96\x31\xc3\x61\xae\x1e\x64\x9d\xce\xb4\x67\x19\x59\xae\x69\xa1\xc0\x11\x03\x73\x02\x10\x5c\x28\x83\x3f\x3d\x68\xc9\x68\xda\xb8\x51\x32\x9d\xa9\xa3\xcb\x27\xc7\xc2\x47\xd3\x5a\xfe\x58\x09\xe0\x87\x5a\x7b\xf6\xda\xdf\x69\x08\xcf\xa9\xf9\x8e\x4e\x4e\xb6\xd6\x7e\x07\x00\x00\xff\xff\xc4\xb6\x1b\x05\xeb\x00\x00\x00")

func clusterrolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterrolebindingYaml,
		"clusterrolebinding.yaml",
	)
}

func clusterrolebindingYaml() (*asset, error) {
	bytes, err := clusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"clusterrole.yaml":        clusterroleYaml,
	"clusterrolebinding.yaml": clusterrolebindingYaml,
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
	"clusterrole.yaml":        {clusterroleYaml, map[string]*bintree{}},
	"clusterrolebinding.yaml": {clusterrolebindingYaml, map[string]*bintree{}},
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
