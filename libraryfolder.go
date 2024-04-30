package steamlocate

import "fmt"

type LibraryFolder struct {
	Path         string
	SteamApps    SteamApps
	SteamCompats SteamCompats // TODO
	Shortcuts    *[]Shortcut  // TODO
}

func (lf *LibraryFolder) discover() error {
	lf.SteamApps = SteamApps{}
	if err := lf.SteamApps.discover(lf.Path); err != nil {
		return fmt.Errorf("lf.SteamApps.discover: %w", err)
	}

	return nil
}
