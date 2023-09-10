package steamlocate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/andygrunwald/vdf"
	// "golang.org/x/sys"
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

func locateSteamDirLinux() string {
	switch runtime.GOOS {
	case "windows":
		return ""

	case "linux":

	}
	homeDir := locateHomeDir()
	dir := path.Join(homeDir, ".steam", "steam")
	steamdDir, _ := filepath.Glob(dir)
	if len(steamdDir) > 0 {
		return steamdDir[0]
	}
	panic("Could not locate steam directory.")
}

// import "golang.org/x/sys" is broken, lovely
// func locateSteamDirWindows(arch int) string {

// 	k, err := sys.windows.registry.OpenKey(sys.windows.registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, sys.windows.registry.QUERY_VALUE)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer k.Close()

// 	s, _, err := k.GetStringValue("SystemRoot")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(s)
// 	return ""

// }

func (s *SteamDir) locateDir() string {
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			return "" // locateSteamDirWindows(64)
		case "386":
			return "" // locateSteamDirWindows(32)
		}

	case "linux":
		return locateSteamDirLinux()
	}
	panic("This operating system is not supported.")
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func getGameName(id int) string {
	dir := path.Join(locateSteamDirLinux(), "steamapps")

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
