package readabilityserver

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/copyembed"
)

//go:embed readability_pb2_grpc.py readability_pb2.py readability_server.py
var content embed.FS

const (
	entryFile            = "readability_server.py"
	DefaultServerAddress = "localhost:9595"
)

func Initialize() {
	var crashErr error
	defer func() {
		if crashErr != nil {
			log.Fatal("readabilityserver.Initialize:", crashErr)
		}
	}()

	dir, err := os.MkdirTemp(os.TempDir(), "readability")
	if err != nil {
		crashErr = err
		return
	}
	defer os.RemoveAll(dir)
	if err := copyembed.CopyDirectory(content, ".", dir); err != nil {
		log.Println(err)
		return
	}

	pyFile := filepath.Join(dir, entryFile)
	cmd := exec.Command("python3", pyFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	listenAddr := fmt.Sprintf("LISTEN_ADDRESS=%s", DefaultServerAddress)
	cmd.Env = append(os.Environ(), listenAddr)
	if err := cmd.Run(); err != nil {
		crashErr = err
		return
	}
}
