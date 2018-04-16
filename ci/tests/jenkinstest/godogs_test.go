package main

import (
	"github.com/DATA-DOG/godog"
)

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a Jenkins install$`, thereIsAJenkinsInstall)
	s.Step(`^I access the login screen$`, iAccessTheLoginScreen)
	s.Step(`^Jenkins should be unlocked$`, jenkinsShouldBeUnlocked)
	s.Step(`^I login using invalid credentials$`, iLoginUsingInvalidCredentials)
	s.Step(`^I am not logged in$`, iAmNotLoggedIn)
	s.Step(`^I login using valid credentials$`, iLoginUsingValidCredentials)
	s.Step(`^I am logged in$`, iAmLoggedIn)
	s.Step(`^I access plugin management$`, iAccessPluginManagement)
	s.Step(`^all the plugins are installed$`, allThePluginsAreInstalled)

	s.BeforeScenario(func(interface{}) {
		createNewHttpClient()
	})
}