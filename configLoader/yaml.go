package configLoader

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func Load(path string, out interface{}){

	yamlFile, readErr := ioutil.ReadFile(path)

	if readErr != nil {
		log.Panic("Config File Loaded Failed !", readErr)
		return
	}
	err := yaml.Unmarshal(yamlFile, out)

	if err != nil {
		log.Panic("Config Setup Failed !", err)
		return
	}

	fmt.Println("Load config Success.")
}