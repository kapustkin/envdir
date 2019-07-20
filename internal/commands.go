package internal

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// Run основной обработчик
func Run(args []string) error {
	envDir, programmPath, err := getInputValues(args)
	if err != nil {
		return err
	}
	env, err := getEnviroment(envDir)
	if err != nil {
		return err
	}

	err = start(programmPath, env)
	if err != nil {
		return err
	}

	return nil
}

func start(command string, env []string) error {
	cmd := exec.Command(command)
	cmd.Env = env
	err := cmd.Start()
	if err != nil {
		return err
	}

	fmt.Printf("%s %v", command, env)
	return nil
}

func getEnviroment(envDir string) ([]string, error) {
	files, err := ioutil.ReadDir(envDir)
	if err != nil {
		return []string{}, err
	}
	result := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		param, err := getEnvParametr(path.Join(envDir, file.Name()))
		if err != nil {
			return []string{}, err
		}
		result = append(result, param)
	}
	return result, nil
}

func getEnvParametr(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	result.WriteString(fileNameWithoutExtension(fileName))
	result.WriteString("=")
	result.WriteString(string(data))
	return result.String(), nil
}

func fileNameWithoutExtension(fn string) string {
	name := filepath.Base(fn)
	return strings.TrimSuffix(name, path.Ext(name))
}

func getInputValues(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", fmt.Errorf("для запуска требуется передать 2 аргумента")
	}

	return args[0], args[1], nil
}
