package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Piszmog/website/gen"
	"github.com/Piszmog/website/log"
	"github.com/Piszmog/website/server"
	"github.com/Piszmog/website/server/router"
	"github.com/Piszmog/website/version"
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
		logger.Info("Clearing public directory")
		if err := os.RemoveAll("./public"); err != nil {
			logger.Error("Failed to clear public directory", err)
			return
		}

		version.Value = *verVal

		logger.Info("Generatng static files")
		if err := gen.GenerateStatic("./public"); err != nil {
			logger.Error("Failed to generate static files", err)
			return
		}

		logger.Info("Compiling tailwind")
		if err := gen.GenerateTailwindCSS(); err != nil {
			logger.Error("Failed to compile tailwind", err)
			return
		}

		logger.Info("Copying assets")
		if err := copyAssets("./dist/assets"); err != nil {
			logger.Error("Failed to copy assets", err)
		}
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
	if err := copyDir(dirPath, "./public/assets"); err != nil {
		return err
	}
	return os.Rename("./public/assets/_headers", "./public/_headers")
}

// copyDir recursively copies a directory from src to dst
func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
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
