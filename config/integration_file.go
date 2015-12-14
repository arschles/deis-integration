package config

import "errors"

var (
	validExpectations = map[string]struct{}{
		"exitcode": struct{}{},
	}
	ErrInvalidExpectation = errors.New("invalid expectation")
)

type IntegrationFile struct {
	Version     int       `yaml:"version"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Commands    []Command `yaml:"commands"`
}

type Command struct {
	Shell      string        `yaml:"shell"`
	PostExpect []Expectation `yaml:"post_expectation"`
}

func (c Command) PostExpectations() []Expectation {
	// TODO: if c.PostExpect doesn't have exitcode: 0, add it?
	return c.PostExpect
}

type Expectation struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func (e Expectation) Validate() error {
	if _, ok := validExpectations[e.Name]; !ok {
		return ErrInvalidExpectation
	}
	return nil
}
