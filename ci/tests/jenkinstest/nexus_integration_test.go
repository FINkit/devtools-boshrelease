package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	MAVEN_JOB_FILE_PATH string = "./test-maven-job.xml"
	MAVEN_JOB_NAME      string = "mavenJob"
)

func iHaveAddedAMavenJob() error {
	body := createJenkinsJob(MAVEN_JOB_NAME, MAVEN_JOB_FILE_PATH)

	if body != "" {
		return fmt.Errorf("Body is not empty - does job already exist? %s", body)
	}

	return nil
}

func iExecuteTheMavenJob() error {
	req, err := http.NewRequest("POST", getJobBuildUrl(MAVEN_JOB_NAME), nil)
	getNewJenkinsCrumb()
	req.Header.Add(crumb.CrumbRequestField, crumb.Crumb)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("iExecuteTheMavenJob error: %s", err)
	}

	body, _ := getBodyString(resp)
	fmt.Printf("%s", body)

	return nil
}

func theJobArtefactIsStoredInNexus() error {
	time.Sleep(40 * time.Second)

	jobResp, err := httpClient.Get(getUrl("/job/" + MAVEN_JOB_NAME + "/1/api/json"))

	if err != nil {
		return fmt.Errorf("Failed to create job %s.", MAVEN_JOB_NAME)
	}

	if jobResp.StatusCode != http.StatusOK {
		return fmt.Errorf("theJobArtefactIsStoredInNexus status code returned %d not 200", jobResp.StatusCode)
	}

	body, _ := getBodyString(jobResp)

	if !strings.Contains(body, `"result":"SUCCESS"`) {
		return fmt.Errorf("theJobArtefactIsStoredInNexus failed: %s", body)
	}

	return nil
}
