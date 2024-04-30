package main

import (
	"log"

	steamlocatego "github.com/ddeityy/steamlocate-go"
)

func main() {
	s := steamlocatego.SteamDir{}
	if err := s.Locate(); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", s.LibraryFolders[0].SteamApps.Apps[440])
}
