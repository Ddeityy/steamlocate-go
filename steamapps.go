package steamlocate

import (
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type SteamApps struct {
	Paths []string
	Apps  map[int]App // Separate apps in steamapps folder
}

func (s *SteamApps) discover(steamPath string, lfPaths []string) {
	var appIds []int
	for _, lfpath := range lfPaths {
		s.Paths = append(s.Paths, lfpath)
		apps, err := os.ReadDir(path.Join(lfpath))
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range apps {
			if strings.Contains(file.Name(), "appmanifest_") {
				id, err := strconv.Atoi(strings.TrimSuffix(strings.Split(file.Name(), "_")[1], ".acf"))
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
