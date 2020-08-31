package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

type FsBackend struct {
	BasePath string
}

func (b *FsBackend) GetFile(filename string) ([]byte, error) {
	f, err := os.Open(b.BasePath + filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (b *FsBackend) PutFile(filename string, content *bytes.Buffer) error {
	return ioutil.WriteFile(b.BasePath+filename, content.Bytes(), 0755)
}

func (b *FsBackend) FileExists(filename string) bool {
	_, err := os.Stat(b.BasePath + filename)
	return err == nil
}

func (b *FsBackend) MkdirAll(dirname string) error {
	return os.MkdirAll(b.BasePath+dirname, 0755)
}

func (b *FsBackend) GetDirectories(dirname string) ([]string, error) {
	files, err := ioutil.ReadDir(b.BasePath + dirname)
	if err != nil {
		return nil, err
	}

	var results []string
	for idx, file := range files {
		if file.IsDir() {
			results = append(results, files[idx].Name())
		}
	}
	return results, nil
}

func (b *FsBackend) GetFiles(dirname string) ([]string, error) {
	files, err := ioutil.ReadDir(b.BasePath + dirname)
	if err != nil {
		return nil, err
	}

	var results []string
	for idx, file := range files {
		if !file.IsDir() {
			results = append(results, files[idx].Name())
		}
	}
	return results, nil
}
