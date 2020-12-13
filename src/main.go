package main

import (
	"log"
	"os"
)

func main() {
	// load source data from json files into slices
	orgs, users, tickets, err := loadData()
	if err != nil {
		log.Fatalf("Error loading data: %s", err)
	}

	// build caches from slices for fast queries
	orgMap, orgIndexes := buildOrgCaches(orgs)
	userMap, userIndexes := buildUserCaches(users)
	ticketMap, ticketIndexes := buildTicketCaches(tickets)

	// build the cli app
	app := buildCLIApp(orgMap, orgIndexes, userMap, userIndexes, ticketMap, ticketIndexes)

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
