package steamlocate

import (
	"fmt"
	"path"
)

type App struct {
	Id   int
	Path string
	Name string
}

func (a *App) New(steamappsPath string, id int) error {
	appmanifestPath := path.Join(steamappsPath, fmt.Sprintf("appmanifest_%d.acf", id))
	if err := PathExists(appmanifestPath); err != nil {
		return fmt.Errorf("PathExists: %w", err)
	}

	k := parseVDF(appmanifestPath)

	appPath := path.Join(steamappsPath, "common", k.String("AppState.installdir"))

	if err := PathExists(appPath); err != nil {
		return fmt.Errorf("PathExists: %w", err)
	}

	a.Path = appPath
	a.Id = id
	a.Name = k.String("AppState.name")

	return nil
}
