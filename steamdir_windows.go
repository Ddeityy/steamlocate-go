//go:build windows
// +build windows

package steamlocate

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func (s *SteamDir) locate() error {
	var k registry.Key
	var err error

	defer k.Close()

	switch runtime.GOARCH {
	case "amd64":
		k, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Valve\Steam`, registry.QUERY_VALUE)
	case "386":
		k, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Valve\Steam`, registry.QUERY_VALUE)
	}
	if err != nil {
		return fmt.Errorf("no steam registry key found: %w", err)
	}

	steamPath, _, err := k.GetStringValue("InstallPath")
	if err != nil {
		return fmt.Errorf("k.GetStringValue: %w", err)
	}
	s.Path = steamPath

	libraryFolders, err := discoverLibraryFolders(steamPath)
	if err != nil {
		return fmt.Errorf("discoverLibraryFolders: %w", err)
	}
	s.LibraryFolders = libraryFolders

	return nil
}
