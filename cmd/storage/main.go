package main

import (
	"fmt"
	"github.com/seoperin/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	st.Name = "My storage"
	file, err := st.Upload("Filename", []byte("Hello world"))
	if err != nil {
		log.Fatal(err)
	}

	fileFromStorage, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	/*unknownID, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	fileNotFound, err := st.GetByID(unknownID)
	if err != nil {
		log.Fatal(err)
	}*/

	fmt.Println(fileFromStorage)
	//fmt.Printf("%d", fileNotFound.String())
}
