package steamlocate

import (
	"fmt"
	"path"
)

type LibraryFolders struct {
	LibraryFolders []LibraryFolder
}

func discoverLibraryFolders(steamPath string) ([]LibraryFolder, error) {
	vdfPath := path.Join(steamPath, "steamapps", "libraryfolders.vdf")
	if err := PathExists(vdfPath); err != nil {
		return nil, fmt.Errorf("LibraryFolders.discover.PathExists: %w", err)
	}

	var parser = parseVDF(vdfPath)

	keys := parser.MapKeys("libraryfolders")

	libraryFolders := make([]LibraryFolder, len(keys))

	for i := range len(keys) {
		folderPath := parser.String(fmt.Sprintf("libraryfolders.%d.path", i))
		folderPath = path.Join(folderPath, "steamapps")
		libraryFolders[i].Path = folderPath
	}

	for i, folder := range libraryFolders {
		if err := folder.discover(); err != nil {
			return nil, fmt.Errorf("LibraryFolder.discover: %w", err)
		}
		libraryFolders[i] = folder
	}

	return libraryFolders, nil
}
