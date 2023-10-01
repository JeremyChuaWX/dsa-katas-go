package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	name := flag.String("name", "", "name of the kata")
	force := flag.Bool("force", false, "replace existing files")
	flag.Parse()

	if name == nil || *name == "" {
		log.Fatal("missing kata name")
	}

	templatePath := fmt.Sprintf("./templates/%s.template", *name)
	stubPath := fmt.Sprintf("./katas/%s.go", *name)

	_, err := os.Stat(stubPath)
	if !errors.Is(err, os.ErrNotExist) && !*force {
		log.Fatalf("%s already exists", stubPath)
	}

	templateFile, err := os.Open(templatePath)
	if err != nil {
		log.Fatal(err)
	}
	defer templateFile.Close()

	stubFile, err := os.Create(stubPath)
	if err != nil {
		log.Fatal(err)
	}
	defer stubFile.Close()

	bytesWritten, err := io.Copy(stubFile, templateFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"generated %s from %s (%d bytes written)\n",
		stubPath,
		templatePath,
		bytesWritten,
	)
}
