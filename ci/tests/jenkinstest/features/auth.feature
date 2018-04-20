Feature: Jenkins authentication
  In order to use Jenkins
  As a user
  I need to be able to login

  @unit,@jenkins
  Scenario: Can access login screen
    Given there is a Jenkins install
    When I access the login screen
    Then Jenkins should be unlocked
  
  @unit,@jenkins
  Scenario: Invalid user can't log in
    Given I access the login screen
    When I login using invalid credentials
    Then I am not logged in

  @unit,@jenkins
  Scenario: Can log in
    Given I access the login screen
    When I login using valid credentials
    Then I am logged in
