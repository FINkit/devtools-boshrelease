package main

import (
	"strings"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"github.com/DATA-DOG/godog"
)

var url string
var body string

func thereIsAJenkinsInstall() error {
	url = os.Getenv("JENKINS_HOST") + "/login"
	return nil
}

func iAccessTheLoginScreen() error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body_bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	body = string(body_bytes)
	return nil
}

func jenkinsShouldBeUnlocked() error {
	if strings.Contains(body, "Unlock Jenkins") {
		return fmt.Errorf("expected %s not to contain 'Unlock Jenkins'", body)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a jenkins install$`, thereIsAJenkinsInstall)
	s.Step(`^I access the login screen$`, iAccessTheLoginScreen)
	s.Step(`^jenkins should be unlocked$`, jenkinsShouldBeUnlocked)
}

