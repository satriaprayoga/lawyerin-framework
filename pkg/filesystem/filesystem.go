package filesystem

import (
	"os"
	"time"
)

type FS interface {
	Put(filename, folder string) error
	Get(destination string, items ...string) error
	List(prefix string) ([]Listing, error)
	Delete(items []string) bool
}

type Listing struct {
	Etag         string
	LastModified time.Time
	Key          string
	IsDir        bool
	Size         float64
}

func New() FS {
	if os.Getenv("SFTP_HOST") != "" {
		fs := NewSFTP(
			os.Getenv("SFTP_HOST"),
			os.Getenv("SFTP_PORT"),
			os.Getenv("SFTP_USER"),
			os.Getenv("SFTP_PASSWORD"),
		)

		return fs
	}
	return nil
}
