package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	TEST_JOB_FILE_PATH string = "./test-job-jenkins.xml"
	JOB_NAME           string = "envJob"
	JOB_CREATION_URL   string = "/createItem?name=" + JOB_NAME
	JOB_URL            string = "/job/" + JOB_NAME
)

func iHaveJobConfig() error {
	return nil
}

func iCreateTheJob() error {
	file, err := os.Open(TEST_JOB_FILE_PATH)
	if err != nil {
		return fmt.Errorf("Failed to open file: %s.", TEST_JOB_FILE_PATH)
	}

	req, err := http.NewRequest("POST", getUrl(JOB_CREATION_URL), file)
	getNewJenkinsCrumb()
	req.Header.Add("Jenkins-Crumb", crumb.Crumb)
	req.Header.Add("Content-Type", "text/xml")

	resp, err := httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("Failed to upload job from %s to %s: %s", TEST_JOB_FILE_PATH, getUrl(JOB_CREATION_URL), err)
	}

	body, _ := getBodyString(resp)
	if body != "" {
		return fmt.Errorf("Body is not empty - does job already exist? %s", body)
	}

	return nil
}

func theJobIsCreated() error {
	jobsResp, err := httpClient.Get(getUrl(JOB_URL))

	if err != nil {
		return fmt.Errorf("Failed to create job %s.", JOB_NAME)
	}

	if jobsResp.StatusCode != 200 {
		return fmt.Errorf("Failed to create job %s with status code: %d.", JOB_NAME, jobsResp.StatusCode)
	}

	return nil
}
