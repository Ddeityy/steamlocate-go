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

func newApp(steamappsPath string, id int) App {
	vdfPath := path.Join(steamappsPath, fmt.Sprintf("appmanifest_%d.acf", id))
	k := parseVDF(vdfPath)

	path := path.Join(steamappsPath, "common", k.String("AppState.installdir"))

	return App{
		Id:   id,
		Path: path,
		Name: k.String("AppState.name"),
	}
}
