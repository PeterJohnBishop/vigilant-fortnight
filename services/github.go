package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetRepositories() (string, error) {
	token := os.Getenv("GITHUB_API_KEY")

	req, err := http.NewRequest("GET", "https://api.github.com/repositories", nil)
	if err != nil {
		return err.Error(), err
	}

	auth := fmt.Sprintf("Bearer %s", token)

	req.Header.Add("Authorization", auth)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err.Error(), err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error(), err
	}

	return string(body), nil
}
