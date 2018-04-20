Feature: Jenkins job management
  In order to use jobs
  As an administrator
  I need to be able to manage jobs

  @unit,@jenkins
  Scenario: Job creation
    Given I login using valid credentials
    And I have job config
    When I create the job
    Then the job is created

  @unit,@jenkins
  Scenario: Slave job execution
    Given I login using valid credentials
    And I have added a slave job
    When I execute the job
    Then the job is executed on a slave
