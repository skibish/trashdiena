package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/skibish/trashdiena/pkg/config"

	"github.com/skibish/trashdiena/pkg/firebase"
	"github.com/skibish/trashdiena/pkg/storage"
	uuid "github.com/satori/go.uuid"
)

func main() {
	c := config.Parse()
	// Try to initialize all mandatory parts of the application
	fbase, err := firebase.New(c.FirebaseDB, c.FirebaseSecretPath)
	if err != nil {
		log.Fatal(err)
	}
	sg := storage.New(fbase)

	file, err := os.Open(c.PathToFileWithTrash)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		err = sg.Trash.Set(&storage.TrashData{
			ID:        uuid.NewV4().String(),
			Data:      record[0],
			Published: false,
		})
	}
}
