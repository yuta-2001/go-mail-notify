package linehelper

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SendNoCommitNotify() {
	fmt.Println("send no commit notify")

	accessToken := os.Getenv("LINE_NOTIFY_TOKEN")
	fmt.Println(accessToken)

	apiUrl := "https://notify-api.line.me/api/notify"

	form := url.Values{}
	form.Set("message", "草が生えてないよ！やばいよ！")
	form.Set("stickerPackageId", "6136")
	form.Set("stickerId", "10551382")

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
}
