package main

import (
	"strings"
	"fmt"
	"os"
	"net/http"
    "net/url"
	"io/ioutil"
	"github.com/DATA-DOG/godog"
    "encoding/json"
)

const (
    JENKINS_HOST string = "JENKINS_HOST"
)

var jenkinsHostUrl string = os.Getenv(JENKINS_HOST)
var jenkinsLogin string = jenkinsHostUrl + "/login"
var requestUrl string
var body string
var pluginsResp string

func thereIsAJenkinsInstall() error {
    requestUrl = jenkinsHostUrl + "/login"
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
	resp, err := http.Get(requestUrl)
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
	u := jenkinsHostUrl + "/pluginManager/api/xml?depth=1"
	pluginsResp, err := http.Get(u)
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

type JenkinsCrumb struct {
    Crumb string `json:"crumb"`
    CrumbRequestField string `json: "crumbRequestField"`
}

var crumb JenkinsCrumb

func getJenkinsCrumb() {
    if crumb.Crumb == "" {
        u := jenkinsHostUrl + "/crumbIssuer/api/json"
        resp, err := http.Get(u)

        defer resp.Body.Close()

        if err != nil {
            fmt.Sprintf("%s", err)
        }

        body_bytes, _ := ioutil.ReadAll(resp.Body)

        json.Unmarshal(body_bytes, &crumb)
    }
}

func iHaveLoggedIntoJenkins() error {
    getJenkinsCrumb()

    loginUrl := jenkinsHostUrl + "/j_acegi_security_check"
	jenkinsPassword := os.Getenv("JENKINS_PASSWORD")

    resp, err := http.PostForm(loginUrl,
    url.Values{"j_username": {"admin"}, "j_password": {jenkinsPassword}, "Jenkins-Crumb": {crumb.Crumb}})

    if err != nil {
        fmt.Printf("%s", err)
    }

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

    return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a jenkins install$`, thereIsAJenkinsInstall)

	s.Step(`^I have logged into Jenkins$`, iHaveLoggedIntoJenkins)

    s.Step(`^I access the login screen$`, iAccessTheLoginScreen)
	s.Step(`^jenkins should be unlocked$`, jenkinsShouldBeUnlocked)

    s.Step(`^I access plugin management$`, iAccessPluginManagement)
	s.Step(`^all the plugins are installed$`, allThePluginsAreInstalled)
}

