package steamlocate

import (
	"fmt"
	"path"
)

type LibraryFolders struct {
	Paths []string
}

func (lf *LibraryFolders) discover(steamPath string) {

	vdfPath := path.Join(steamPath, "steamapps", "libraryfolders.vdf")

	var k = parseVDF(vdfPath)

	for i := 0; i < len(k.MapKeys("libraryfolders")); i++ {
		lbf := k.String(fmt.Sprintf("libraryfolders.%d.path", i))
		lbf = path.Join(lbf, "steamapps")
		lf.Paths = append(lf.Paths, lbf)
	}
}
