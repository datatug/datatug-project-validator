package main

import (
	"context"
	"fmt"
	"os"

	"github.com/datatug/datatug-core/pkg/storage/filestore"
)

func main() {
	projPath := os.Args[1]
	_, err := fmt.Println("Project path:", projPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to print project path: %s\n", err)
		os.Exit(1)
	}
	loader, _ := filestore.NewSingleProjectLoader(projPath)
	ctx := context.Background()
	_, err = loader.LoadProject(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load project: %s\n", err)
		os.Exit(1)
	}
	if _, err = fmt.Fprintln(os.Stderr, "project is valid"); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to print project info: %s\n", err)
		os.Exit(1)
	}
}
