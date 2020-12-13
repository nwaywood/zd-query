// helper.go contains misc helper functions for file IO
package main

import (
	"encoding/json"
	"io/ioutil"
)

// loadData loads the source data files into memory
func loadData() ([]*organization, []*user, []*ticket, error) {
	// load organisations into slice
	orgs, err := loadOrgsFile("./data/organizations.json")
	if err != nil {
		return nil, nil, nil, err
	}
	// load tickets into slice
	tickets, err := loadTicketsFile("./data/tickets.json")
	if err != nil {
		return nil, nil, nil, err
	}
	// load users into slice
	users, err := loadUsersFile("./data/users.json")
	if err != nil {
		return nil, nil, nil, err
	}

	return orgs, users, tickets, nil
}

// loadUsersFile is a helper function for loading users.json
func loadUsersFile(path string) ([]*user, error) {
	userBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var users []*user
	err = json.Unmarshal(userBytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// loadTicketsFile is a helper function for loading tickets.json
func loadTicketsFile(path string) ([]*ticket, error) {
	ticketBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var tickets []*ticket
	err = json.Unmarshal(ticketBytes, &tickets)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// loadOrgsFile is a helper function for loading organizations.json
func loadOrgsFile(path string) ([]*organization, error) {
	orgBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var orgs []*organization
	err = json.Unmarshal(orgBytes, &orgs)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}
