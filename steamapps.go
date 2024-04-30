package steamlocate

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type SteamApps struct {
	Path string
	Apps map[int]App
}

func (s *SteamApps) discover(libraryFolderPath string) error {
	s.Path = path.Join(libraryFolderPath)

	var appIds []int

	apps, err := os.ReadDir(s.Path)
	if err != nil {
		return fmt.Errorf("os.ReadDir: %w", err)
	}

	for _, file := range apps {
		if strings.Contains(file.Name(), "appmanifest_") {
			id, err := strconv.Atoi(strings.TrimSuffix(strings.Split(file.Name(), "_")[1], ".acf"))
			if err != nil {
				return fmt.Errorf("strconv.Atoi: %w", err)
			}
			appIds = append(appIds, id)
		}
	}

	s.Apps = make(map[int]App)

	for _, id := range appIds {
		app := App{}
		if err := app.New(s.Path, id); err != nil {
			// skip if there's a mismatch between libfolders and real dirs
			log.Println("app.New: %w", err)
		}
		s.Apps[id] = app
	}

	return nil
}
