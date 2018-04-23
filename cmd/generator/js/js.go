package js

import (
	"encoding/json"
	"errors"
	"github.com/Nordgedanken/matrix_dsl/cmd/lexer"
	"io/ioutil"
	"os"
	"os/exec"
)

func GenerateBot(mx *lexer.Section) error {

	err := os.MkdirAll("js_project", os.ModeDir)
	if err != nil {
		return err
	}

	packageJson := &pkgJson{}
	packageJson.Name, err = getName(mx.Properties)
	if err != nil {
		return err
	}

	packageJson.Description, err = getDesc(mx.Properties)
	if err != nil {
		return err
	}

	packageJson.Main = "index.js"

	packageJsonM, err := json.Marshal(packageJson)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./js_project/package.json", packageJsonM, 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command("npm", "install", "--save", "matrix-bot-sdk")
	cmd.Dir = "./js_project"
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func getName(p []*lexer.PropertyArrays) (string, error) {
	var name string
	for _, v := range p {
		if v.Key == "Name" {
			if v.Value != nil {
				name = *v.Value.String
			} else {
				return "", errors.New("value of the \"Name\" Key is empty. Please fix the Value to contain a string")
			}
			break
		}
	}
	if name == "" {
		return "", errors.New("value of \"Name\" is empty. Please fix the Value to contain a string")
	}

	return name, nil
}

func getDesc(p []*lexer.PropertyArrays) (string, error) {
	var desc string
	for _, v := range p {
		if v.Key == "Description" {
			if v.Value != nil {
				desc = *v.Value.String
			} else {
				return "", errors.New("value of the \"Description\" Key is empty. Please fix the Value to contain a string")
			}
			break
		}
	}
	if desc == "" {
		return "", errors.New("value of \"Description\" is empty. Please fix the Value to contain a string")
	}

	return desc, nil
}
