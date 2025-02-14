package main

import (
	"app/cleaner"
	"app/dbadmin"
	"app/digestwriter"
	"app/manager"
	"app/pyxis"
	"app/vmsync"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "dbadmin":
			dbadmin.Start()
		case "manager":
			manager.Start()
		case "pyxis":
			pyxis.Start()
		case "vmsync":
			vmsync.Start()
		case "digestwriter":
			digestwriter.Start()
		case "cleaner":
			cleaner.Start()
		default:
			log.Fatalf("Unknown service name: %s\n", os.Args[1])
		}
	} else {
		log.Fatal("Please specify service name as the first argument.\n")
	}
}
