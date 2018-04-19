package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	MAVEN_JOB_FILE_PATH string = "./test-maven-job.xml"
	MAVEN_JOB_NAME      string = "mavenJob"
	NEXUS_ARTEFACT_URL  string = "https://nexus.dev-build-create.build.finkit.io/repository/releases-libs/com/example/demo/0.0.1/demo-0.0.1.jar"
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
	nexusResp, err := httpClient.Get(NEXUS_ARTEFACT_URL)

	fmt.Printf("nexus response from %s to %s .", NEXUS_ARTEFACT_URL, err)

	if nexusResp.StatusCode == http.StatusOK {
		return fmt.Errorf("returned 200 unexpectedly")
	}

	time.Sleep(40 * time.Second)

	nexusResp, err = httpClient.Get(NEXUS_ARTEFACT_URL)

	if err != nil {
		return fmt.Errorf("Failed to retrieve artefact from %s.", NEXUS_ARTEFACT_URL)
	}

	if nexusResp.StatusCode != http.StatusOK {
		return fmt.Errorf("theJobArtefactIsStoredInNexus status code returned %d not 200", nexusResp.StatusCode)
	}

	return nil
}
