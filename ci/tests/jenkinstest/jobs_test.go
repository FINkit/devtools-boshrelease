package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	TEST_JOB_FILE_PATH             string        = "./test-job-jenkins.xml"
	JOB_CREATION_URL               string        = "/createItem?name="
	JOB_URL                        string        = "/job/"
	BUILD_URL                      string        = "/build"
	INITIAL_JOB_NAME               string        = "initialJob"
	SLAVE_JOB_NAME                 string        = "slaveJob"
	SLAVE_JOB_WAIT_TIME_IN_SECONDS time.Duration = 10
)

func getJobUrl(jobName string) string {
	return getJenkinsUrl(JOB_URL + jobName)
}

func getJobBuildUrl(jobName string) string {
	return getJenkinsUrl(JOB_URL + jobName + BUILD_URL)
}

func getJobCreateUrl(jobName string) string {
	return getJenkinsUrl(JOB_CREATION_URL + jobName)
}

func iHaveJobConfig() error {
	return nil
}

func createJenkinsJob(jobName string, filePath string) string {
	file, err := os.Open(filePath)

	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s.", filePath))
	}

	req, err := http.NewRequest("POST", getJobCreateUrl(jobName), file)
	getNewJenkinsCrumb()
	req.Header.Add(crumb.CrumbRequestField, crumb.Crumb)
	req.Header.Add("Content-Type", "text/xml")

	resp, err := httpClient.Do(req)

	if err != nil {
		panic(fmt.Errorf("Failed to upload job from %s to %s: %s", filePath, getJobCreateUrl(jobName), err))
	}

	body, _ := getBodyString(resp)

	return body
}

func iCreateTheJob() error {
	body := createJenkinsJob(INITIAL_JOB_NAME, TEST_JOB_FILE_PATH)

	if body != "" {
		return fmt.Errorf("Body is not empty - does job already exist? %s", body)
	}

	return nil
}

func iCreateTheSlaveJob() error {
	body := createJenkinsJob(SLAVE_JOB_NAME, TEST_JOB_FILE_PATH)

	if body != "" {
		return fmt.Errorf("Body is not empty - does job already exist? %s", body)
	}

	return nil
}

func theJobIsCreated() error {
	jobsResp, err := httpClient.Get(getJobUrl(INITIAL_JOB_NAME))

	if err != nil {
		return fmt.Errorf("Failed to create job %s.", INITIAL_JOB_NAME)
	}

	if jobsResp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create job %s with status code: %d.", INITIAL_JOB_NAME, jobsResp.StatusCode)
	}

	return nil
}

func iExecuteTheJob() error {
	req, err := http.NewRequest("POST", getJobBuildUrl(SLAVE_JOB_NAME), nil)
	getNewJenkinsCrumb()
	req.Header.Add(crumb.CrumbRequestField, crumb.Crumb)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("iExecuteTheJob error: %s", err)
	}

	body, _ := getBodyString(resp)
	fmt.Printf("%s", body)

	return nil
}

func theJobIsExecuted() error {
	time.Sleep(SLAVE_JOB_WAIT_TIME_IN_SECONDS * time.Second)

	jobResp, err := httpClient.Get(getJenkinsUrl(JOB_URL + SLAVE_JOB_NAME + "/1/api/json"))

	if err != nil {
		return fmt.Errorf("Failed to create job %s.", INITIAL_JOB_NAME)
	}

	if jobResp.StatusCode != http.StatusOK {
		return fmt.Errorf("theJobIsExecuted status code returned %d not 200", jobResp.StatusCode)
	}

	body, _ := getBodyString(jobResp)

	if !strings.Contains(body, `"result":"SUCCESS"`) {
		return fmt.Errorf("theJobIsExecuted failed: %s", body)
	}

	return nil
}
