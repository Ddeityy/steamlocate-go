# steamlocate

A (worse) Go version of [steamlocate-rs](https://github.com/WilliamVenner/steamlocate-rs) which locates any Steam installation and Steam application on the filesystem.

**This library supports Linux, Windows and Macos.**

## TODO

* Shortcuts (non-steam games)
* Steam compatibility tools (proton etc.)

Download:

```console
go get -u github.com/ddeityy/steamlocate-go
```

## Examples

### Locate the installed Steam directory

```go
import "github.com/ddeityy/steamlocate-go"

s := steamlocatego.SteamDir{}
if err := s.Locate(); err != nil {
    log.Fatalln(err)
}

fmt.Println(s.Path)
```

```go
SteamDir {
    "/home/$USER/.steam/steam"
}
```

### Locate all installed Steam apps or a specific one by its app ID

```go
import "github.com/ddeityy/steamlocate-go"

s := steamlocatego.SteamDir{}
if err := s.Locate(); err != nil {
    log.Fatalln(err)
}

fmt.Println(s.LibraryFolders[0].SteamApps.Apps[440])
```

```go
App {
    440
    /home/$USER/.local/share/Steam/steamapps/common/Team Fortress 2
    Team Fortress 2
}
```

#### Locate all Steam library folders

```go
import "github.com/ddeityy/steamlocate-go"

s := steamlocatego.SteamDir{}
if err := s.Locate(); err != nil {
    log.Fatalln(err)
}

fmt.Println(s.LibraryFolders)
```

```go
{
    "/home/$USER/.steam/steam/steamapps"
    "/second_drive/steam/steamapps"
    ...
}
```
