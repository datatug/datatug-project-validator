package main

import (
	"context"
	"fmt"
	"os"

	"github.com/datatug/datatug-core/pkg/datatug"
	"github.com/datatug/datatug-core/pkg/storage"
	"github.com/datatug/datatug-core/pkg/storage/filestore"
)

func main() {
	projPath := os.Args[1]
	_, err := fmt.Println("Project path:", projPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to print project path: %s\n", err)
		os.Exit(1)
	}
	store, _ := filestore.NewSingleProjectStore(projPath, "")
	ctx := context.Background()
	projectStore := store.GetProjectStore(storage.SingleProjectID)
	var p *datatug.Project
	p, err = projectStore.LoadProject(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load project: %s\n", err)
		os.Exit(1)
	}
	if err = p.Validate(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to validate project: %s\n", err)
		os.Exit(1)
	}
	_, _ = fmt.Fprintln(os.Stderr, "DataTug project is valid")
}
