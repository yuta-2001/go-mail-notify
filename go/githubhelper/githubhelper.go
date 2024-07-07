package githubhelper


import (
    "os"
    "time"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "fmt"
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
                            ContributionCount int    `json:"contributionCount"`
                            Date              string `json:"date"`
                        } `json:"contributionDays"`
                    } `json:"weeks"`
                } `json:"contributionCalendar"`
            } `json:"contributionCollection"`
        } `json:"user"`
    } `json:"data"`
}


func GetContributesCount() (int, error) {
    accessToken := os.Getenv("GITHUB_TOKEN")
    if accessToken == "" {
        return 0, fmt.Errorf("GITHUB_TOKEN is not set")
    }

    nowUTC := time.Now().UTC()
    startOfTodayUTC := time.Date(nowUTC.Year(), nowUTC.Month(), nowUTC.Day(), 0, 0, 0, 0, time.UTC)
    endOfTodayUTC := startOfTodayUTC.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

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
    variables := map[string]interface{}{
        "userName": os.Getenv("GITHUB_USER"),
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
    req.Header.Set("Authorization", "Bearer " + accessToken)

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

    var contributesCount int
    for _, week := range response.Data.User.ContributionCollection.ContributionCalendar.Weeks {
        for _, day := range week.ContributionDays {
            contributesCount += day.ContributionCount
        }
    }

    return contributesCount, nil
}
