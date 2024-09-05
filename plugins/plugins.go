package plugins

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/clysec/greq"
)

const PluginDirectory = "plugins.dagosy.com"

type PluginVersion struct {
	Source     string `json:"src"`
	ApiVersion string `json:"apiVersion"`
	Namespace  string `json:"namespace"`
}

type PluginMetadata struct {
	Name        string                   `json:"name"`
	Slug        string                   `json:"slug"`
	Description string                   `json:"description"`
	Versions    map[string]PluginVersion `json:"versions"`
}

func GetPluginMetadata(plugin, apiVersion string) (PluginMetadata, error) {
	resp, err := greq.GetRequest(fmt.Sprintf("https://%s/%s/%s/metadata.json", PluginDirectory, apiVersion, plugin)).
		Execute()
	if err != nil {
		return PluginMetadata{}, err
	}

	var metadata PluginMetadata
	err = resp.BodyUnmarshalJson(&metadata)
	if err != nil {
		return PluginMetadata{}, err
	}

	return metadata, nil
}

func DownloadPlugin(pluginItem, pluginName, pluginSource string) error {
	pluginFile, err := os.OpenFile(pluginName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer pluginFile.Close()

	resp, err := http.Get(pluginSource)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(pluginFile, resp.Body)

	return nil
}
