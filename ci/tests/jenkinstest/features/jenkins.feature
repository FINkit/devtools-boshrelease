Feature: jenkins configuration
  In order have a valid jenkins installation
  As an administrator
  I need to be able to login to jenkins

  Scenario: Can access login screen
    Given there is a jenkins install
    When I access the login screen
    Then jenkins should be unlocked

  Scenario: Plugins are installed
    Given I have logged into Jenkins
    When I access plugin management
    Then all the plugins are installed
