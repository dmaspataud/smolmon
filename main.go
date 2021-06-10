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
	pterm.EnableDebugMessages()
	for i := 0; i < len(c.Targets); i++ {
		if c.Targets[i].Address != "" {
			wg.Add(1)
			go checkHealth(c.Targets[i], &wg)
		}
	}
	wg.Wait()
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
	resp, err := http.Get(target.Address)
	if err != nil {
		pterm.Error.Println(target.Name, "(", target.Address, ")")
		return
	}
	if resp.StatusCode != 200 {
		pterm.Error.Println(target.Name, "(", target.Address, ")")
	} else if resp.StatusCode == 200 {
		pterm.Success.Println(target.Name, "(", target.Address, ")")
	}
}
