package main

import (
	"no-commit-notify/githubhelper"
	"github.com/joho/godotenv"
)

// Load .env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Load .env file
	loadEnv()

	// get contributes count from github
	contributesCount := githubhelper.GetContributesCount()

	println(contributesCount)
}
