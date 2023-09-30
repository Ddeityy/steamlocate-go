package main

import (
	"fmt"
	steamlocate "steamlocate/pkg"
)

func main() {
	var s steamlocate.SteamDir

	s.Locate()
	s.SteamApps.Discover()
	fmt.Println(s.SteamApps)
}
