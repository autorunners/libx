package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Protocol  string
	Host      string
	Port      string
	ClientId  string `yaml:"client_id"`
	User      string
	Pass      string
	KeepAlive int64     `yaml:"keep_alive"`
	SecretKey string    `yaml:"secret_key"`
	CertFiles CertFiles `yaml:"cert_files"`
	Topics    []string  `yaml:"topics"`
}

type CertFiles struct {
	CaFile   string
	CertFile string
	KeyFile  string
}


func Parse(configFile string) Config {
	var c Config
	yamlFile, err := ioutil.ReadFile(configFile)
	//log.Println(string(yamlFile))
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Println(err)
	}
	return c
}



