package steamlocate

type SteamDir struct {
	Path           string
	LibraryFolders []LibraryFolder
}

func (s *SteamDir) Locate() error {
	return s.locate()
}
