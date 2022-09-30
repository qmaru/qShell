package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ListScripts() ([]string, error) {
	root, err := LoadScripts()
	if err != nil {
		return nil, err
	}

	dirs, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	scripts := make([]string, 0)
	index := 0
	for _, dir := range dirs {
		name := dir.Name()
		if !dir.IsDir() && name != "README.md" {
			sFile := filepath.Join(root, name)
			sText, err := os.ReadFile(sFile)
			if err != nil {
				return nil, err
			}
			describe := readDescribe(sText)
			if describe != "" {
				index++
				tips := fmt.Sprintf("%d. %s - %s", index, name, describe)
				scripts = append(scripts, tips)
			}
		}
	}
	return scripts, nil
}

func RunScript(name string, paras []string) (string, error) {
	root, err := LoadScripts()
	if err != nil {
		return "", err
	}

	script := filepath.Join(root, name)
	cmd := exec.Command(script, paras...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	results := strings.TrimSpace(string(output))
	return results, nil
}
