//go:build linux
// +build linux

package cmd

func ConsoleInit() error {
	return nil
}

func defaultConfigFile() string {
	return "/etc/skynx/skynx-node.yml"
}
