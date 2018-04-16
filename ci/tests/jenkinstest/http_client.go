package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const (
	LOGIN_FORM_URL string = "/j_acegi_security_check"
	CRUMB_URL string = "/crumbIssuer/api/json"
)

var requestUrl string
var body string
var pluginsResp string
var cookieJar *cookiejar.Jar
var crumb JenkinsCrumb

type JenkinsCrumb struct {
	Crumb             string `json:"crumb"`
	CrumbRequestField string `json:"crumbRequestField"`
}

func getUrl(path string) string {
	return jenkinsHostUrl + path
}

func createNewCookieJar() {
	cookieJar, _ = cookiejar.New(&cookiejar.Options{})
}

func createNewHttpClient() {
	createNewCookieJar()

	httpClient = &http.Client{
		Jar: cookieJar,
	}

	getNewJenkinsCrumb()	
}

func getBodyString(resp *http.Response)(string, error) {
	defer resp.Body.Close()
	body_bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "Error", fmt.Errorf("%s", err)
	}

	return string(body_bytes), nil
}

func getNewJenkinsCrumb() error {
	u := getUrl(CRUMB_URL)
	resp, err := httpClient.Get(u)

	if err != nil {
		return fmt.Errorf("expected response from crumbIssuer, got: %s", body)
	}

	defer resp.Body.Close()

	body_bytes, _ := ioutil.ReadAll(resp.Body)

	if ! strings.Contains(body, `{"_class":"hudson.security.csrf.DefaultCrumbIssuer","crumb":`) {
		return fmt.Errorf("expected %s to contain a crumb", body)
	}

	json.Unmarshal(body_bytes, &crumb)

	return nil
}

func jenkinsLogin(username, password string)(string, error) {
	if crumb.Crumb == "" {
		getNewJenkinsCrumb()
	}

	u := getUrl(LOGIN_FORM_URL)

	resp, err := httpClient.PostForm(u,
		url.Values{"j_username": {username}, "j_password": {password}, "Jenkins-Crumb": {crumb.Crumb}})

	if err != nil {
		return "Error", fmt.Errorf("%s", err)
	}

	body, err := getBodyString(resp)
	return body, err
}
