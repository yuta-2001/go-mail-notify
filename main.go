package main

// For testing in local environment
// import (
// 	"no-commit-notify/githubhelper"
// 	"no-commit-notify/linehelper"
// 	"github.com/joho/godotenv"
// )

// // Load .env file
// func loadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func main() {
// 	// Load .env file
// 	loadEnv()

// 	// get contributes count from github
// 	contributesCount := githubhelper.GetContributesCount()

// 	if (contributesCount == 0) {
// 		linehelper.SendNoCommitNotify()
// 	} else {
// 		println("commit")
// 	}
// }




// For AWS Lamda Environment
import (
	"no-commit-notify/githubhelper"
	"no-commit-notify/linehelper"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest is the entry point for AWS Lambda function
func HandleRequest() {
	// get contributes count from github
	contributesCount := githubhelper.GetContributesCount()

	if (contributesCount == 0) {
		linehelper.SendNoCommitNotify()
		println("no commit")
	} else {
		println("commit")
	}
}

func main() {
	lambda.Start(HandleRequest)
}
