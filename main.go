package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Piszmog/website/log"
	"github.com/Piszmog/website/server"
	"github.com/Piszmog/website/server/router"
)

func main() {
	generate := flag.Bool("gen", false, "Generate the static files to ./public")
	verVal := flag.String("ver", "dev", "The version of the application")
	flag.Parse()

	logger := log.New(
		log.GetLevel(),
		log.GetOutput(),
	)

	if *generate {
	} else {
		svr := server.New(
			logger,
			":8080",
			server.WithRouter(router.New(logger)),
		)

		svr.StartAndWait()
	}
}

func copyAssets(dirPath string) error {
	return copyDir(dirPath, "./public/assets")
}

// copyDir recursively copies a directory from src to dst
func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	directory, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range directory {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// copyFile copies a single file from src to dst
func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
