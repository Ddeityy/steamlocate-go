//go:build windows
// +build windows

package steamlocate

import (
	"log"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func (s *SteamDir) locate() {
	var k registry.Key
	var err error

	switch runtime.GOARCH {
	case "386":
		k, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Valve\Steam`, registry.QUERY_VALUE)
	case "amd64":
		k, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Valve\Steam`, registry.QUERY_VALUE)
	default:
		log.Fatalf("Failed to locate steam.")
	}

	log.Println(err)

	if err != nil {
		log.Fatalf("Failed to locate steam.")
	}
	defer k.Close()

	steamPath, _, err := k.GetStringValue("InstallPath")
	if err != nil {
		log.Fatal(err)
	}
	s.Path = steamPath
	s.LibraryFolders.discover(steamPath)
	s.SteamApps.discover(steamPath)
}
