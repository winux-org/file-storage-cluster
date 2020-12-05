package main

import "gopkg.in/yaml.v2"
import "io/ioutil"
import "log"


type conf struct {
    Env string `yaml:"env"`
}

func loadConf() *conf {
	var c *conf
    yamlFile, err := ioutil.ReadFile("conf.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}


var Conf conf



func init() {
	//Conf.SetEnv("sd")
	//log.Fatalf("Config: %v", Conf.Env)
	// var f = loadConf()
	// log.Fatalf("Config: %v", f.Env)
}