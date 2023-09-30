package steamlocate

import (
	"fmt"
	"log"
	"path"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
)

type LibraryFolders struct {
	Paths []string
}

func (lf *LibraryFolders) Discover(p string) {

	vdfPath := path.Join(p, "steamapps", "libraryfolders.vdf")

	var k = koanf.New(".")

	if err := k.Load(file.Provider(vdfPath), Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	for i := 0; i < len(k.MapKeys("libraryfolders")); i++ {
		lf.Paths = append(lf.Paths, k.String(fmt.Sprintf("libraryfolders.%d.path", i)))
	}
}
