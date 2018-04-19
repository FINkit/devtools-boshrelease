Feature: Nexus integration
  In order to use artefacts 
  As a developer
  I need to be able to store artefacts
    
  Scenario: Slave job to nexus
    Given I login using valid credentials
    And I have added a maven job
    When I execute the maven job
    Then the job artefact is stored in nexus
