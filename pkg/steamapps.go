package steamlocate

import (
	"fmt"
	"path"
	"strconv"
)

type SteamApps struct {
	SteamApps map[int]App
}

func (s *SteamApps) Discover() {
	var steamDir SteamDir
	steamDir.Locate()
	steamApps := path.Join(steamDir.Path, "steamapps")
	libf := path.Join(steamApps, "libraryfolders.vdf")

	appIds := make([]string, 0)

	k := ParseVDF(libf)

	fmt.Println(len(k.MapKeys("libraryfolders.0.apps")))

	for i := range k.MapKeys("libraryfolders") {
		appIds = append(appIds, k.MapKeys(fmt.Sprintf("libraryfolders.%d.apps", i))...)
	}

	s.SteamApps = make(map[int]App)

	for _, value := range appIds {
		id, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		app := newApp(steamApps, id)
		s.SteamApps[id] = app
	}
}
