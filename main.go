package main

import (
	steamlocate "steamlocate/pkg"

	"github.com/kr/pretty"
)

func main() {
	var s steamlocate.SteamDir

	s.SteamApps.Discover()
	pretty.Println(s.SteamApps.Apps[440].Path)
}
