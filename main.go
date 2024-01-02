package main

import (
	"no-commit-notify/githubhelper"
	"no-commit-notify/linehelper"
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

	if (contributesCount == 0) {
		linehelper.SendNoCommitNotify()
	} else {
		println("commit")
	}
}
