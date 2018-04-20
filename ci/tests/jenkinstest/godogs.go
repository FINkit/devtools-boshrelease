package main

import (
	"net/http"
	"os"
)

const (
	JENKINS_HOST string = "JENKINS_HOST"
	NEXUS_HOST   string = "NEXUS_HOST"
)

var jenkinsHostUrl string
var nexusHostUrl string
var httpClient *http.Client

func init() {
	jenkinsHostUrl = os.Getenv(JENKINS_HOST)

	if jenkinsHostUrl == "" {
		panic("JENKINS_HOST is empty")
	}

	nexusHostUrl = os.Getenv(NEXUS_HOST)

	if nexusHostUrl == "" {
		panic("NEXUS_HOST is empty")
	}
}

func main() {

}
