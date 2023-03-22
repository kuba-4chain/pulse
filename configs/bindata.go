// Package configs Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// sample-bsvd.conf
package configs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
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
	clErr := gz.Close() //nolint:all

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

var _sampleBsvdConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x7b\x6b\x73\x1b\x37\xb2\xf6\x77\xfe\x8a\xae\xad\x6c\x49\xde\xa2\x28\x52\xbe\x25\x64\xe8\x2a\xd9\xce\xee\xea\x7d\x6d\x4b\x65\x39\x39\x67\x2b\x95\x4a\x81\x33\x3d\x1c\xac\x30\xc0\x04\xc0\x90\x62\x4e\x6d\x7e\xfb\xa9\x6e\x00\x73\x21\xa5\xc8\xc9\xb1\xfc\x21\x91\x30\x40\xa3\xd1\xd7\xa7\x1b\xd0\x8f\xe7\x75\xad\x64\x26\xbc\x34\x1a\x2e\x6b\xfa\x9f\xfb\x69\x34\x5a\xc0\xc9\x17\xfd\x37\x5a\xc0\x5b\xe1\x05\x38\xf4\x5e\xea\xb5\xfb\xf2\x1b\x8c\x16\xf0\xa9\x44\xc8\xa5\xc5\xcc\x1b\xbb\x03\x6f\xc0\x79\x63\x11\x72\xde\xb8\xc9\x4a\x10\x0e\x7c\x89\xb0\x52\x26\xbb\x81\xac\x14\x52\x83\xd0\x39\xd4\x88\x16\x44\x9e\x5b\x74\x0e\xdd\x04\x88\xd0\x68\x31\x98\xe6\xc5\x0d\x3a\x70\xb8\x41\x2b\x14\xfc\xe3\xf5\x18\x9c\x01\x5f\x4a\x07\xca\x44\xe1\x55\x8d\xf3\x50\x8a\x0d\x82\x00\x65\x3c\x98\x02\x0a\x8b\x08\xae\x16\x19\x4e\x12\x7b\x58\x88\x46\x79\x90\x0e\x7e\x3b\x9d\xac\xb2\x32\x3f\x65\xf6\x8c\x86\xab\xcb\xeb\x8b\xff\x86\xcb\x6b\x74\x63\xf8\xea\xdd\xe5\x9b\xf3\x77\xe7\x57\x57\x6f\xcf\x3f\x9d\x9f\xbe\xee\x4f\xfb\x2f\xa9\x73\xb3\x75\xe3\xd1\x02\x7e\x3b\x7d\x27\x57\x56\xd8\xdd\x69\x5f\x89\xd7\x4d\x5d\x1b\xeb\x87\xab\xde\x8b\x0c\x2e\xaf\xc7\x7c\xdc\xaf\x4a\x53\xe1\x69\x7f\xef\xd1\x02\xae\x94\xd0\xdf\x4c\x00\xbe\xd3\x1b\x69\x8d\xae\x50\x7b\xd8\x08\x2b\xc5\x4a\xa1\x03\x61\x11\xf0\xb6\x16\x3a\xc7\x3c\x9c\x1c\x77\x50\x89\x1d\xac\x10\x1a\x87\xf9\x04\xe0\xc3\xe5\xa7\xef\xe6\x89\xbb\xd1\x02\xf0\x5e\x42\x7e\x57\xcb\x4c\x28\xb5\x83\xbf\xfe\x70\xfe\xf1\xe2\xfc\xf5\xbb\xef\xfe\x3a\x86\x55\xe3\x23\x59\x92\xe3\x0a\x41\x64\x19\xe9\x23\x87\xad\xf4\xe5\x68\x01\x5f\xa5\xc9\x50\xa2\xc5\x09\xc0\xb9\x72\x66\x0c\xbf\x91\x2c\x5b\xde\xbc\x19\xca\xae\x27\x31\x52\x01\x9d\x37\x97\x76\xd9\x97\xfd\xe8\x51\xac\xfd\x03\xfa\xad\xb1\x37\x8f\x6b\xf0\xdf\x3b\x04\x8f\xce\x6b\xf4\x74\xba\xf8\xe3\x72\x46\xdf\xde\x18\xad\x31\xf3\xb0\x91\x02\x04\x5c\x5f\xbe\xf9\xff\xd7\xcf\xa1\xb6\xe6\x76\xd7\x2a\xeb\xba\xc6\x4c\x16\x3b\xa9\xd7\x20\xc2\x27\xd8\x4a\xa5\x20\x97\x8e\xb4\x05\x4a\x3a\x8f\x5a\xea\xf5\x68\x01\x85\xb1\x20\x75\x66\x2a\x9a\x9d\x05\xda\x14\x2b\xa0\xd1\x0a\x9d\x8b\x73\x3b\x2f\x62\x4d\xd7\xd6\x6c\x24\xa9\x85\x98\x20\xcf\x3b\x0a\xd3\x8e\x46\x0b\x30\x1c\x6b\x88\x6d\xde\x79\x39\x3b\x7b\x39\x99\x4e\xa6\x93\xd9\xfc\x9b\xe9\xf3\x69\x1a\x6e\x1c\xda\x65\xfa\xa5\x16\xce\x2d\x93\xa3\xf7\x4f\x04\x62\x65\x36\x48\x86\x20\x9c\x6b\xaa\x60\x07\x2b\x84\x4f\xc6\xc2\x71\xe9\x7d\xed\xe6\xa7\xa7\xdb\xed\x76\xe2\x8d\xad\xad\xf9\x37\x66\x7e\x62\xec\xfa\x09\xed\x7e\x51\x30\x67\x81\x8e\x74\xa0\x8d\x07\x6f\x2c\x0f\x16\x46\x29\xb3\xa5\x13\xf7\x6c\x9d\x68\xd7\x16\x37\x64\xd8\x8d\xa3\x8f\xde\xd8\xd1\x02\x1c\x4b\x53\x66\x81\x33\xf8\xa5\x41\x2b\xd1\xd1\x6c\x65\xcc\x4d\x53\xf7\x64\x73\xcc\x91\x43\xea\xcc\xa2\x60\x59\x69\xa3\x77\x95\xf4\x3b\xd8\x96\xa8\x23\x3d\xe9\xc2\x76\xab\x5d\xda\x8e\xf6\xda\x99\xc6\xc2\xc5\x15\xac\x90\x7e\x53\x28\x6e\xa2\x78\xdf\x7e\xb8\xe6\xf3\x68\x63\xb4\x34\x3a\x18\x01\x19\x88\xd0\x20\x94\x47\xab\x85\x97\x9b\x74\x50\x6f\x92\x12\xe9\xc7\x09\x2f\xe9\x87\xc0\x4f\x7d\x91\x44\xa1\x92\x85\xb1\x58\x05\x0b\x56\x9b\x1c\x27\xf0\xc1\xe8\x83\xe5\xc1\x8a\x56\x48\x5b\x78\x91\xf9\xe8\xc3\x2c\xd2\x8a\xa2\x29\x53\x26\x1b\xb0\xfc\xc1\x34\xbe\x35\x40\x59\x80\x36\x9a\x75\xe9\x82\x55\xc7\xe3\xf4\xcd\x63\x96\x86\x93\x79\xf0\x2f\xad\x79\x7c\xa7\xd9\x7c\x89\x49\xe7\x2d\x8a\x0a\xa4\x33\x2a\x44\xc7\xd5\x0e\xac\xd0\xb9\xa9\xe4\xaf\x24\xc0\xb0\x29\x91\x81\xcc\x62\x4e\x42\x16\xca\x81\x45\xd7\x28\x96\xb7\xa4\xe0\x48\x94\x48\x57\x3e\xb8\x8a\xc6\x2d\x64\xd2\x66\x8d\xf4\xec\x17\x28\xb2\xb2\xe7\x13\x9c\x40\xa4\x83\x8a\x73\x86\xf4\x50\x71\x16\x92\x45\x21\xb3\x46\xf9\x20\xc6\xcc\x58\x8b\x4a\x78\xec\x3b\x13\x3b\xb1\xb1\x2d\xb7\x9d\x12\xbf\xd7\x72\x83\xd6\x09\x05\x57\xaa\x59\x73\x18\xbf\x52\x62\x07\xc7\xdf\x5f\xe9\xab\x27\x44\x50\x34\xde\x54\xc2\xc7\xc8\x6a\x6a\xb2\xa3\x32\xf9\x30\x50\x46\x18\x2d\x78\x9d\x59\x79\xce\x68\x25\x85\x74\xb6\x0b\x45\x06\x15\xb5\x07\x85\x35\x15\xb8\x90\x43\x30\x87\x1c\x37\x32\xe3\xa4\x18\x62\x06\x9f\x2c\xf8\xee\x68\x11\x14\xcd\x39\x4f\x1b\xc0\xa2\x20\x7b\x92\x05\xe0\xed\x3e\xd9\x18\x10\xa2\x8f\x60\x4e\x27\x6d\x6a\x5d\x87\x13\xc6\x40\x74\x1f\x4f\xe8\x82\xd9\x93\xc5\x91\x61\xb4\x71\x09\x58\xd8\x97\x1a\x5b\xee\x6b\x24\xd7\x51\x52\x53\x82\xa0\x10\x1f\x58\x24\xa7\x8e\xc6\x08\x4f\x6d\x7e\x52\x0b\xeb\x77\xe0\xa4\x0f\x0e\x1a\x45\xd2\x6e\x2d\x7b\xce\x1a\x32\x37\xa9\x13\x85\x76\x74\xba\x9d\x69\xf8\x30\x2b\x2c\xa5\xce\xe1\xc3\xf9\xa7\x71\x8f\xbf\x76\x3f\x72\x14\x32\x43\xd2\x4d\xbe\x41\xeb\x25\x39\xe3\x68\x01\x96\xec\x85\x3f\x25\xae\x63\x0c\x25\xc2\x2e\x8a\x42\x7a\x4e\x73\xe4\x22\x18\xcc\x99\xe3\x27\xc9\xec\x28\xca\x1f\x8e\x85\x26\xa7\x8c\x98\x63\x5f\x67\x1c\x0d\xd2\x91\x64\xbd\x9c\x4d\xce\x26\x4f\x27\xcf\x86\x83\x67\xd3\xe9\xd9\x7c\x3e\x3b\x7b\xfa\x8c\xf4\xf0\xb7\x2f\xfa\x8f\x14\xdb\x54\x95\xb0\x3b\xc2\x42\x47\x22\xcf\x09\x68\x1d\x01\x19\x72\xe3\xe0\x28\x1a\xfe\xd1\x64\xb4\x18\x2d\xe0\x52\x93\xd9\x6a\xa4\xb9\xc3\xd8\xeb\xb7\x26\x9e\xd8\x8d\x7b\x64\xc8\x96\x5b\x1a\xe3\x18\xa1\xbb\x38\xcc\xa0\xe4\xb5\xf1\x25\x08\x22\xc4\xc2\x25\x48\x18\xe5\x4b\x34\x08\x0d\x0a\xcf\x5f\xb6\x42\xfb\x80\x18\xc5\x2e\x79\x24\x47\x3c\x8a\x4f\x2d\x38\x21\x95\x8b\x8d\xa4\x54\xe8\xc0\x29\xb9\x2e\xbd\xda\xb1\x67\xa3\x45\xed\x69\xc3\x5d\x82\x78\xe3\x9e\xf9\x51\x2e\xd8\x51\x84\x0e\xbe\x5f\xc8\x08\x3a\xdd\x30\x0e\x33\xa8\xec\xd9\x42\x52\x6c\x4a\x2c\x94\x2d\x8c\x4e\x50\x96\x68\x95\xc6\x11\xa6\x74\x99\x95\x2b\x4a\x15\xa8\xcc\x96\x8d\x91\x02\xdb\x4a\xac\xd4\x0e\xb6\x46\x1f\x79\xd0\x18\x12\x57\x65\x72\x3a\xbd\xd0\x3b\x5f\x92\x6c\x19\x4a\xb1\xfc\x3b\xc1\xe6\x06\x43\x1a\x4c\x59\xae\x9f\xee\xc9\x5b\x7c\x89\x36\xf2\x9f\x4b\x97\x99\x0d\x5a\xcc\x39\x70\x30\x4f\x18\xbf\x25\x3f\x69\xc5\xc9\xae\xa0\x73\x10\xca\x19\x50\xe8\x03\x1a\xb7\x58\x19\x9f\xd6\xdc\xe8\xa8\x2a\x61\x49\x97\x62\x23\xa4\x62\xeb\x4f\xa0\x33\x13\x9a\x78\xa3\x43\xf4\xf9\x68\xbf\x0d\x13\xdb\xce\x34\x31\x1a\xb7\x88\x03\x2a\x52\x5b\x4c\xe6\x45\xa3\xfa\x1e\x4d\xca\x0d\x49\x61\xa5\xb0\x72\xac\xa8\x18\xf2\xc9\xb5\x29\xd6\x3b\x53\x61\xf0\x61\x52\xc5\x71\x8d\xb6\x14\xb5\x83\xbc\x09\x8e\x0e\x85\xb4\xb8\x15\x4a\x3d\x89\x52\xed\x0c\xd4\x84\x98\x1b\xb8\x2e\x85\xce\xc7\xc1\x38\x2e\x3f\xbc\xfb\x57\x9f\x67\x9a\xd4\xda\x70\x3c\x5e\x70\x74\x1d\x65\x4f\xd1\xf8\xc2\x07\x31\x46\xac\xd6\x0f\x8a\xc7\x3d\x13\xc2\x5b\x2a\x0c\x24\x99\xa9\x43\x1f\x27\x91\x60\xdb\xc8\xba\x0f\xcd\xa2\x98\x9e\xb0\xa6\xde\x7e\xb8\x06\x87\x48\x42\x60\xe3\x64\x57\xe9\x02\x1c\x13\x8a\xa1\x2d\xa7\xea\x8a\xb0\x46\xab\x32\x2e\xab\xe2\x81\x3a\x8b\xe8\x9d\x94\x76\x08\xe6\x49\x58\xbf\xa6\x04\xd5\x37\x35\xce\x44\x42\x0f\x14\x3d\x01\xb8\x36\xe3\xc0\x70\x12\x6d\x52\x6c\xc8\x3f\x72\x83\x6a\x17\x7c\x9e\x94\x1e\xdd\xde\x50\x74\xe9\x6d\xfd\x17\x6f\x1b\xe7\x31\xff\x4b\x24\xfb\xe5\x83\xdf\x68\x01\xe7\x39\xe9\xcf\x3a\x16\xac\xbf\xcb\xe3\x49\x66\x39\x3a\x69\x39\x5a\x51\x22\x63\xa1\xd5\x68\x43\x0e\x1b\x2d\xe0\x5f\xa6\xe1\xd8\x96\x02\x17\x83\x8d\x5e\xbe\x66\x64\x35\x04\x52\xc6\xfa\x50\xc4\xb6\xe5\x26\x0d\xb1\xe2\xa8\xac\xe5\xbc\x43\xfa\x1a\x20\x06\x59\x40\xc4\x5d\xa4\xdb\xce\x00\x63\x84\x80\x18\x1e\x96\xb3\x6f\xce\x26\xb3\x17\x5f\x4f\x66\x93\x59\x7f\x74\xca\xe8\xec\x6c\xfe\xf5\xd3\xa7\x4f\x7b\xe3\x05\x7e\x3d\x9d\xcf\xfb\x33\x7f\x0c\x43\x67\x3f\x85\xa9\xf7\x8a\x29\x45\x66\x76\x8f\x14\x9e\x1f\x92\x1c\x55\x0a\x49\x76\xf0\x7f\x12\x1d\x15\x8b\xfb\xc2\xfb\xb3\xa2\x3b\xa8\xb6\x7c\x07\xa2\xa0\x14\x2e\x1a\xb8\x93\x39\x46\x23\x76\xf1\x78\x31\xae\xc7\xf2\x46\xc7\xf0\x7a\x7f\x2a\x05\x17\x13\xae\x8b\x50\xb4\x73\xa9\x3d\xc5\xb5\xa3\x7b\x8a\x4b\xe3\x9d\xe2\xd2\xc8\xa1\xe2\xde\x8b\x5b\x59\x35\x15\xe8\xa6\x5a\xa1\xa5\xc4\x2d\xf5\xca\x34\x14\xe1\x09\x67\x36\x3e\xfc\xd2\x7a\x58\x25\x6e\xf9\xe7\xe5\xec\xec\x39\xad\x7f\x1b\xeb\xcc\x95\xd0\x01\xcd\x15\x50\x49\x17\x72\x2c\x85\xe1\xb4\x4e\x9b\x38\x23\xa0\xc5\xb4\x2d\xfb\x38\x65\x3d\xa1\xc1\x65\x86\x11\x59\x11\xa0\xb6\x4b\x71\x44\x07\xac\x9c\x76\xb8\x93\xfc\x4a\x68\x5f\x5a\x74\xa5\x51\xf9\x72\x36\x9d\xd2\x1e\xff\x34\x5b\x50\x86\x6b\x3b\xa6\x7f\xb8\x10\x7e\x10\x4a\xe6\xe0\x65\x85\xd0\x68\xe9\x03\xc0\xfd\x1f\x37\x86\x6a\x0c\xe5\x7f\x88\xf0\x7b\xa9\x99\xd1\x59\xda\x26\x6f\x6c\xc0\xf5\x67\xcf\xca\xbd\x91\xd9\xac\x7c\x3a\xad\x66\xcf\x5d\xf2\x88\x6d\x29\x3d\x72\x4c\xce\xc9\x66\x75\xe8\x27\x70\x2a\x80\x8b\x2b\x37\x49\x05\x7e\x9b\x23\xb6\x0c\x06\x2e\xae\xa0\x12\x3e\x2b\x09\x70\x13\x7c\x49\x54\xba\xb0\xcd\xa8\xc2\x97\x28\x6d\x4f\x72\xa9\x16\x65\x60\xde\x2e\xea\xaa\xae\xc1\x68\xb0\x8b\xde\xac\x68\x59\xd3\xc9\xf4\xf4\xec\xd9\xe0\x53\x91\x4f\xa7\xf3\xf9\xe9\xec\x45\x5f\xdf\xbd\xac\xc2\x39\x35\x45\xf6\x3e\x78\xe2\x42\x98\x11\x94\xf3\xc2\x7a\x37\x26\x44\xcc\x67\x68\x1c\x39\x27\xd1\xf0\x26\x22\x2b\x22\x32\xcc\x3b\x83\x38\x4b\xfe\x1d\xec\x28\xd7\x8e\x36\x3e\xac\x3a\xa4\xf6\x68\x0b\x91\xc5\x82\x3d\x14\x4d\x6d\x75\x31\x6c\x6e\x0c\xc2\x73\xaa\x8a\xf6\x62\x2d\xd5\x0b\x04\xb5\x64\x28\xe0\x09\x35\x24\x0c\xdb\xb6\x22\x8f\x62\xbf\xe6\x88\x53\xab\xa4\x45\x8c\x2c\x32\x53\x55\x98\xba\x59\x5d\x44\xd9\xc5\xf8\x14\x21\x14\x61\x5a\x2e\x9a\x89\x9b\xb4\x77\xa8\xcf\x33\xb2\x04\x0a\x16\x0f\x63\x49\x6f\x20\x8f\xa8\x62\x2b\x1d\x9f\xe8\x5c\xa9\xbe\x38\x8c\x1e\x9e\x2c\xf6\x2e\x02\x76\x8b\x5f\x9e\xcc\x47\x0b\x88\x52\x5a\x26\x12\xf5\xe6\xd9\xef\xd0\xe9\xaf\x08\x21\x68\xda\x2d\x7c\xf1\xd0\xc2\xb4\x72\x3e\xbf\x93\x61\x66\x94\xa2\xd4\x70\x72\x0c\x71\xf7\x70\x77\xf7\xa2\xc8\xdb\xde\xda\x7d\x06\xef\x5e\xfb\xe3\x7c\xfe\x53\x5a\xc8\xe5\x0d\xef\xaa\x4c\x26\x54\x69\x9c\xbf\x7f\x61\xd7\xec\xd8\x5b\xfd\xe2\x73\x56\xff\x38\x9f\xcf\x1e\xda\x57\x1b\x7d\xe2\xbc\xd0\xb9\xb0\x79\x4b\xe6\xc5\xfd\x4c\xbc\xb8\x53\xce\x9f\x41\x65\xb0\xf8\x50\xe8\x9f\x41\xa1\xa7\x81\x17\xf7\x6b\xe0\x33\x08\x25\x75\x0c\x62\x51\x07\x9b\xef\xeb\x70\xa6\x76\x0e\x07\x1f\xaa\x75\x2c\xe5\x69\xa1\x54\x5c\xdb\x26\xa9\x24\xb7\x3e\x79\x86\x73\x2b\x65\x4c\x05\x85\x54\x1e\xad\xd4\x6b\xc2\xb0\x88\xf0\xfa\xe2\x6a\x3a\x9b\xcd\xc2\x5a\x9a\xc7\xd3\xc2\x2c\x17\xa8\x50\x1e\x10\x79\x2e\x89\x0f\xa1\x20\x2b\x31\xbb\xa9\x8d\xd4\xde\x4d\xe0\xef\xc6\x56\xc2\xcf\xe1\xe8\xdb\x12\xa9\xa2\x79\x35\xff\xb6\x14\xae\x7c\x75\x14\xb0\x56\x37\x77\xb9\x37\x21\x11\x0e\x61\xc6\xbb\x04\xd1\xb9\x09\x26\xd6\x04\xc1\x18\x7c\x49\xd7\x07\xf7\xde\x74\x09\xf3\x7d\xe3\x3c\x07\x26\xa9\x33\xd5\xe4\x14\x70\x84\x15\x19\xf1\x0d\x47\xa7\x47\x63\x38\x9a\xd3\x7f\x8e\x63\x8d\xfe\xe4\x88\xbb\x3d\x22\x6e\xb8\xec\xcb\x87\xc6\xa4\x4f\x49\xac\x93\x11\x1c\xbf\xf9\x7b\x6c\x67\x66\x03\x91\x7c\xe1\x26\xfa\x02\x3e\x5e\xbd\x01\x87\x76\x43\x00\x26\x06\xe8\x13\x8e\xe7\x5d\x03\x22\x8d\x67\x46\x7b\x6b\x54\xe8\x04\x34\x52\xf9\x13\xa9\x7b\xeb\x43\xe2\xcb\xca\xb6\x73\x1b\x52\x10\x2f\x21\x41\x84\x5c\x25\x75\xc1\x9a\x23\xec\x17\x2a\x1c\xb0\x4d\x40\x27\x9c\xee\x6a\x6b\x32\x74\x2e\x94\x8f\x5d\x76\xe9\xb1\x29\x5d\xaa\xfa\x38\xb7\xb4\x57\x4a\x05\xd8\x3a\x63\x2d\x9e\x7f\x78\x4b\x3f\xd7\xc2\xb9\x31\x70\x33\xd9\xd6\x99\x92\x95\xf4\xfd\xcf\x3c\x10\xe6\x10\x74\x19\x60\xd7\xc9\xa3\xdc\x57\x5c\x63\xd6\xd8\xd0\xc9\xa2\xf3\x9c\x5f\x5d\x70\x72\xec\x03\xe3\x60\x87\x5a\x54\x18\xae\xe6\x84\x73\x5b\x63\xf3\x88\xe6\x33\x6e\x66\x3b\xd3\xb6\x79\x28\x29\xf2\x39\x30\xff\xdd\x85\x7c\xa9\xd4\x2e\xf1\xa0\x50\x70\x20\x24\x28\x51\x34\x4a\x71\x75\x6c\x8a\x41\x13\xf8\xa4\xa5\x4c\xf0\x22\xaf\xa4\x86\x13\x88\x37\x03\x3d\x75\x74\x65\x55\xd2\xca\x24\x08\x9c\x9b\xd3\xdb\x52\x78\xdc\xa0\xfd\x99\x09\xfc\x9c\x78\xfc\x79\x67\x9a\x9f\xa9\xaa\x09\x53\x43\xeb\x7a\xa8\xa6\x6e\x69\x64\xe3\xbe\xc5\xad\x1e\x97\xbf\x83\x6a\x8a\x43\xc6\x1f\x46\x39\x5d\x2b\xf5\x8b\xc0\x1c\xd2\x5a\x04\x3a\x5f\x00\xe6\x50\x69\xc2\x40\xe7\x4f\xc0\x9c\x21\xd6\xf4\x5c\xee\xed\xa9\x34\xb4\x0c\x5a\x19\xf5\xd2\x27\x89\xf2\xe2\x6a\xf3\x2c\x42\xf1\xcd\x8b\x87\x51\x53\xc8\x43\xac\xab\x3f\x8a\x91\x7a\xab\xfe\x04\x4e\xea\x16\x3f\x00\x95\x9e\x1d\xcc\xa7\xc1\x87\xd1\xd2\xc1\xba\x5e\xba\x7e\xf6\x30\x60\x3a\x58\x9e\x92\xf4\xb3\x87\x31\xd3\xc1\xda\x01\x62\x79\xf6\x30\x6c\xba\x6b\xf3\xd9\x43\xbb\xdf\x09\x34\x5e\xfe\x2e\x2b\x2f\x3f\x1f\x3c\x1d\x10\x1a\xac\xff\x4c\xfc\x74\x40\xa4\xa7\x93\x97\x7f\x10\x42\x1d\xd0\x4a\x0a\x7a\xb9\x1f\x6c\xaa\x83\x6e\x40\x66\x74\xd6\x58\x4b\x88\x82\x9c\x2b\x53\x92\xf1\x06\xb7\x51\xd3\x3e\x7b\xb7\x5f\xb6\xce\x2a\x71\x1b\x67\x2e\x67\xd3\x3f\xbc\xc9\x16\x57\xce\x64\x37\xe8\xd3\x76\x1d\xd5\xf6\x93\x5b\x86\xfe\xc3\x7b\x69\x6d\x6c\xe9\xc2\xff\xbb\xbe\xfc\x70\x42\x04\x7e\x69\x24\x95\xd9\xa6\x80\xd7\xd2\x67\x46\x6a\x78\x43\x05\xf2\xc9\x49\x8c\x81\x6f\xa5\xcb\x4c\x63\xc5\x9a\x32\x0e\x07\x9e\xd1\x22\x08\x92\x42\x94\x58\x49\x25\xfd\x0e\xa4\x73\x0d\xba\xb6\xed\xbe\x42\xa0\xea\x1d\x73\x10\xd6\x34\xdc\xcc\xb5\x75\x16\xb6\xea\x2e\xf9\x86\xd7\x1e\xf1\xe5\x00\xc7\xae\x88\x96\xf6\x02\x38\x6e\x50\x53\xda\xe7\x2e\x7a\x4c\xfd\x22\x11\xa7\x84\x30\xbc\x73\x0b\x6d\xa4\x04\x67\x43\x87\x94\x1b\x57\x5c\x47\xcb\xec\x86\xef\x32\x06\x3b\x51\xc4\x4e\x21\x31\xf4\xc8\x62\x0f\xc4\x1b\xee\xdb\x6f\x70\x90\x31\x19\xcf\xd0\xd2\xcc\xe8\x42\xae\x09\xce\x61\x00\x71\xb6\xce\xfe\xc0\x39\x3f\xbd\xbb\xbe\x23\x61\x0d\xee\x20\xbb\x7e\x3e\xc7\xe9\xd0\x1b\x89\xb2\x18\x62\xa5\xd0\x84\xe2\x6b\xf8\xe4\xc6\x3d\xbb\x3f\x4e\x90\x2d\xf6\xd6\x52\x45\x1b\xd8\xf6\x8a\xd5\xf3\x18\xb8\xf3\x3d\x56\xb5\x31\x0a\xae\xe3\x03\x91\xfb\x50\xe7\x23\x01\xb1\x70\x9f\x55\xc5\x0e\x95\xb7\x42\x3b\x91\x05\x50\x8a\x98\xfa\x08\x46\x3b\x99\xf3\xe5\x8e\xe0\xe8\xf0\x2b\x5a\x43\xdf\x29\x93\x57\x52\x5b\x54\x62\xe7\x6f\x0b\x44\x8a\x30\xd3\xe9\x94\x35\xfc\x51\x78\x3c\x61\x4c\x12\xde\x3f\xf5\x68\xb7\xb5\xc6\x46\xa8\x06\x61\xf6\x1c\xfe\x06\xb3\xe9\x74\x0a\xab\x9d\xc7\x84\x38\x2a\xa9\x1b\xcf\x66\xc3\x44\x88\x06\x6f\xb4\x9c\xb1\xd3\x7e\x44\x72\x1b\x84\x52\xae\x4b\xa8\xad\x34\x96\xfc\x8d\xac\x85\x67\x71\x25\x47\xdb\x1a\x0b\xca\x6c\x4f\x8a\x3d\x0e\xa2\x39\xd2\xd4\xb4\x78\xc9\x51\xe6\x1d\x73\x6c\x6c\x5d\x0a\x3d\x90\x07\x6b\xc9\x1b\x62\xf4\x80\x52\x25\x6e\xc3\x0a\x7f\x9b\x3a\x87\x6f\x03\x98\x09\x37\x1a\xc3\xc3\xb3\x7b\xf4\xaf\xbb\x26\xe9\x95\x99\x23\x23\x0e\x0e\xf2\x91\x58\x1b\xc6\xe2\x01\x11\x8b\x6b\x61\x73\x06\x3c\xa6\x68\x33\xbe\xde\x7b\x6e\xc4\x51\x8f\x28\x69\xa3\x9d\xcf\x13\xe9\x7f\x63\xe6\xbf\x10\x6d\x22\xd5\x11\x7f\x0c\x0f\x09\xaf\x11\x85\x82\x0b\x9d\xe3\x2d\x3e\x8e\x27\xbc\x6e\xa4\x0a\xfd\xea\x4a\x48\xcd\xcf\x00\x44\xa8\x08\xa8\x58\x3e\x59\x09\x2e\xe4\x7a\xe6\x20\x89\x99\x58\xe8\x85\x17\x1e\x82\xf1\xff\x40\x90\x5d\x33\x32\x5d\xaf\xad\xd1\x5b\xb1\xed\x13\xfa\x78\xf5\x86\xdf\x7b\xdc\x32\xc5\xa0\xa1\xfb\xb9\x89\xa8\xfc\xb3\x18\x0a\x78\xd6\xa1\xb0\x59\x39\xdc\xd4\x71\x6c\x6c\xb9\x8b\xf7\x3a\xb6\xc7\xc1\x5b\x54\xe8\x43\x88\xa6\xa8\x6e\xbb\x47\x0a\x61\x1f\xa3\x43\x43\x16\x9a\x9a\x81\xb3\x06\xbc\x95\xfc\x4e\x27\xb7\xa6\xee\x88\x4d\x1f\xc7\x26\xae\xe5\x5a\x0b\x4f\x45\xe4\x0f\x68\x09\xe1\x87\x42\xfa\x8d\xc8\xf8\xcc\x5f\x78\xbf\x36\x2c\xf0\x9d\x6c\xbb\x75\x46\xbb\x85\xeb\xde\x4a\xdc\x92\xab\x3c\xa7\xe8\x47\x02\xb3\x32\xbc\x18\x71\x72\xcd\xb3\x2a\x71\xeb\xe4\xaf\xb8\xe4\x09\x8f\x23\x92\x37\x84\x51\xfe\x81\x1a\xc3\x4d\x02\x1c\xbf\x97\x5a\xea\xf5\x93\x07\x33\x4b\xbf\x9f\x31\x5a\xc0\xba\x23\x61\x8a\xf8\xf8\xd5\x63\x55\x2b\x41\xb1\x39\x3d\x45\x6b\xdf\xc8\x54\xbc\x0b\x88\xee\xe5\x29\x19\x9e\x35\xcd\xba\x8c\xc8\x21\x13\x4a\xf1\x45\xd6\x16\xc9\x80\xdd\xb0\x6f\xf2\xe6\xea\x7b\xa2\x81\x16\x8e\x65\x11\xdf\xb8\xe4\x4f\x1e\xa7\xef\x10\x1f\x84\xed\xef\x2d\xf5\x7a\xd8\x66\x89\x8d\xed\xf6\x7d\x2a\x43\x8b\xf8\x4c\x80\xf1\x08\x3a\x86\x2a\x75\x63\x6b\xe3\xb0\xab\x71\x63\x89\x1b\x7a\x2d\x4e\x56\xf4\xb3\x93\x3a\x0b\x8e\xd4\x3e\xfd\x22\x8a\xec\xd9\xf4\x5d\x3a\x28\x84\x05\x6f\x4c\xc8\x65\xb4\x41\xc7\x58\x5b\xe3\x6e\x8d\xf5\x65\x78\x5f\xb4\x2d\x23\x9e\x8a\xaa\xc2\x65\x21\x94\xc3\x5e\xa7\x30\xde\xe9\x7b\x03\xb5\xd8\xb1\x78\xf3\x98\x60\xf8\xdd\xc9\x70\x07\x6e\x93\x18\x1f\xc0\x1b\xd5\x3a\x3a\x5a\xc2\xbe\xee\xd3\x76\x79\x0b\xca\xd6\xe8\x79\x52\x9a\xc3\xc1\xec\xe0\x25\x56\x7b\x49\x12\x36\xa4\x2f\xcb\x19\x9d\x64\x15\x90\x75\x9c\xfa\xe0\x84\xb3\x07\x67\x3c\x3d\xa8\x13\x22\xb0\x09\x47\x21\x17\xa4\xe3\x05\x98\xe1\x4d\x78\xcc\xb7\xdf\x06\x20\x6d\xef\x07\xf2\x10\x56\xb9\xab\x80\x9a\x6d\xbb\x40\xaa\x96\x2c\x88\xa0\xb5\x38\xda\x02\x91\x74\x4f\x1c\x5b\xa3\x39\x17\x07\x3d\x09\xee\xc9\x76\x02\xc3\x57\xb8\x77\xf0\xcd\x14\xa5\x76\x1e\x05\x51\x13\x9e\x16\xfb\xf0\x8e\x48\xa9\xfb\x49\x43\x53\xc7\x67\x97\xfd\x03\x35\xda\x4b\x15\x7b\x30\xc2\x13\xc8\xe2\x4e\xd8\xf0\xd1\x49\x57\x5d\xb1\xc4\x5a\x98\x52\x49\xcd\xd1\xec\xde\xb2\xec\x21\x71\xb3\x8b\x05\xf0\x97\x04\x15\x9b\x5b\xe1\x65\x54\xac\x1d\x1c\x6a\xd7\xb8\xf0\x89\x40\x7d\x60\x57\x09\xbb\x46\xb2\x3f\x02\x68\xa5\xf0\x81\x50\xc7\x5c\x0c\xb5\x2f\x63\xac\x1d\x72\x38\xe4\x29\xbe\x31\x23\x0d\x9e\x24\xd5\x9d\x26\xb8\x28\x2c\x8a\x70\xff\xd8\xbd\xf9\x4c\x9b\xd0\x19\x0e\xed\x83\x81\xb2\xe3\x06\x22\x33\x09\xa2\x32\x8d\xf6\x6e\x0c\x46\xe5\x20\x75\xdd\xd0\xcf\xe4\x6f\xae\x0a\x68\x81\xd8\x71\xed\x55\x2c\x33\x82\xce\xb7\x66\x14\x7d\x89\x45\xf1\x4b\x83\x14\x48\x98\xe5\x74\xf3\x26\x3c\x45\x9d\xc8\x70\x78\x48\x7b\x88\xb5\xc5\x1a\xc7\xe1\x61\x94\x15\xd2\xf1\xe9\x5b\x33\xf5\x25\x5a\x5c\xed\x08\x30\xc4\xbe\x6b\xc5\xcf\x54\x19\x75\xdf\xa0\xda\xc5\x12\x20\x99\x71\x08\x0e\xfc\x28\x37\x4b\x29\x22\x3c\x84\x6d\xc5\xd2\xa9\x56\xba\x9e\x56\x57\xbb\xf8\x37\x17\xa2\x55\x51\xaa\xdc\xf8\xfd\xcf\xd0\x0c\xe8\x58\x54\x34\x87\x26\x2a\xd3\x4e\x3c\x3f\x76\x1e\x7d\x8b\xab\x66\xfd\x28\xe9\x87\x29\x83\x32\xeb\x75\x78\xb3\xbd\x41\x45\xa7\x0b\x0f\x08\xf8\xd7\xf8\x78\xc0\x5b\x91\xe1\x18\x72\x9a\x3f\xe6\x9b\x82\x31\x6c\x85\xd5\x63\x40\x6b\x8d\x1d\x43\x66\x25\x3f\x3a\xf9\x4f\xef\x61\x10\x3f\x03\x4b\xbd\xd3\x6f\x5d\xb3\x72\x3b\xe7\xb1\x7a\xb5\xfc\x96\x49\xbf\x1a\x77\x63\x67\xdd\xe0\x64\x32\xe1\xf7\x8e\xc8\x31\xc0\x44\xb6\xe2\x25\x58\x2e\x37\x32\x6f\x84\x82\x76\x25\x05\x2a\x2a\xdd\xf9\x8a\xe2\xe4\x84\x39\xe4\x15\x4b\x57\x9a\x6d\xba\x33\x1f\xbe\xd8\xeb\xd6\xf2\x63\x9a\x76\x05\x9d\x2b\xbd\xe2\xe7\x36\x53\xba\x2c\x89\xcd\x56\xe2\xe1\x9f\x9f\x3e\x5d\x41\x6d\x4d\x21\x15\x82\x25\x1f\x70\x3e\xfe\x65\x4e\x3b\x7c\x6f\xff\x3d\xf4\x01\xba\x27\x36\xfb\x0f\x73\xf6\xe8\xf4\x6f\x64\x32\xa1\xe3\xbb\xb6\xf4\xd7\x27\xc2\x43\xe9\x7d\x3d\x3f\x3d\x6d\xfb\x07\xf3\x6f\xe3\x52\xe2\xfe\xd5\x29\x1f\xed\xb4\xa6\x31\x30\xe4\xaa\xf1\x3e\x27\xfe\x6d\x03\x4d\x5c\xbe\x98\xbe\x98\x8d\xfe\x37\x00\x00\xff\xff\x13\xec\x79\xba\x6c\x35\x00\x00")

func sampleBsvdConfBytes() ([]byte, error) {
	return bindataRead(
		_sampleBsvdConf,
		"sample-bsvd.conf",
	)
}

func sampleBsvdConf() (*asset, error) {
	bytes, err := sampleBsvdConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sample-bsvd.conf", size: 13676, mode: os.FileMode(0420), modTime: time.Unix(1540529716, 0)}
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
	"sample-bsvd.conf": sampleBsvdConf,
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
	"sample-bsvd.conf": {sampleBsvdConf, map[string]*bintree{}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
