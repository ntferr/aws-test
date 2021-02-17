package settings

import (
	"github.com/Netflix/go-env"
	"log"
)
var environment Environment

type Environment struct {
	Port	string	`env:"PORT"`
	Host	string	`env:"HOST"`
}


func init() {
	if _, err := env.UnmarshalFromEnviron(&environment); err != nil {
		log.Fatal(err)
	}
}

func GetEnv() Environment{

	return environment
}