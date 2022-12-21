package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type spec struct {
	Openapi    string                 `yaml:"openapi"`
	Info       map[string]interface{} `yaml:"info"`
	Servers    []interface{}          `yaml:"servers"`
	Security   []interface{}          `yaml:"security"`
	Paths      map[string]interface{} `yaml:"paths"`
	Components map[string]interface{} `yaml:"components"`
	Tags       []interface{}          `yaml:"tags"`
}

func (c *spec) getSpec() *spec {
	yamlFile, err := ioutil.ReadFile("openapi-full.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func (c *spec) writeSpec() *spec {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}

	err = ioutil.WriteFile("openapi.yaml", data, 0644)
	if err != nil {
		log.Printf("yamlFile.Write err   #%v ", err)
	}
	return c
}

func main() {
	var c spec
	c.getSpec()

	allowedPaths := []string{
		"/workspaces",
		"/workspaces/{workspaceId}",
		"/environments",
		"/environments/{environmentId}",
	}
	paths := make(map[string]interface{})
	for _, allowedPath := range allowedPaths {
		paths[allowedPath] = c.Paths[allowedPath]
	}
	c.Paths = paths
	c.writeSpec()
}
