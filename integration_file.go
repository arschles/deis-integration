package main

type IntegrationFile struct {
  Version int `yaml:"version"`
  Name string `yaml:"name"`
  Description string `yaml:"description"`
  Commands []string `yaml:"commands"`
  Expect []Expectation `yaml:"expect"`
}

type Expectation struct {
  What string `yaml:"what"`
  Value string `yaml:"value"`
}
