package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Release represents a GitHub release
type Release struct {
	TagName   string `json:"tag_name"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"` // Add CreatedAt field to capture the release creation date
}

// fetchGithubReleases fetches the last 20 releases of a GitHub project
func fetchGithubReleases(repoURL string, limit int) ([]Release, error) {
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
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Fatal("close file: ", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response body
	var releases []Release
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, fmt.Errorf("failed to parse releases response: %w", err)
	}

	if len(releases) > limit {
		releases = releases[:limit]
	}

	return releases, nil
}

// FileContent represents the response from the GitHub API for file contents
type FileContent struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

// fetchFileFromGitHub fetches the contents of a specific file from a GitHub repository at a given tag
func fetchFileFromGitHub(repoURL, tag, filePath string) string {
	// Extract the owner and repo name from the URL
	parts := strings.Split(strings.TrimPrefix(repoURL, "https://github.com/"), "/")
	if len(parts) < 2 {
		return fmt.Errorf("invalid GitHub repository URL").Error()
	}
	owner, repo := parts[0], parts[1]

	// Construct the GitHub API URL for the file
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s", owner, repo, filePath, tag)

	// Make the HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("failed to fetch file: %v", err)
		return ""
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Printf("close file: %v", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
		return ""
	}

	// Parse the response body
	var fileContent FileContent
	err = json.NewDecoder(resp.Body).Decode(&fileContent)
	if err != nil {
		log.Printf("failed to parse file content response: %v", err)
		return ""
	}

	// Decode the file content if it's base64 encoded
	if fileContent.Encoding == "base64" {
		decodedContent, err := base64.StdEncoding.DecodeString(fileContent.Content)
		if err != nil {
			log.Printf("failed to decode file content: %v", err)
			return ""

		}
		return string(decodedContent)
	}

	return fileContent.Content
}

// generateReleaseNotesURL generates a URL to view the release notes for a specific tag in a GitHub project.
func generateReleaseNotesURL(repoURL, tag string) string {
	// Extract the owner and repo name from the URL
	parts := strings.Split(strings.TrimPrefix(repoURL, "https://github.com/"), "/")
	if len(parts) < 2 {
		log.Printf("invalid GitHub repository URL")
		return ""
	}
	owner, repo := parts[0], parts[1]

	// Construct the URL for the release notes
	releaseNotesURL := fmt.Sprintf("https://github.com/%s/%s/releases/tag/%s", owner, repo, tag)
	return releaseNotesURL
}
