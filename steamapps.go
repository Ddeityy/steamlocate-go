package steamlocate

import (
	"fmt"
	"path"
	"strconv"
)

type SteamApps struct {
	Path string
	Apps map[int]App
}

func (s *SteamApps) discover(steamPath string) {
	steamApps := path.Join(steamPath, "steamapps")
	libf := path.Join(steamApps, "libraryfolders.vdf")

	appIds := make([]string, 0)

	k := parseVDF(libf)

	for i := range k.MapKeys("libraryfolders") {
		appIds = append(appIds, k.MapKeys(fmt.Sprintf("libraryfolders.%d.apps", i))...)
	}

	s.Path = steamApps

	s.Apps = make(map[int]App)

	for _, value := range appIds {
		id, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		app := newApp(steamApps, id)
		s.Apps[id] = app
	}
}
