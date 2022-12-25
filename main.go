package main

import (
	"log"
	"path/filepath"

	"github.com/medenzel/task/cmd"
	"github.com/medenzel/task/models"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	//opening db
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Error using homedir!")
	}
	dbPath := filepath.Join(home, "tasks.db")
	db, err := models.Open(dbPath)
	if err != nil {
		log.Fatal("Error opening db!")
	}
	defer db.Close()

	//set task service with db methods
	ts := models.TaskService{
		DB: db,
	}

	//set command service with cobra commands
	cs := cmd.CommandService{
		TS: ts,
	}
	//executing cobra commands
	cs.Execute()
}
