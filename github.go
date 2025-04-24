package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Release represents a GitHub release
type Release struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
}

// fetchGithubReleases fetches the last 20 releases of a GitHub project
func fetchGithubReleases(repoURL string) ([]Release, error) {
	// Extract the owner and repo name from the URL
	parts := strings.Split(strings.TrimPrefix(repoURL, "https://github.com/"), "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid GitHub repository URL")
	}
	owner, repo := parts[0], parts[1]

	// Construct the GitHub API URL for releases
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)

	// Make the HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch releases: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response body
	var releases []Release
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, fmt.Errorf("failed to parse releases response: %w", err)
	}

	// Limit to the last 20 releases
	if len(releases) > 20 {
		releases = releases[:20]
	}

	return releases, nil
}

func findGitHubReleases() {
	repoURL := "https://github.com/codefresh-io/gitops-runtime-helm"

	releases, err := fetchGithubReleases(repoURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Last 20 releases:")
	for _, release := range releases {
		fmt.Printf("Tag: %s, Name: %s\n", release.TagName, release.Name)
	}
}

// FileContent represents the response from the GitHub API for file contents
type FileContent struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

// fetchFileFromGitHub fetches the contents of a specific file from a GitHub repository at a given tag
func fetchFileFromGitHub(repoURL, tag, filePath string) (string, error) {
	// Extract the owner and repo name from the URL
	parts := strings.Split(strings.TrimPrefix(repoURL, "https://github.com/"), "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid GitHub repository URL")
	}
	owner, repo := parts[0], parts[1]

	// Construct the GitHub API URL for the file
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s", owner, repo, filePath, tag)

	// Make the HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	// Parse the response body
	var fileContent FileContent
	err = json.NewDecoder(resp.Body).Decode(&fileContent)
	if err != nil {
		return "", fmt.Errorf("failed to parse file content response: %w", err)
	}

	// Decode the file content if it's base64 encoded
	if fileContent.Encoding == "base64" {
		decodedContent, err := base64.StdEncoding.DecodeString(fileContent.Content)
		if err != nil {
			return "", fmt.Errorf("failed to decode file content: %w", err)
		}
		return string(decodedContent), nil
	}

	return fileContent.Content, nil
}

func readContent() {
	repoURL := "https://github.com/codefresh-io/gitops-runtime-helm"
	tag := "0.18.2"
	filePath := "charts/gitops-runtime/Chart.yaml"

	content, err := fetchFileFromGitHub(repoURL, tag, filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("File Content:")
	fmt.Println(content)
}
