package main

import (
	"strings"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"github.com/DATA-DOG/godog"
)

const JENKINS_HOST string = "JENKINS_HOST"

var jenkinsHostUrl string = os.Getenv(JENKINS_HOST)
var url string
var body string
var pluginsResp string

func thereIsAJenkinsInstall() error {
	url = jenkinsHostUrl + "/login"
	return nil
}

func getBodyString(resp *http.Response) string {
	defer resp.Body.Close()
	body_bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Sprintf("%s", err)
	}

	return string(body_bytes)
}

func iAccessTheLoginScreen() error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

    body = getBodyString(resp)
	return nil
}

func jenkinsShouldBeUnlocked() error {
	if strings.Contains(body, "Unlock Jenkins") {
		return fmt.Errorf("expected %s not to contain 'Unlock Jenkins'", body)
	}
	return nil
}

func iAccessPluginManagement() error {
	url = jenkinsHostUrl + "/pluginManager/api/xml?depth=1"
	pluginsResp, err := http.Get(url)
	if err != nil {
		return err
	}
    body = getBodyString(pluginsResp)
	return nil
}

func allThePluginsAreInstalled() error {
	if ! strings.Contains(body, "<shortName>cucumber-reports</shortName>") {
		return fmt.Errorf("expected %s to contain 'cucumber-reports'", body)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a jenkins install$`, thereIsAJenkinsInstall)

    s.Step(`^I access the login screen$`, iAccessTheLoginScreen)
	s.Step(`^jenkins should be unlocked$`, jenkinsShouldBeUnlocked)

    s.Step(`^I access plugin management$`, iAccessPluginManagement)
	s.Step(`^all the plugins are installed$`, allThePluginsAreInstalled)
}

