package main

import (
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	RabbitMQ struct {
		Host     string `use_yaml:"host"`
		Port     int64  `use_yaml:"port"`
		Username string `use_yaml:"username"`
		Password string `use_yaml:"password"`
	}
}

func readConf(filename string) (*Conf, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Conf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return &conf, nil
}

func main() {
	conf, err := readConf("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("%v", conf)
}
