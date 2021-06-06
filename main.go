package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	address := "https://www.google.com"
	resp, err := http.Get(address)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.StatusCode)
}

func readConf() {

}

func checkHealth(address string) {

}
