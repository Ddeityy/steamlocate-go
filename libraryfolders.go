package steamlocate

import (
	"fmt"
	"path"
)

type LibraryFolders struct {
	Paths []string
}

func (lf *LibraryFolders) discover(p string) {

	vdfPath := path.Join(p, "steamapps", "libraryfolders.vdf")

	var k = ParseVDF(vdfPath)

	for i := 0; i < len(k.MapKeys("libraryfolders")); i++ {
		lbf := k.String(fmt.Sprintf("libraryfolders.%d.path", i))
		lbf = path.Join(lbf, "steamapps")
		lf.Paths = append(lf.Paths, lbf)
	}
}
