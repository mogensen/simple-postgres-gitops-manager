package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// State is a representation of a desired or current state
type State struct {
	Databases []Database `yaml:"databases"`
}

// Database ...
type Database struct {
	Name string `yaml:"name"`
}

func newStateFromFile(fileName string) (*State, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	state := &State{}
	err = yaml.Unmarshal(bytes, state)
	if err != nil {
		return nil, err
	}
	return state, nil

}
