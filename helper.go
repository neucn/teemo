package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func getImgPath() (upPath, downPath string) {
	dir := getExecutableDir()
	return filepath.Join(dir, "img", "up.png"), filepath.Join(dir, "img", "down.png")
}

func getDefaultConfigPath() string {
	dir := getExecutableDir()
	return filepath.Join(dir, "config.yaml")
}

func getExecutableDir() string {
	p, _ := os.Executable()
	absPath, _ := filepath.Abs(p)
	return filepath.Dir(absPath)
}

func readBody(resp *http.Response) (body string) {
	res, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return string(res)
}
