Feature: Nexus integration
  In order to use artefacts 
  As a developer
  I need to be able to store artefacts
    
  @integration
  Scenario: Slave job to Nexus
    Given I login using valid credentials
    And I have added a Maven job
    When I execute the Maven job
    Then the job artefact is stored in Nexus
