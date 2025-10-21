package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"vigilant-fortnight/models"
)

func GetRepositories(username string) ([]models.Repository, error) {
	token := os.Getenv("GITHUB_API_KEY")
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("GitHub API error: %s", string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var repos []models.Repository
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}
