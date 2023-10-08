package main

import (
	"fmt"
	"io"
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
		err := errors.Wrap(err, "Error creating project directory")
		return err
	}

	projectRoot := filepath.Join(cwd, projectName)

	// Create the project directory
	err = os.Mkdir(projectRoot, os.ModePerm)
	if err != nil {
		err := errors.Wrap(err, "Error creating project directory")
		return err
	}

	// Get the absolute path of a specific directory.
	sourceDir := filepath.Join(cwd, "templates", "REST")

	// Walk through the source directory and copy all files and folders.
	err = filepath.Walk(sourceDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate the destination path by replacing the source directory with the project root directory.
		relPath, err := filepath.Rel(sourceDir, srcPath)
		if err != nil {
			return err
		}
		destPath := filepath.Join(projectRoot, relPath)

		if info.IsDir() {
			// Create directories in the destination path.
			err := os.MkdirAll(destPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			// Copy files from source to destination.
			err := copyFile(srcPath, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		err = errors.Wrap(err, "Error copying files and folders")
		return err
	}

	fmt.Printf("Starter Golang API project template created in '%s'\n", projectRoot)

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
