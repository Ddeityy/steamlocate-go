package steamlocate

import (
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type SteamApps struct {
	SteamDir *SteamDir
	Paths    []string
	Apps     map[int]App // Separate apps in steamapps folder
}

func (s *SteamApps) Discover() {
	s.SteamDir.LibraryFolders.Discover()

	var appIds []int
	for _, lfpath := range s.SteamDir.LibraryFolders.Paths {
		s.Paths = append(s.Paths, lfpath, "steamapps")
		apps, err := os.ReadDir(path.Join(lfpath, "steamapps"))
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range apps {
			if strings.Contains(file.Name(), "appmanifest_") {
				id, err := strconv.Atoi(strings.TrimSuffix(strings.Split(file.Name(), "_")[1], ".dem"))
				if err != nil {
					log.Panic(err)
				}
				appIds = append(appIds, id)
			}
		}
	}

	s.Apps = make(map[int]App)

	for _, id := range appIds {
		var app App

		for _, spath := range s.Paths {
			app = newApp(spath, id)
			s.Apps[id] = app
		}
	}
}
