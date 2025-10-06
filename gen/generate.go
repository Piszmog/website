package gen

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/Piszmog/website/components/core"
	"github.com/Piszmog/website/components/pages"
	"github.com/Piszmog/website/data"
	"github.com/Piszmog/website/version"
	"github.com/a-h/templ"
)

func GenerateTempl() error {
	cmd := exec.Command(
		"go",
		"tool",
		"templ",
		"generate",
		"-path",
		"./components",
	)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}

func GenerateStatic(dirPath string) error {
	// write root
	if err := writeIndexFile(dirPath, "", pages.Home()); err != nil {
		return err
	}

	// write about
	if err := writeIndexFile(path.Join(dirPath, "about"), "about", pages.About()); err != nil {
		return err
	}

	// write projects
	if err := writeIndexFile(path.Join(dirPath, "projects"), "projects", pages.Projects(data.ProjectsData)); err != nil {
		return err
	}

	// write experience
	if err := writeIndexFile(path.Join(dirPath, "experience"), "experience", pages.Experience(data.ExperienceData)); err != nil {
		return err
	}

	return nil
}

func writeIndexFile(dirPath string, page string, component templ.Component) error {
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	f, err := os.Create(path.Join(dirPath, indexFile))
	if err != nil {
		return err
	}
	defer f.Close()

	return core.HTML(page, component).Render(context.Background(), f)
}

const indexFile = "index.html"

func GenerateTailwindCSS() error {
	if err := os.RemoveAll("./dist/assets/css"); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	cmd := exec.Command(
		"go",
		"tool",
		"go-tw",
		"-i",
		"./styles/input.css",
		"-o",
		fmt.Sprintf("./dist/assets/css/output@%s.css", version.Value),
		"--minify",
	)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
