package config

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type Configuration struct {
	File    string
	Mapping map[string]string
}

func New(file string) (conf Configuration) {
	conf = Configuration{
		File:    file,
		Mapping: make(map[string]string),
	}
	return
}

func Load(file string) (conf Configuration, err error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	conf = Configuration{
		File:    file,
		Mapping: make(map[string]string),
	}
	for _, v := range strings.Split(string(bytes), "\n") {
		if v == "" {
			continue
		}

		key := strings.SplitN(v, "=", 2)[0]
		value := strings.SplitN(v, "=", 2)[1]

		conf.Mapping[key] = value
	}

	return
}

func (c *Configuration) Save() (err error) {
	var conf string

	for v := range c.Mapping {
		conf += v + "=" + c.Mapping[v] + "\n"
	}

	err = os.WriteFile(c.File, []byte(conf), os.ModePerm)

	return
}

func (c *Configuration) Add(key string, value string) (err error) {
	if _, ok := c.Mapping[key]; ok {
		err = errors.New("configuration already contains spesified key")
		return
	}

	c.Mapping[key] = value
	return
}

func (c *Configuration) Set(key string, newValue string) (err error) {
	if _, ok := c.Mapping[key]; !ok {
		err = errors.New("configuration does not contain the spesified key (try using Add first)")
		return
	}

	c.Mapping[key] = newValue
	return
}

func (c *Configuration) Get(key string) (out string) {
	if _, ok := c.Mapping[key]; !ok {
		out = ""
		return
	}

	out = c.Mapping[key]
	return
}
