package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
			err = generateAPIProject(cmd.Context(), projectName)
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
func generateAPIProject(ctx context.Context, projectName string) error {
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
	err = downloadAndUnzip(ctx, projectRoot)
	if err != nil {
		err := errors.Wrap(err, "Error downloading and unzipping template")
		return err
	}

	fmt.Printf("Starter Golang API project template created in '%s'\n", projectRoot)

	return nil
}

func downloadAndUnzip(ctx context.Context, targetDir string) error {
	templateURL := "https://github.com/victormazeli/api-gen-cli/raw/main/templates/template.zip"
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, templateURL, http.NoBody)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
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

	_, err = io.Copy(out, resp.Body)
	out.Close() // Close the file after copying is complete

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

	// Find the template folder within the ZIP file
	var templateFolder *zip.File
	for _, f := range r.File {
		if strings.HasPrefix(f.Name, "template/") {
			templateFolder = f
			break
		}
	}

	if templateFolder == nil {
		return fmt.Errorf("template folder not found in the ZIP file")
	}

	for _, f := range r.File {
		if !strings.HasPrefix(f.Name, "template/") {
			continue
		}

		// Construct the new file path
		extractedFilePath := filepath.Join(dest, strings.TrimPrefix(f.Name, "template/"))

		if f.FileInfo().IsDir() {
			// Create directory if it doesn't exist
			err = os.MkdirAll(extractedFilePath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		// Extract file from the template folder
		rc, err := f.Open()
		if err != nil {
			return err
		}

		outFile, err := os.Create(extractedFilePath)
		if err != nil {
			rc.Close()
			return err
		}

		limitReader := io.LimitedReader{R: rc, N: 1 << 20}
		_, err = io.Copy(outFile, &limitReader)
		if err != nil {
			rc.Close()
			outFile.Close()
			return err
		}

		rc.Close()
		outFile.Close()
	}

	// Remove the _MACOSX folder if it exists
	macOSXPath := filepath.Join(dest, "__MACOSX")
	err = os.RemoveAll(macOSXPath)
	if err != nil {
		return err
	}

	return nil
}

// func copyFile(src, dst string) error {
//	sourceFile, err := os.Open(src)
//	if err != nil {
//		return err
//	}
//	defer sourceFile.Close()
//
//	destFile, err := os.Create(dst)
//	if err != nil {
//		return err
//	}
//	defer destFile.Close()
//
//	_, err = io.Copy(destFile, sourceFile)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
