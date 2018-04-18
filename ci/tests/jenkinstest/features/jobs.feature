Feature: Jenkins job management
  In order to use jobs 
  As an administrator
  I need to be able to manage jobs

  Scenario: Job creation
    Given I login using valid credentials
    And I have job config
    When I create the job
    Then the job is created

  Scenario: Slave job execution
    Given I login using valid credentials
    And I have added a slave job
    When I execute the job
    Then the job is executed on a slave
