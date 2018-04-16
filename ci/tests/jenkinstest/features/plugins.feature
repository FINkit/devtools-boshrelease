Feature: Jenkins plugin installation and configuration
  In order to use Jenkins plugins
  As an administrator
  I need to be able to install and configure plugins

  Scenario: Plugins are installed
    Given I login using valid credentials
    When I access plugin management
    Then all the plugins are installed
