package gpsbabel

import (
	"os"
	"path"
)

var (
	TestAssetPath = ""
)

func init() {
	goPath := os.Getenv("GOPATH")
	projectPath := path.Join(goPath, "src", "github.com", "dsoprea", "go-gpsbabel")
	TestAssetPath = path.Join(projectPath, "test", "asset")
}
