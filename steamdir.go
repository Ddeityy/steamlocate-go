package steamlocate

type SteamDir struct {
	Path           string
	SteamApps      SteamApps
	LibraryFolders LibraryFolders
}

// Initializes the SteamDir struct and locates everything.
func (s *SteamDir) Locate() {
	s.locate()
}
