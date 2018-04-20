package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	MAVEN_JOB_FILE_PATH            string        = "./test-maven-job.xml"
	MAVEN_JOB_NAME                 string        = "mavenTestJob"
	MAVEN_JOB_WAIT_TIME_IN_SECONDS time.Duration = 40
	NEXUS_ARTEFACT_PATH            string        = "/repository/releases-libs/com/example/demo/0.0.1/demo-0.0.1.jar"
)

func getNexusUrl(path string) string {
	return nexusHostUrl + path
}

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
	nexusResp, err := httpClient.Get(getNexusUrl(NEXUS_ARTEFACT_PATH))

	if err != nil {
		return fmt.Errorf("Unexpected error pinging %s prior to creation of test Maven job", getNexusUrl(NEXUS_ARTEFACT_PATH))
	}

	if nexusResp.StatusCode != http.StatusNotFound {
		return fmt.Errorf("Nexus returned OK response to test Maven job prior to creation")
	}

	time.Sleep(MAVEN_JOB_WAIT_TIME_IN_SECONDS * time.Second)

	nexusResp, err = httpClient.Get(getNexusUrl(NEXUS_ARTEFACT_PATH))

	if err != nil {
		return fmt.Errorf("Failed to retrieve artefact from %s.", getNexusUrl(NEXUS_ARTEFACT_PATH))
	}

	if nexusResp.StatusCode != http.StatusOK {
		return fmt.Errorf("theJobArtefactIsStoredInNexus status code returned %d not 200", nexusResp.StatusCode)
	}

	return nil
}
