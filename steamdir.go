package steamlocate

type SteamDir struct {
	Path           string
	SteamApps      SteamApps // steamapps folder
	LibraryFolders LibraryFolders
}

// Initializes the SteamDir struct and locates everything.
func (s *SteamDir) Locate() {
	s.locate()
}
