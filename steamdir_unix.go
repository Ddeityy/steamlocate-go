//go:build linux
// +build linux

package steamlocate

import (
	"fmt"
	"log"
	"os"
	"path"
)

func locateHomeDir() string {
	return os.Getenv("HOME")
}

func (s *SteamDir) locate() error {
	homeDir := locateHomeDir()
	steamPath := path.Join(".steam", "steam")

	standardInstall := path.Join(homeDir, steamPath)

	if err := PathExists(standardInstall); err != nil {
		log.Println("Standard installation not found")
	} else {
		s.Path = standardInstall
		libraryFolders, err := discoverLibraryFolders(standardInstall)
		if err != nil {
			return fmt.Errorf("discoverLibraryFolders: %w", err)
		}
		s.LibraryFolders = libraryFolders
		return nil
	}

	var flatpakInstall string = path.Join(homeDir, ".var", "app", "com.valvesoftware.Steam", steamPath)

	if err := PathExists(flatpakInstall); err != nil {
		log.Println("Flatpak installation not found")
	} else {
		s.Path = flatpakInstall
		libraryFolders, err := discoverLibraryFolders(standardInstall)
		if err != nil {
			return fmt.Errorf("discoverLibraryFolders: %w", err)
		}
		s.LibraryFolders = libraryFolders
		return nil
	}

	return fmt.Errorf("no steam installations found")
}
