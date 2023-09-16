package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func generateCmd() *cobra.Command {
	genCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a starter Golang API project template",
		Run: func(cmd *cobra.Command, args []string) {
			projectName, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Printf("Error getting project name: %v", err)
				return
			}
			generateAPIProject(projectName)
		},
	}

	genCmd.Flags().StringP("name", "n", "", "Name of the API project")
	return genCmd
}

// Function to generate project files for a starter api project.
func generateAPIProject(projectName string) {
	if projectName == "" {
		fmt.Println("Please provide a project name using the -n flag.")
		return
	}

	projectRoot := filepath.Join(".", projectName)

	// Create the project directory
	err := os.Mkdir(projectRoot, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating project directory: %v\n", err)
		return
	}

	// Create necessary directories and files
	directoryStructure := []string{
		"cmd",
		"internal/api/handlers",
		"internal/api/middlewares",
		"internal/config",
		"internal/db",
		"internal/models",
		"pkg/utils",
		"tests",
	}

	for _, dir := range directoryStructure {
		dirPath := filepath.Join(projectRoot, dir)
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// Create main.go file
	mainFileContent := `package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Add your API routes and handlers here

	r.Run(":8080")
}
`
	mainFilePath := filepath.Join(projectRoot, "cmd", "main.go")
	err = os.WriteFile(mainFilePath, []byte(mainFileContent), 0o600)
	if err != nil {
		fmt.Printf("Error creating main.go file: %v\n", err)
		return
	}

	fmt.Printf("Starter Golang API project template created in '%s'\n", projectRoot)
}
