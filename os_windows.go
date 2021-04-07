package main

import (
	"os"
	"os/user"
	"path/filepath"
)

func findDataDir() (dir string) {
	user, err := user.Current()
	if err != nil {
		return "."
	}
	dir = filepath.Join(user.HomeDir, "AppData", "Roaming")
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		".."
	}
	dir = filepath.Join(dir, "edpc-player")
	os.MkdirAll(dir, 0777)
	return dir
}
