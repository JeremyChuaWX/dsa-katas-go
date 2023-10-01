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
	force := flag.Bool("force", false, "name of the kata")
	flag.Parse()

	if name == nil || *name == "" {
		log.Fatal("missing kata name")
	}

	templatePath := fmt.Sprintf("./templates/%s.tmpl", *name)
	newPath := fmt.Sprintf("./katas/%s.go", *name)

	_, err := os.Stat(newPath)
	if !errors.Is(err, os.ErrNotExist) && !*force {
		log.Fatalf("%s already exists", newPath)
	}

	templateFile, err := os.Open(templatePath)
	if err != nil {
		log.Fatal(err)
	}
	defer templateFile.Close()

	newFile, err := os.Create(newPath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, templateFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"generated ./katas/%s.go from ./templates/%s.tmpl (%d bytes written)\n",
		*name,
		*name,
		bytesWritten,
	)
}
