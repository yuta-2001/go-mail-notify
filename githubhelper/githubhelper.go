package githubhelper


import (
	"os"
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type GraphQLRequest struct {
    Query     string                 `json:"query"`
    Variables map[string]interface{} `json:"variables,omitempty"`
}

// get contributes count from github
func GetContributesCount() int {
	accessToken := os.Getenv("GITHUB_TOKEN")

	url := "https://api.github.com/graphql"
	query := `
		query($userName: String!, $from: DateTime!, $to: DateTime!) {
			user(login: $userName) {
				contributionsCollection(from: $from, to: $to) {
					contributionCalendar {
						totalContributions
						weeks {
							contributionDays {
								contributionCount
								date
							}
						}
					}
				}
			}
		}
	`
	// クエリの変数
	variables := map[string]interface{}{
		"userName": os.Getenv("GITHUB_USER"),
		"from":     "2024-01-03T00:00:00Z",
		"to":       "2024-01-03T11:59:59Z",
	}

	// GraphQLリクエストオブジェクトを作成
	graphqlRequest := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	// GraphQLリクエストオブジェクトをJSONに変換
	requestBody, err := json.Marshal(graphqlRequest)
    if err != nil {
        panic(err)
    }

    // POSTリクエストを作成
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        panic(err)
    }

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + accessToken)

    // HTTPクライアントを作成し、リクエストを実行
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	// レスポンスを読み込み
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	return 0
}
