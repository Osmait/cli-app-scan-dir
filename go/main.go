package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	var ignoreDirs []string

	rootCmd := &cobra.Command{
		Use:   "explorer",
		Short: "Traverses the file system and allows ignoring specified directories",
		Run: func(cmd *cobra.Command, args []string) {
			scanDir(args, ignoreDirs)
		},
	}

	rootCmd.Flags().StringSliceVarP(&ignoreDirs, "ignore", "i", nil, "List of directories to ignore")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func scanDir(args []string, ignoreDirs []string) {
	if len(args) < 1 {
		fmt.Println("Usage: explorer <path>")
		os.Exit(1)
	}

	root := args[0]

	ignoredDirsMap := make(map[string]bool)
	for _, dir := range ignoreDirs {
		ignoredDirsMap[dir] = true
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && ignoredDirsMap[info.Name()] {
			return filepath.SkipDir // Skip the directory and its contents
		}

		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
	}
}
