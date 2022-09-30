package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"

	"github.com/qmaru/minitools"
)

func readDescribe(text []byte) string {
	rule := regexp.MustCompile(`# describe:(.*)`)
	data := rule.FindAllSubmatch(text, -1)
	if len(data) != 0 {
		describe := string(data[0][1])
		return describe
	}
	return ""
}

func loadRoot(subpath ...string) (string, error) {
	f := new(minitools.FileSuiteBasic)
	folder, err := f.RootPath(subpath...)
	if err != nil {
		return "", err
	}
	folderPath, err := f.Mkdir(folder)
	if err != nil {
		return "", err
	}
	return folderPath, nil
}

func LoadScripts() (string, error) {
	return loadRoot("scripts")
}

func LoadFile() (string, error) {
	return loadRoot("scripts", "files")
}

func loadUserfile() (string, error) {
	root, err := loadRoot()
	if err != nil {
		return "", err
	}
	userCfg := filepath.Join(root, "users.json")
	return userCfg, nil
}

func LoadUserConfig() (map[string]string, error) {
	userCfg, err := loadUserfile()
	if err != nil {
		return nil, err
	}
	userRaw, err := os.ReadFile(userCfg)
	if err != nil {
		return nil, err
	}

	var userJson map[string]string
	err = json.Unmarshal(userRaw, &userJson)
	if err != nil {
		return nil, err
	}
	return userJson, nil
}
