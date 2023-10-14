package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

// Generate Command to generate new API project.
func generateCmd() *cobra.Command {
	genCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a starter Golang API project template",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName, err := cmd.Flags().GetString("name")
			if err != nil {
				return errors.New("Project name must be provided")
			}
			err = generateAPIProject(projectName)
			if err != nil {
				return err
			}
			return nil
		},
	}

	genCmd.Flags().StringP("name", "n", "", "Name of the API project")
	return genCmd
}

// Function to generate project files for a starter api project.
func generateAPIProject(projectName string) error {
	if projectName == "" {
		err := errors.New("please provide a project name using the -n flag")
		return err
	}

	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		err := errors.Wrap(err, "Error getting current working directory")
		return err
	}

	projectRoot := filepath.Join(cwd, projectName)

	// Create the project directory
	err = os.Mkdir(projectRoot, os.ModePerm)
	if err != nil {
		err := errors.Wrap(err, "Error creating project directory")
		return err
	}

	// Download and unzip the template from the GitHub URL
	templateURL := "https://github.com/victormazeli/api-gen-cli/raw/main/templates/template.zip"
	err = downloadAndUnzip(templateURL, projectRoot)
	if err != nil {
		err := errors.Wrap(err, "Error downloading and unzipping template")
		return err
	}

	fmt.Printf("Starter Golang API project template created in '%s'\n", projectRoot)

	return nil
}

func downloadAndUnzip(url, targetDir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Create a temporary file to save the downloaded ZIP archive
	tmpZipFile := filepath.Join(targetDir, "template.zip")

	out, err := os.Create(tmpZipFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Unzip the downloaded ZIP archive
	err = unzip(tmpZipFile, targetDir)
	if err != nil {
		return err
	}

	// Remove the temporary ZIP file
	err = os.Remove(tmpZipFile)
	if err != nil {
		return err
	}

	return nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}

			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
