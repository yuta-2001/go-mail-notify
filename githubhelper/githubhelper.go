package githubhelper


import (
	"os"
	"time"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type GraphQLRequest struct {
    Query     string                 `json:"query"`
    Variables map[string]interface{} `json:"variables,omitempty"`
}

type Response struct {
	Data struct {
		User struct {
			ContributionCollection struct {
				ContributionCalendar struct {
					TotalContributions int `json:"totalContributions"`
					Weeks []struct {
						ContributionDays []struct {
							ContributionCount int `json:"contributionCount"`
							Date string `json:"date"`
						} `json:"contributionDays"`
					} `json:"weeks"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}


func GetContributesCount() int {
	accessToken := os.Getenv("GITHUB_TOKEN")
	// JSTのロケーションを取得
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	// 現在の日付と時刻をJSTで取得
	nowJST := time.Now().In(loc)

	// 現在の日付の0時0分0秒をJSTで取得
	startOfTodayJST := time.Date(nowJST.Year(), nowJST.Month(), nowJST.Day(), 0, 0, 0, 0, loc)

	// 現在の日付の23時59分59秒をJSTで取得
	endOfTodayJST := startOfTodayJST.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	// JSTをUTCに変換
	startOfTodayUTC := startOfTodayJST.UTC()
	endOfTodayUTC := endOfTodayJST.UTC()

	// UTCの時間をISO8601フォーマットの文字列に変換
	startOfTodayStr := startOfTodayUTC.Format(time.RFC3339)
	endOfTodayStr := endOfTodayUTC.Format(time.RFC3339)

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
		"from":     startOfTodayStr, // JSTの2024年1月3日の開始をUTCで表現
		"to":       endOfTodayStr, // JSTの2024年1月3日の終了をUTCで表現
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

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	// レスポンスからコントリビュート数を取得
	var contributesCount int
	for _, week := range response.Data.User.ContributionCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			contributesCount += day.ContributionCount
		}
	}

	return contributesCount
}
