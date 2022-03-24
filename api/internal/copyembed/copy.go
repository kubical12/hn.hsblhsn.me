package copyembed

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyDirectory(em embed.FS, srcDir, dest string) error {
	entries, err := em.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		file, err := em.Open(sourcePath)
		if err != nil {
			return err
		}

		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(em, destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(em, sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(em, sourcePath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func Copy(em embed.FS, srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := em.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	_, err = io.CopyBuffer(out, in, make([]byte, 4096))
	if err != nil {
		return err
	}

	return nil
}

func Exists(em embed.FS, filePath string) bool {
	if _, err := em.Open(filePath); err != nil {
		return false
	}
	return true
}

func CreateIfNotExists(em embed.FS, dir string, perm os.FileMode) error {
	if Exists(em, dir) {
		return nil
	}
	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}
	return nil
}
