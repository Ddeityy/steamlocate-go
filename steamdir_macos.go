//go:build darwin
// +build darwin

package steamlocate

import (
	"fmt"
	"os"
	"path"
)

func locateHomeDir() string {
	return os.Getenv("HOME")
}

func (s *SteamDir) locate() error {
	homeDir := locateHomeDir()

	steamPath := path.Join(homeDir, "Library", "Application Support", "Steam")
	if err := PathExists(steamPath); err != nil {
		return fmt.Errorf("no steam installation found: %w", err)
	}
	s.Path = steamPath

	libraryFolders, err := discoverLibraryFolders(steamPath)
	if err != nil {
		return fmt.Errorf("discoverLibraryFolders: %w", err)
	}
	s.LibraryFolders = libraryFolders

	return nil
}
