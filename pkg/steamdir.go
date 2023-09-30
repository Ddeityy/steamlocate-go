package steamlocate

import (
	"log"
	"os"
	"path"
)

type SteamDir struct {
	Path           string
	Apps           Apps
	LibraryFolders LibraryFolders
}

func (s *SteamDir) Locate() {

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
}

func locateHomeDir() string {
	return os.Getenv("HOME")
}
