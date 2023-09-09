package steamlocate

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/andygrunwald/vdf"
)

func locateHomeDir() string {
	return os.Getenv("HOME")
}

func locateSteamDir() string {
	homeDir := locateHomeDir()
	dir := path.Join(homeDir, ".steam")
	steamdDir, _ := filepath.Glob(dir)
	if len(steamdDir) > 0 {
		return steamdDir[0]
	}
	panic("Could not locate steam directory.")
}

func getGameName(id int) {
	dir := path.Join(locateSteamDir(), "steamapps")

	f, err := os.Open(path.Join(dir, "libraryfolders.vdf"))
	if err != nil {
		panic(err)
	}

	p := vdf.NewParser(f)

	m, err := p.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}

func locateGameDirectory(id int) {
	//dir := path.Join(locateSteamDir(), "steamapps", "common")
	//gameDir := path.Join(dir, getGameName(id))
	//fmt.Println(gameDir)

}
