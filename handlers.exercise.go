package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	exercises := getAllExercises()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": exercises}, "index.html")
}

func getRandomExercise(c *gin.Context) {
	exercises := getAllExercises()

	i := randomPick(len(exercises))
	e := exercises[i]

	render(c, gin.H{
		"payload": e}, "random.html")
}

// pick generates a random int number up to n with a random seed
func randomPick(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}
