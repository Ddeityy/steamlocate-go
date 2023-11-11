//go:build darwin
// +build darwin

package steamlocate

import (
	"log"
	"os"
	"path"
)

func (s *SteamDir) locate() {
	homeDir := locateHomeDir()

	steamPath := path.Join(homeDir, "Library", "Application Support", "Steam")

	if _, err := os.Stat(steamPath); os.IsNotExist(err) {
		log.Fatalf("Steam not found.")
	}
	s.Path = steamPath
	s.LibraryFolders.discover(steamPath)
	s.SteamApps.discover(steamPath)
	return
}
