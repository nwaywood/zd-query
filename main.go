package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// load source data from json files into slices
	orgs, users, tickets, err := loadData()
	if err != nil {
		log.Fatalf("Error loading data: %s", err)
	}
	// fmt.Printf("%v %v %v", orgs[0], users[0], tickets[0])
	// fmt.Printf("\n%s\n", orgs[0].ID.String())

	// build caches from slices for fast queries
	orgMap, orgIndexes := buildOrgCaches(orgs)
	userMap, userIndexes := buildUserCaches(users)
	ticketMap, ticketIndexes := buildTicketCaches(tickets)

	// build the cli app
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: refactor this
func loadData() ([]*Organization, []*User, []*Ticket, error) {
	// load organisations into slice
	orgBytes, err := ioutil.ReadFile("./data/organizations.json")
	if err != nil {
		return nil, nil, nil, err
	}
	var orgs []*Organization
	err = json.Unmarshal(orgBytes, &orgs)
	if err != nil {
		return nil, nil, nil, err
	}
	// load tickets into slice
	ticketBytes, err := ioutil.ReadFile("./data/tickets.json")
	if err != nil {
		return nil, nil, nil, err
	}
	var tickets []*Ticket
	err = json.Unmarshal(ticketBytes, &tickets)
	if err != nil {
		return nil, nil, nil, err
	}
	// load users into slice
	userBytes, err := ioutil.ReadFile("./data/users.json")
	if err != nil {
		return nil, nil, nil, err
	}
	var users []*User
	err = json.Unmarshal(userBytes, &users)
	if err != nil {
		return nil, nil, nil, err
	}

	return orgs, users, tickets, nil
}
