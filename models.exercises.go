package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type exercise struct {
	Name        string
	Description string
	URL         string
}

func getAllExercises() []exercise {
	var exercises []exercise

	data, err := ioutil.ReadFile("./exercises.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(data), &exercises)
	if err != nil {
		log.Fatal(err)
	}

	return exercises
}
