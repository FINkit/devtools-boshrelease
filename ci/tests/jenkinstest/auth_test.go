package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const LOGIN_URL string = "/login"

var loginScreenBody string
var validLoginAttemptBody string
var invalidLoginAttemptBody string

func thereIsAJenkinsInstall() error {
	// TODO - Write a sensible test
	return nil
}

func iAccessTheLoginScreen() error {
	resp, err := http.Get(getJenkinsUrl(LOGIN_URL))
	if err != nil {
		return err
	}

	loginScreenBody, _ = getBodyString(resp)
	return nil
}

func jenkinsShouldBeUnlocked() error {
	if strings.Contains(loginScreenBody, "Unlock Jenkins") {
		return fmt.Errorf("expected %s not to contain 'Unlock Jenkins'", loginScreenBody)
	}
	return nil
}

func iLoginUsingValidCredentials() error {
	validLoginAttemptBody, _ = loginToJenkins(os.Getenv("JENKINS_PASSWORD"))
	return nil
}

func iLoginUsingInvalidCredentials() error {
	invalidLoginAttemptBody, _ = loginToJenkins("garbagepasswordthatwontwork123")
	return nil
}

func loginToJenkins(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("%s", "Empty password")
	}

	body, err := jenkinsLogin("administrator", password)

	if err != nil {
		return "ERROR", err
	}

	return body, nil
}

func iAmLoggedIn() error {
	if !strings.Contains(string(validLoginAttemptBody), `<a href="/logout"><b>log out</b></a>`) {
		return fmt.Errorf("expected %s to contain '/logout' link", validLoginAttemptBody)
	}

	return nil
}

func iAmNotLoggedIn() error {
	if strings.Contains(string(invalidLoginAttemptBody), `<a href="/logout"><b>log out</b></a>`) {
		return fmt.Errorf("expected %s to not contain '/logout' link", invalidLoginAttemptBody)
	}

	return nil
}
