# steamlocate

A (worse) Go version of [steamlocate-rs](https://github.com/WilliamVenner/steamlocate-rs) which locates any Steam installation and Steam application on the filesystem.

**This library supports Linux, Windows and Macos.**

## TODO
* Shortcuts (non-steam games)
* Steam compatibility tools (proton etc.)


## Using steamlocate
Download:
```console
$ go get -u github.com/ddeityy/steamlocate-go
```

## Examples

#### Locate the installed Steam directory

```go
import "github.com/ddeityy/steamlocate-go"

var s steamlocate.SteamDir

s.Locate()

fmt.Println(s.Path)
```
```go
SteamDir {
    Path string: "/home/$USER/.steam/steam"
}
```

#### Locate all installed Steam apps or a specific one by it's app ID

```go
import "github.com/ddeityy/steamlocate-go"

var s steamlocate.SteamDir

s.Locate()

fmt.Println(s.SteamApps.Apps)

fmt.Println(s.SteamApps.Apps[440].Name)

```
```go
Apps {
    440: {
    ID int: 440 
    Path string: /home/deity/.steam/steam/steamapps/common/Team Fortress 2
    Name string: Team Fortress 2}
    ...
}
```

#### Locate all Steam library folders
```go
import "github.com/ddeityy/steamlocate-go"

var s steamlocate.SteamDir

s.Locate()

fmt.Println(s.LibraryFolders.Paths)
```
```go
{
    "/home/$USER/.steam/steam/steamapps"
    "/second_drive/steam/steamapps"
    ...
}
```
