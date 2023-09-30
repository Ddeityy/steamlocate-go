package main

import (
	"fmt"
	steamlocate "steamlocate/pkg"
)

func main() {
	var s steamlocate.SteamDir

	s.Locate()

	s.LibraryFolders.Discover(s.Path)
	fmt.Println(s.LibraryFolders.Paths)

}
