package steamlocate

type SteamDir struct {
	Path           string
	SteamApps      SteamApps
	LibraryFolders LibraryFolders
}

func (s *SteamDir) Locate() {
	s.locate()
}
