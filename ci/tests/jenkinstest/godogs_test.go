package main

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

const (
	JENKINS_HOST string = "JENKINS_HOST"
)

var jenkinsHostUrl string = os.Getenv(JENKINS_HOST)
var jenkinsLogin string = jenkinsHostUrl + "/login"
var requestUrl string
var body string
var pluginsResp string
var cookieJar *cookiejar.Jar
var crumb JenkinsCrumb
var httpClient *http.Client

type JenkinsCrumb struct {
	Crumb             string `json:"crumb"`
	CrumbRequestField string `json: "crumbRequestField"`
}

func init() {
	createNewHttpClient()
}

func createNewCookieJar() {
	cookieJar, _ = cookiejar.New(&cookiejar.Options{})
}

func createNewHttpClient() {
	createNewCookieJar()

	httpClient = &http.Client{
		Jar: cookieJar,
	}
}

func thereIsAJenkinsInstall() error {
	requestUrl = jenkinsHostUrl + "/login"
	return nil
}

func getBodyString(resp *http.Response)(string, error) {
	defer resp.Body.Close()
	body_bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	return string(body_bytes), nil
}

func iAccessTheLoginScreen() error {
	resp, err := http.Get(requestUrl)
	if err != nil {
		return err
	}

	body, _ = getBodyString(resp)
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
	pluginsResp, err := httpClient.Get(u)

	if err != nil {
		return err
	}
	body, _ = getBodyString(pluginsResp)
	return nil
}

func allThePluginsAreInstalled() error {
	if !strings.Contains(body, "<shortName>cucumber-reports</shortName>") {
		return fmt.Errorf("expected %s to contain 'cucumber-reports'", body)
	}
	return nil
}

func getNewJenkinsCrumb() error {
	u := jenkinsHostUrl + "/crumbIssuer/api/json"
	resp, err := httpClient.Get(u)

	if err != nil {
		return fmt.Errorf("expected response from crumbIssuer, got: %s", body)
	}

	defer resp.Body.Close()

	body_bytes, _ := ioutil.ReadAll(resp.Body)

	if ! strings.Contains(body, `{"_class":"hudson.security.csrf.DefaultCrumbIssuer","crumb":`) {
		return fmt.Errorf("expected %s to contain '/logout' link", body)
	}

	json.Unmarshal(body_bytes, &crumb)

	return nil
}

func iHaveLoggedIntoJenkins() error {

	getNewJenkinsCrumb()

	loginUrl := jenkinsHostUrl + "/j_acegi_security_check"
	jenkinsPassword := os.Getenv("JENKINS_PASSWORD")

	resp, err := httpClient.PostForm(loginUrl,
		url.Values{"j_username": {"administrator"}, "j_password": {jenkinsPassword}, "Jenkins-Crumb": {crumb.Crumb}})

	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if !strings.Contains(string(body), `<a href="/logout"><b>log out</b></a>`) {
		return fmt.Errorf("expected %s to contain '/logout' link", body)
	}

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