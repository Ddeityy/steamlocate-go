//go:build windows
// +build windows

package steamlocate

import (
	"fmt"
	"log"
)

func (s *SteamDir) Locate() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\Microsoft\Windows NT\CurrentVersion", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("SystemRoot")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Windows system root is %q\n", s)
}
