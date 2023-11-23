package steamlocate

import (
	"fmt"
	"path"
)

type LibraryFolders struct {
	SteamDir *SteamDir
	Paths    []string
}

func (lf *LibraryFolders) Discover() {

	vdfPath := path.Join(lf.SteamDir.Path, "steamapps", "libraryfolders.vdf")

	var k = parseVDF(vdfPath)

	for i := 0; i < len(k.MapKeys("libraryfolders")); i++ {
		lbf := k.String(fmt.Sprintf("libraryfolders.%d.path", i))
		lbf = path.Join(lbf, "steamapps")
		lf.Paths = append(lf.Paths, lbf)
	}
}
