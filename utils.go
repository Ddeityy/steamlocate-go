package steamlocate

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
)

func ParseVDF(vdfpath string) *koanf.Koanf {
	var k = koanf.New(".")

	if err := k.Load(file.Provider(vdfpath), Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	return k
}
