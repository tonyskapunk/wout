// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	// Handle GET random exercise
	router.GET("/rand", getRandomExercise)
}
