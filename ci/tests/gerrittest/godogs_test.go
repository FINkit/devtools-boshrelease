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

func thereIsAGerritInstall() error {
	url = os.Getenv("GERRIT_HOST")
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

func gerritShouldBeUnlocked() error {
	if !strings.Contains(body, "Gerrit Code Review") {
		return fmt.Errorf("expected %s to contain 'Gerrit Code Review'", body)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a gerrit install$`, thereIsAGerritInstall)
	s.Step(`^I access the login screen$`, iAccessTheLoginScreen)
	s.Step(`^gerrit should be unlocked$`, gerritShouldBeUnlocked)
}

