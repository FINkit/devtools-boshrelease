package main

import (
	"fmt"
	"strings"
)
const (
	PLUGIN_MANAGER_URL string = "/pluginManager/api/xml?depth=1"
)

func iAccessPluginManagement() error {
	pluginsResp, err := httpClient.Get(getUrl(PLUGIN_MANAGER_URL))

	if err != nil {
		return err
	}

	body, _ = getBodyString(pluginsResp)
	return nil
}

func allThePluginsAreInstalled() error {
	if !strings.Contains(body, "<shortName>cucumber-reports</shortName>") {
		return fmt.Errorf("expected %s to contain 'cucumber-reports'", body)
	}
	return nil
}

