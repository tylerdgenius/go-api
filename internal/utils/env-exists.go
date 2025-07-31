package utils

import "os"

func EnvExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
