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
		log.Println("Standard installation not found")
	} else {
		s.Path = standardInstall
		s.LibraryFolders.discover(standardInstall)
		s.SteamApps.discover(standardInstall)
		return
	}

	// Check flatpak installation
	var flatpakInstall string = path.Join(homeDir, ".var", "app", "com.valvesoftware.Steam", steamPath)

	if _, err := os.Stat(flatpakInstall); os.IsNotExist(err) {
		log.Println("Flatpak installation not found")
	} else {
		s.Path = flatpakInstall
		s.LibraryFolders.discover(flatpakInstall)
		s.SteamApps.discover(flatpakInstall)
		return
	}

	log.Fatalf("No steam installations found.")

}
