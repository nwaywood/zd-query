package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// load source data from json files into slices
	orgs, users, tickets, err := loadData()
	if err != nil {
		log.Fatalf("Error loading data: %s", err)
	}
	// fmt.Printf("%v \n%v \n%v", orgs[0], users[0], tickets[0])
	// fmt.Printf("\n%s\n", orgs[0].ID.String())

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

// TODO: refactor this
func loadData() ([]*organization, []*user, []*ticket, error) {
	// load organisations into slice
	orgBytes, err := ioutil.ReadFile("../data/organizations.json")
	if err != nil {
		return nil, nil, nil, err
	}
	var orgs []*organization
	err = json.Unmarshal(orgBytes, &orgs)
	if err != nil {
		return nil, nil, nil, err
	}
	// load tickets into slice
	ticketBytes, err := ioutil.ReadFile("../data/tickets.json")
	if err != nil {
		return nil, nil, nil, err
	}
	var tickets []*ticket
	err = json.Unmarshal(ticketBytes, &tickets)
	if err != nil {
		return nil, nil, nil, err
	}
	// load users into slice
	userBytes, err := ioutil.ReadFile("../data/users.json")
	if err != nil {
		return nil, nil, nil, err
	}
	var users []*user
	err = json.Unmarshal(userBytes, &users)
	if err != nil {
		return nil, nil, nil, err
	}

	return orgs, users, tickets, nil
}
