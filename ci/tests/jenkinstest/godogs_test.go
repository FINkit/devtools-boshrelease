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
	s.Step(`^I have job config$`, iHaveJobConfig)
	s.Step(`^I create the job$`, iCreateTheJob)
	s.Step(`^the job is created$`, theJobIsCreated)
	s.Step(`I have added a slave job$`, iCreateTheSlaveJob)
	s.Step(`^I execute the job$`, iExecuteTheJob)
	s.Step(`^the job is executed on a slave$`, theJobIsExecuted)
	s.Step(`^I have added a Maven job$`, iHaveAddedAMavenJob)
	s.Step(`^I execute the Maven job$`, iExecuteTheMavenJob)
	s.Step(`^the job artefact is stored in Nexus$`, theJobArtefactIsStoredInNexus)

	s.BeforeScenario(func(interface{}) {
		createNewHttpClient()
	})
}
