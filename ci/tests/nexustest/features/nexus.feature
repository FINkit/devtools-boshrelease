Feature: nexus configuration
  In order have a valid nexus installation
  As an administrator
  I need to be able to login to nexus

  Scenario: Can access login screen
    Given there is a nexus install
    When I access the login screen
    Then nexus should be unlocked
