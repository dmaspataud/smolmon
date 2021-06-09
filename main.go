package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

type YamlTarget struct {
	Name    string
	Address string
}

type YamlConfig struct {
	Targets []YamlTarget
}

func main() {
	c := YamlConfig{}
	var wg sync.WaitGroup
	err := yaml.Unmarshal(readConf(), &c)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < len(c.Targets); i++ {
		if c.Targets[i].Address != "" {
			wg.Add(1)
			checkHealth(c.Targets[i], &wg)
		}
	}
}

func readConf() []byte {
	conf, err := ioutil.ReadFile("./targets.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

func checkHealth(target YamlTarget, wg *sync.WaitGroup) {
	defer wg.Done()
	checkSpinner, _ := pterm.DefaultSpinner.Start("Checking ...", target.Name)
	resp, err := http.Get(target.Address)
	if err != nil {
		checkSpinner.Fail(target.Name, " (", target.Address, ") => ", err)
		return
	}
	if resp.StatusCode != 200 {
		checkSpinner.Fail(target.Name, " (", target.Address, ") => ", resp.Status)
	} else if resp.StatusCode == 200 {
		checkSpinner.Success(target.Name, " (", target.Address, ")")
	}
}
