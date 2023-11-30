package main

import (
	"fmt"

	"github.com/ddeityy/steamlocate-go"
)

func main() {
	s := steamlocate.SteamDir{}
	s.Locate()
	fmt.Println(s.SteamApps.Apps[440].Path)
}
