package linehelper

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"strconv"
	"io/ioutil"
)

func SendMessage(contributesCount int) {
	fmt.Println("send no commit notify")

	accessToken := os.Getenv("LINE_NOTIFY_TOKEN")
	fmt.Println(accessToken)

	apiUrl := "https://notify-api.line.me/api/notify"

	form := url.Values{}

	if contributesCount == 0 {
		form.Set("message", "草が生えてないよ！やばいよ！")
		form.Set("stickerPackageId", "6136")
		form.Set("stickerId", "10551382")
	} else {
		form.Set("message", "草が生えてるよ！本日のコミット数は" + strconv.Itoa(contributesCount) + "だよ！")
		form.Set("stickerPackageId", "446")
		form.Set("stickerId", "1989")
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }
    fmt.Println("Response Body:", string(body))
}
