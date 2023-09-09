package steamlocate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/andygrunwald/vdf"
)

type SteamApp struct {
	AppId int
	Path  string
	Name  string
}

type SteamApps struct {
	Apps       map[int]SteamApp
	Discovered bool
}

type LibraryFolders struct {
	Paths      []string
	Discovered bool
}

type SteamDir struct {
	Path           string
	SteamApps      SteamApp
	LibraryFolders LibraryFolders
}

func locateHomeDir() string {
	return os.Getenv("HOME")
}

func locateSteamDir() string {
	homeDir := locateHomeDir()
	dir := path.Join(homeDir, ".steam", "steam")
	steamdDir, _ := filepath.Glob(dir)
	if len(steamdDir) > 0 {
		return steamdDir[0]
	}
	panic("Could not locate steam directory.")
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func getGameName(id int) string {
	dir := path.Join(locateSteamDir(), "steamapps")

	f, err := os.Open(path.Join(dir, fmt.Sprintf("appmanifest_%d.acf", id)))
	if err != nil {
		panic(err)
	}

	p := vdf.NewParser(f)

	m, err := p.Parse()
	if err != nil {
		panic(err)
	}

	gameName := m["AppState"].(map[string]interface{})["name"]

	return fmt.Sprintf("%v", gameName)
}

func locateGameDirectory(id int) {
	//dir := path.Join(locateSteamDir(), "steamapps", "common")
	//gameDir := path.Join(dir, getGameName(id))
	//fmt.Println(gameDir)

}
