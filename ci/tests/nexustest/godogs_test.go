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

func thereIsANexusInstall() error {
	url = os.Getenv("NEXUS_HOST") + "/login"
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

func nexusShouldBeUnlocked() error {
	if !strings.Contains(body, "Nexus Repository Manager") {
		return fmt.Errorf("expected %s to contain 'Nexus Repository Manager'", body)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a nexus install$`, thereIsANexusInstall)
	s.Step(`^I access the login screen$`, iAccessTheLoginScreen)
	s.Step(`^nexus should be unlocked$`, nexusShouldBeUnlocked)
}

