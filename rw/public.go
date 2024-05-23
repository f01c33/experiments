// Code generated by go-bindata. DO NOT EDIT.
// sources:
// keys/public.key (808B)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _keysPublicKey = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd3\xb9\xb6\xa2\xcc\x02\xc5\xf1\x9c\xa7\x38\x39\xeb\x2e\x06\x8b\xc1\xa0\x83\x02\x0a\x29\xc0\x52\x19\x54\xc8\x80\xc3\x64\x31\x28\x32\x3f\xfd\x5d\xdd\xf1\xb7\xd3\x5f\xba\xff\xff\xfb\x3b\x0d\x9d\x30\xf9\xf1\x7c\xf8\x73\x0d\x35\x17\xeb\x3f\x0e\x8a\xfe\x01\x73\xc6\x58\xc7\x2f\x48\xb4\x92\x7e\x2a\x5a\x9f\x8e\x0b\xaf\xc1\x1b\x32\x21\xbc\xe8\xb0\x54\xe1\x5f\xd7\x4b\x47\x87\x25\x82\x73\x3d\xea\xd7\x7b\x4e\x09\xd7\x0f\xe7\x2d\xba\xe3\x5f\x6c\x31\xe6\x4a\x97\xbe\x67\x29\x7b\x2d\x94\xed\x70\x48\x28\xfd\xaa\x49\xf5\x68\x4e\xe2\xdd\x2b\xc1\x6d\xed\xe5\x19\xe7\x4b\xc2\xf9\x52\xdb\xe8\x56\x9e\x3a\x10\x7d\x82\x49\x57\x9f\x2e\x6b\x3d\x86\xc8\x63\x3c\xf7\x48\x88\xfe\xb1\x6f\x62\xd0\x5e\x2b\xd4\x5c\x62\x0a\x45\xe1\x53\x4d\xd9\x8d\x6a\x8f\xc0\x06\x11\x2c\x13\x55\x51\x5b\x25\x1e\x5f\x80\x76\xd3\x18\x6f\xef\x33\x00\x1f\x1c\xfa\xeb\x4d\x15\x99\x39\xcd\x86\xb0\x8d\x7a\xf7\xf5\xaa\xbd\xcb\x1e\xfc\xfa\x8e\x5d\x85\x7d\x4e\x7f\x59\xaa\xdf\x35\x3d\x3b\x2f\x25\xbd\x5f\x9f\x58\xba\x8e\xf9\xd6\x0b\xa4\x3a\x64\xeb\xd1\xf0\x7c\x63\x92\x64\x0c\x31\xa3\x5f\x32\xcb\x1a\xe2\x8a\xdf\x1a\x5d\x4f\x7c\xae\x9d\x32\x22\x46\x42\x30\x3f\x79\x41\x3a\xd9\x06\x36\xa2\x6f\xed\x46\xf7\x28\x28\x92\x82\x12\x6e\xf1\xb8\x5a\x11\x54\xb7\xad\xfc\xe2\x31\x84\x17\xc6\x4c\x2d\xb1\x1d\xcb\xca\xda\x85\x77\xde\x77\x74\xdd\x0a\x7b\x97\x5a\x9a\x29\x12\xb1\x45\x9f\x07\xc1\x2e\xd4\xe6\x61\xfb\xd0\x34\x95\xe5\x20\x96\xed\x63\x95\xd3\x2c\x54\xd4\xd2\x4d\xe3\x14\x31\xc9\x8b\x8c\xac\xb4\xe2\xa1\x5f\xb1\xf2\x2a\x72\x22\xc5\xb7\xf0\xdc\x8b\xdd\x81\x4c\x9d\xf1\xdd\x7e\x9d\xb9\x37\x05\xce\x59\x56\x18\x79\xa5\x70\xe5\x86\xbb\x68\xee\xa1\x59\xf9\xc8\xfa\xce\x6e\xca\x04\xfe\xa8\x3c\x5b\x71\x80\x28\x22\xb7\xe6\xe5\xb0\xf1\x6d\x19\x0c\x65\xe6\x3d\x31\x1c\x17\xaf\x13\xd8\x1d\x68\x73\xbc\x85\xe2\x5e\xb7\xb5\x93\x4c\x83\xf6\xc6\xea\xa8\x82\x19\xb9\x5b\xeb\x37\xcc\x2a\xa7\x16\x3a\x1e\x00\x8f\xc9\x58\xd7\xae\xa1\xb2\xac\xb3\xe6\x1d\xa0\xdb\xec\x64\xaf\xfa\xe4\x5c\x1b\x8a\x11\xec\x1a\xa5\x9b\x9c\x18\x5e\x2c\xc9\xb7\x1e\x4f\xae\xb5\xd7\x44\x05\x48\xb2\x98\x33\x00\xcf\xa3\x3d\xa8\x84\x5c\xf9\xde\x50\xe4\x87\x67\x84\xa7\xa3\xdf\x28\xcf\x35\xf9\x20\xfa\x1e\x9e\xf2\xca\x67\x8d\x2f\x9d\xda\xc8\x2e\xde\x67\x6e\xd3\xd6\x32\xf5\x83\x4b\x9a\xd2\x01\x96\x09\xe3\x5f\xe8\x09\xa4\x00\x56\x56\x27\x6f\x72\x83\x40\xa7\xd6\x3c\x6e\xf7\x6a\x84\xb0\xeb\xa7\x62\x7c\x17\x4d\xa5\x0b\xa3\xd5\x65\xe1\xf8\xdd\xd9\xfa\x04\xac\x3a\xf8\xf2\x5f\x96\x68\x66\x17\x02\x26\x5f\xc6\x38\x33\x79\x45\x10\xd9\x73\x61\x72\xc6\xf9\x81\xf0\x62\xa9\x3a\x5c\x10\x84\xb7\x3f\x7f\x98\x7f\x8f\x47\xc4\xf8\xcf\x10\xfe\x1f\x00\x00\xff\xff\xd5\x9c\xc2\xcc\x28\x03\x00\x00")

func keysPublicKeyBytes() ([]byte, error) {
	return bindataRead(
		_keysPublicKey,
		"keys/public.key",
	)
}

func keysPublicKey() (*asset, error) {
	bytes, err := keysPublicKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/public.key", size: 808, mode: os.FileMode(0666), modTime: time.Unix(1568851326, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5b, 0xc5, 0x2e, 0x95, 0x8b, 0x76, 0x9f, 0x18, 0xf5, 0xe7, 0xff, 0x96, 0xbe, 0x25, 0xc3, 0xba, 0x44, 0xd2, 0x31, 0x39, 0xc4, 0xfd, 0xcb, 0xa0, 0x38, 0xa4, 0x22, 0x7, 0x41, 0x71, 0xed, 0x9f}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"keys/public.key": keysPublicKey,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"keys": &bintree{nil, map[string]*bintree{
		"public.key": &bintree{keysPublicKey, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
