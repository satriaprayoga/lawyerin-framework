package filesystem

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFTP struct {
	Host     string
	Port     string
	User     string
	Password string
}

func NewSFTP(h, p, u, pass string) FS {
	return &SFTP{Host: h, Port: p, User: u, Password: pass}
}

func (s *SFTP) getCredentials() (*sftp.Client, error) {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	config := ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	cwd, err := client.Getwd()
	if err != nil {
		return nil, err
	}
	log.Println("current directory: ", cwd)

	return client, nil
}

func (s *SFTP) Put(filename, folder string) error {
	var filepath string
	client, err := s.getCredentials()
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	defer client.Close()

	f, err := os.Open(filename)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	defer f.Close()

	if folder != "" {
		client.MkdirAll(folder)
		filepath = fmt.Sprintf("%s/%s", folder, path.Base(filename))
	} else {
		filepath = filename
	}

	f2, err := client.Create(filepath)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	defer f2.Close()

	if _, err := io.Copy(f2, f); err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}

func (s *SFTP) List(prefix string) ([]Listing, error) {
	var listing []Listing

	client, err := s.getCredentials()
	if err != nil {
		return listing, err
	}
	defer client.Close()

	files, err := client.ReadDir(prefix)
	if err != nil {
		return listing, err
	}

	for _, x := range files {
		var item Listing

		if !strings.HasPrefix(x.Name(), ".") {
			b := float64(x.Size())
			kb := b / (1024)
			mb := kb / (1024)
			item.Key = x.Name()
			item.Size = mb
			item.LastModified = x.ModTime()
			item.IsDir = x.IsDir()
			listing = append(listing, item)

		}
	}

	return listing, err
}

func (s *SFTP) Delete(items []string) bool {
	client, err := s.getCredentials()
	if err != nil {
		return false
	}
	defer client.Close()

	for _, x := range items {
		deleteErr := client.Remove(x)
		if deleteErr != nil {
			return false
		}
	}

	return true
}

func (s *SFTP) Get(destination string, items ...string) error {
	client, err := s.getCredentials()
	if err != nil {
		return err
	}
	defer client.Close()

	for _, item := range items {
		dstFile, err := os.Create(fmt.Sprintf("%s/%s", destination, path.Base(item)))
		if err != nil {
			return err
		}
		defer dstFile.Close()

		srcFile, err := client.Open(item)
		if err != nil {
			return err
		}

		bytes, err := io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}

		log.Printf("%d bytes copied ", bytes)

		err = dstFile.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}
