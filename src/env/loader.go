package env

import (
	"os"
)

func Load() (*Vars) {
	return &Vars{
		Host:  os.Getenv("COMMANDER_HOST"),
		Port:  os.Getenv("COMMANDER_PORT"),
		Secret: os.Getenv("COMMANDER_SECRET"),
	}
}