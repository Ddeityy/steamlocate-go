package steamlocate

type Apps struct {
	Apps       map[int]App
	Discovered bool
}

// func (s *Apps) Discover() {
// 	var dir SteamDir
// 	dir.Locate()

// 	steamapps := path.Join(dir.Path, "steamapps")

// 	f, err := os.Open(path.Join(steamapps, fmt.Sprintf("appmanifest_%d.acf")))
// 	if err != nil {
// 		log.Fatalf("%s", err)
// 	}

// 	p := vdf.NewParser(f)

// 	m, err := p.Parse()
// 	if err != nil {
// 		log.Fatalf("%s", err)
// 	}

// name := m["AppState"].(map[string]interface{})["name"].(string)
// installDir := m["AppState"].(map[string]interface{})["installdir"].(string)
// if err != nil {
// 	log.Fatalf("%s", err)
// }

//}
