//go:build linux
// +build linux

package steamlocate

import (
	"log"
	"os"
	"path"
)

func locateHomeDir() string {
	return os.Getenv("HOME")
}

func (s *SteamDir) locate() {

	homeDir := locateHomeDir()
	steamPath := path.Join(".steam", "steam")

	// Check normal installation
	standardInstall := path.Join(homeDir, steamPath)

	if _, err := os.Stat(standardInstall); os.IsNotExist(err) {
		log.Fatalf("%s", err)
	} else {
		s.Path = standardInstall
		return
	}

	// Check flatpak installation
	var flatpakInstall string = path.Join(homeDir, ".var", "app", "com.valvesoftware.Steam", steamPath)

	if _, err := os.Stat(flatpakInstall); os.IsNotExist(err) {
		log.Fatalf("%s", err)
	} else {
		s.Path = flatpakInstall
		return
	}

	s.LibraryFolders.discover(steamPath)
	s.SteamApps.discover()

}
