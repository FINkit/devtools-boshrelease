Feature: gerrit configuration
  In order have a valid gerrit installation
  As an administrator
  I need to be able to login to gerrit

  Scenario: Can access login screen
    Given there is a gerrit install
    When I access the login screen
    Then gerrit should be unlocked
