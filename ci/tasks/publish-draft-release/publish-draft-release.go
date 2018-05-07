package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	gitHubUrl                      string = "https://github.com/"
	gitHubCredentialsVariable      string = "GITHUB_ACCESS_KEY"
	defaultVersion                 string = "latest"
	defaultBranch                  string = "master"
	minimumNumberOfCommandLineArgs int    = 3
)

var (
	releaseApiUrl string = gitHubUrl + "repos/" + owner + "/" + repo + "/releases/"
	owner         string
	repo          string
	branch        string
	version       string
	description   string
	credentials   string = os.Getenv(gitHubCredentialsVariable)
)

type Release struct {
	TagName    string `json:"tag_name"`
	Branch     string `json:"target_commitish"`
	Name       string `json:"name"`
	Body       string `json:"body"`
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
}

func getReleaseApiUrl() string {
	return releaseApiUrl + version
}

func sendRequest(body io.Reader, bodySize int64) (*http.Response, error) {
	req, err := http.NewRequest("PATCH", getReleaseApiUrl(), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", credentials))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.ContentLength = bodySize

	os.Stdout.WriteString(fmt.Sprintf("Sending request to %s with data %s", getReleaseApiUrl(), body))

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return resp, err
	}

	return resp, err
}

func publishDraftRelease() error {
	release := Release{
		TagName:    version,
		Branch:     branch,
		Name:       version,
		Body:       description,
		Draft:      false,
		Prerelease: false,
	}

	releaseData, err := json.Marshal(release)

	if err != nil {
		return fmt.Errorf("error setting JSON data %s when publishing draft release to %s due to %s", releaseData, getReleaseApiUrl(), err)
	}

	releaseBuffer := bytes.NewBuffer(releaseData)

	resp, err := sendRequest(releaseBuffer, int64(releaseBuffer.Len()))

	if err != nil {
		return fmt.Errorf("error publishing draft release to %s with response %s", getReleaseApiUrl(), resp)
	}

	if resp == nil {
		return fmt.Errorf("error publishing draft release to %s with nil response", getReleaseApiUrl())
	}

	code := resp.StatusCode

	if code != http.StatusOK {
		return fmt.Errorf("error publishing draft release to %s with response code %d", getReleaseApiUrl(), code)
	}

	return nil
}

func main() {
	if credentials == "" {
		os.Stderr.WriteString("Must provide GitHub credentials via GITHUB_CREDENTIALS\n")
		os.Exit(1)
	}

	if len(os.Args) != minimumNumberOfCommandLineArgs {
		os.Stderr.WriteString(fmt.Sprintf("Only found %d arguments - Must provide owner and repo\n", len(os.Args)))
		os.Exit(1)
	}

	owner = os.Args[1]
	repo = os.Args[2]
	version = os.Args[3]
	branch = os.Args[4]
	description = os.Args[5]

	if owner == "" {
		os.Stderr.WriteString("Must provide owner as the first argument\n")
		os.Exit(1)
	}

	if repo == "" {
		os.Stderr.WriteString("Must provide repo as the second argument\n")
		os.Exit(1)
	}

	if version == "" {
		version = defaultVersion
	}

	if branch == "" {
		branch = defaultBranch
	}

	if description == "" {
		description = version
	}

	err := publishDraftRelease()

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Failed to publish draft release - %s\n", err))
	}
}
