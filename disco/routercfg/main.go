package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

const (
    cfgTemplate = "bird.conf"
    cfgSource = "bird.json"
)

type BirdConf struct {
    RouterID string `json:"router_id"`
    BGPLocalAs string `json:"bgp_local_as"`
    BGPNeightbourHost string `json:"bgp_neightbour_host"`
    BGPNeightbourAs string `json:"bgp_neightbour_as"`
    BGPSourceAddr string `json:"bgp_source_addr"`
}

func main() {
    bconf := readConfigFile()
    tmpl, err := template.New(cfgTemplate).ParseFiles(cfgTemplate)
    if err != nil {
        log.Fatalf("Failed to parse template file: %v\n", err)
    }

    err = tmpl.Execute(os.Stdout, bconf)
    if err != nil {
        log.Fatalf("Failed to execute template: %v\n", err)
    }
}

func readConfigFile() BirdConf {
    bconf := BirdConf{}
    content, err := os.ReadFile(cfgSource)
    if err != nil {
        log.Fatalf("Failed to read config file: %v\n", err)
    }

    err = json.Unmarshal(content, &bconf)
    if err != nil {
        log.Fatalf("Failed to parse config file content: %v\n", err)
    }
    return bconf
}
