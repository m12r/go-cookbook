package fs

import (
	"embed"
	"io/fs"
	"os"
	"strings"
)

const (
	staticDir = "static"
	useDirEnv = "USE_DIR"
)

//go:embed static
var embedFS embed.FS

func FileSystem() fs.FS {
	fileSystem, _ := fs.Sub(embedFS, staticDir)
	if v := os.Getenv(useDirEnv); strings.EqualFold("true", v) || v == "1" {
		fileSystem = os.DirFS(staticDir)
	}
	return fileSystem
}
