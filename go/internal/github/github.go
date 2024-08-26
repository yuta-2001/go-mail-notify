package github

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "time"
)

type GraphQLRequest struct {
    Query     string                 `json:"query"`
    Variables map[string]interface{} `json:"variables,omitempty"`
}

type Response struct {
    Data struct {
        User struct {
            ContributionCollection struct {
                TotalCommitContributions int `json:"totalCommitContributions"`
            } `json:"contributionCollection"`
        } `json:"user"`
    } `json:"data"`
}

func GetContributesCount(username string, token string) (int, error) {
    // 日本標準時（JST）を取得
    nowJST := time.Now().In(time.FixedZone("JST", 9*60*60))
    // その日の開始時刻（JST）
    startOfTodayJST := time.Date(nowJST.Year(), nowJST.Month(), nowJST.Day(), 0, 0, 0, 0, nowJST.Location())
    // その日の終了時刻（JST）
    endOfTodayJST := startOfTodayJST.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

    // UTCに変換してISO 8601形式にフォーマット
    startOfTodayStr := startOfTodayJST.UTC().Format(time.RFC3339)
    endOfTodayStr := endOfTodayJST.UTC().Format(time.RFC3339)

    url := "https://api.github.com/graphql"
    query := `
        query($userName: String!, $from: DateTime!, $to: DateTime!) {
            user(login: $userName) {
                contributionsCollection(from: $from, to: $to) {
                    totalCommitContributions
                }
            }
        }
    `
    variables := map[string]interface{}{
        "userName": username,
        "from":     startOfTodayStr,
        "to":       endOfTodayStr,
    }

    graphqlRequest := GraphQLRequest{
        Query:     query,
        Variables: variables,
    }

    requestBody, err := json.Marshal(graphqlRequest)
    if err != nil {
        return 0, err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        return 0, err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    var response Response
    if err := json.Unmarshal(body, &response); err != nil {
        return 0, err
    }

    contributesCount := response.Data.User.ContributionCollection.TotalCommitContributions

    return contributesCount, nil
}
