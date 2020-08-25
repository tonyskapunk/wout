package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type exercise struct {
	Name        string
	Description string
	URL         string
}

func main() {
	var exercises []exercise

	data, err := ioutil.ReadFile("./exercises.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(data), &exercises)
	if err != nil {
		log.Fatal(err)
	}

	i := pick(len(exercises))
	e := exercises[i]

	fmt.Printf("%v (%v)\n", e.Name, e.URL)
}

// pick generates a random int number up to n with a random seed
func pick(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}
