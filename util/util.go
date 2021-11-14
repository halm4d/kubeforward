package util

import (
	"os"
	"path/filepath"
)

func GetEnvOrDefault(e string, d string) string {
	if env := os.Getenv(e); env != "" {
		return env
	} else {
		return d
	}
}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func ReadNamespaceArg() string {
	ns := os.Args[1]
	if ns == "" {
		panic("Namespace is required.")
	}
	return ns
}
