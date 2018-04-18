Feature: Jenkins job creation
  In order to run a job 
  As an administrator
  I need to be able to create a job

  Scenario: Job creation
    Given I have job config
    And I login using valid credentials
    When I create the job
    Then the job is created
