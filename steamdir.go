package steamlocate

import (
	"fmt"
)

type SteamDir struct {
	Path           string
	LibraryFolders []LibraryFolder
}

func (s *SteamDir) Locate() error {
	return s.locate()
}

func (s *SteamDir) FindApp(id int) (*App, error) {
	for _, lf := range s.LibraryFolders {
		for _, app := range lf.SteamApps.Apps {
			if app.Id == id {
				return &app, nil
			}
		}
	}
	return nil, fmt.Errorf("app with id %d not found", id)
}
