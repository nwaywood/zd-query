package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var err error

	// declare data structures for the caches
	var orgMap OrgMap
	var userMap UserMap
	var ticketMap TicketMap
	var orgIndexes IndexTypeMap
	var userIndexes IndexTypeMap
	var ticketIndexes IndexTypeMap

	// load cache files, if they exist and are not corrupted
	orgMap, orgIndexes, userMap, userIndexes, ticketMap, ticketIndexes, err = loadCacheFiles()
	// if err, load cache from source files and write cache files to disk
	if err != nil {
		// load source data from json files into slices
		orgs, users, tickets, err := loadData()
		if err != nil {
			log.Fatalf("Error loading data: %s", err)
		}

		// build caches from slices for fast queries
		orgMap, orgIndexes = buildOrgCaches(orgs)
		userMap, userIndexes = buildUserCaches(users)
		ticketMap, ticketIndexes = buildTicketCaches(tickets)

		// write cache files for future usage
		err = writeCacheFiles(orgMap, orgIndexes, userMap, userIndexes, ticketMap, ticketIndexes)
		if err != nil {
			fmt.Printf("Error storing cache: %s", err)
		}
	}

	// build the cli app
	app := buildCLIApp(orgMap, orgIndexes, userMap, userIndexes, ticketMap, ticketIndexes)

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
