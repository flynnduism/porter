package plugins

import "fmt"

// TestPluginProvider helps us test Porter.Plugins in our unit tests without actually hitting any real plugins on the file system.
type TestPluginProvider struct{}

func (p *TestPluginProvider) List() ([]string, error) {
	mixins := []string{"plugin1", "plugin2", "unknown"}
	return mixins, nil
}

func (p *TestPluginProvider) GetMetadata(pluginName string) (*Metadata, error) {
	var impl []Implementation
	if pluginName != "unknown" {
		impl = []Implementation{
			{Type: "instance-storage", Name: "blob"},
			{Type: "instance-storage", Name: "mongo"},
		}
	}
	return &Metadata{
		Name:            pluginName,
		ClientPath:      fmt.Sprintf("/home/porter/.porter/plugins/%s", pluginName),
		Implementations: impl,
		VersionInfo:     VersionInfo{Version: "v1.0", Commit: "abc123", Author: "Porter Authors"},
	}, nil
}
