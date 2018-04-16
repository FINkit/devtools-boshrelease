package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	PLUGIN_MANAGER_URL string = "/pluginManager/api/json?depth=1&tree=plugins[shortName,version]"
)

type Plugins struct {
	Plugins []Plugin
}

type Plugin struct {
	Name    string `json:"shortName"`
	Version string `json:"version"`
}

func iAccessPluginManagement() error {
	pluginsResp, err := httpClient.Get(getUrl(PLUGIN_MANAGER_URL))

	if err != nil {
		return err
	}

	body, _ = getBodyString(pluginsResp)
	return nil
}

func createPluginArray(plugins Plugins) map[string]string {
	a := make(map[string]string)

	for _, plugin := range plugins.Plugins {
		a[plugin.Name] = plugin.Version
	}
	return a
}

func getAllInstalledPlugins() (Plugins, error) {
	var plugins Plugins
	json.Unmarshal([]byte(body), &plugins)

	if len(plugins.Plugins) == 0 {
		return plugins, fmt.Errorf("%s", "No plugins installed")
	}

	return plugins, nil
}

func getAllExpectedPlugins() (Plugins, error) {
	var plugins Plugins
	filePath, _ := os.Getwd()

	p, err := ioutil.ReadFile(filePath + "/../../../src/jenkins/plugins.txt")

	if err != nil {
		return Plugins{}, fmt.Errorf("%s", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(fmt.Sprintf("%s", p)))
	for scanner.Scan() {
		pluginLine := scanner.Text()
		splitPlugin := strings.Split(pluginLine, ":")
		plugin := Plugin{Name: splitPlugin[0], Version: splitPlugin[1]}
		plugins.Plugins = append(plugins.Plugins, plugin)
	}

	if len(plugins.Plugins) == 0 {
		return plugins, fmt.Errorf("%s", "No expected plugins")
	}

	return plugins, nil
}

func allThePluginsAreInstalled() error {
	installedPlugins, err := getAllInstalledPlugins()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	installedPluginsArray := createPluginArray(installedPlugins)
	expectedPlugins, err := getAllExpectedPlugins()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	for _, v := range expectedPlugins.Plugins {
		if installedPluginsArray[v.Name] == "" {
			return fmt.Errorf("Expected plugin %s is not installed", v.Name)
		}
		if installedPluginsArray[v.Name] != v.Version {
			return fmt.Errorf("Unexpected plugin version %s for %s (should be version %s)", installedPluginsArray[v.Name], v.Name, v.Version)
		}
	}
	return nil
}
