// file_io.go contains misc helper functions for file IO
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

// loadCacheFiles loads the cache files from disk and unmarshalls them into data structures
func loadCacheFiles() (OrgMap, IndexTypeMap, UserMap, IndexTypeMap, TicketMap, IndexTypeMap, error) {
	orgMapBytes, err := ioutil.ReadFile("./.cache/orgMap.json")
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	var orgMap OrgMap
	err = json.Unmarshal(orgMapBytes, &orgMap)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	userMapBytes, err := ioutil.ReadFile("./.cache/userMap.json")
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	var userMap UserMap
	err = json.Unmarshal(userMapBytes, &userMap)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	ticketMapBytes, err := ioutil.ReadFile("./.cache/ticketMap.json")
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	var ticketMap TicketMap
	err = json.Unmarshal(ticketMapBytes, &ticketMap)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	orgIndexesBytes, err := ioutil.ReadFile("./.cache/orgIndexes.json")
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	var orgIndexes IndexTypeMap
	err = json.Unmarshal(orgIndexesBytes, &orgIndexes)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	userIndexesBytes, err := ioutil.ReadFile("./.cache/userIndexes.json")
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	var userIndexes IndexTypeMap
	err = json.Unmarshal(userIndexesBytes, &userIndexes)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	ticketIndexesBytes, err := ioutil.ReadFile("./.cache/ticketIndexes.json")
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	var ticketIndexes IndexTypeMap
	err = json.Unmarshal(ticketIndexesBytes, &ticketIndexes)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return orgMap, orgIndexes, userMap, userIndexes, ticketMap, ticketIndexes, nil
}

// writeCacheFiles write the data structures used for indexing to disk
func writeCacheFiles(orgMap OrgMap, orgIndexes IndexTypeMap, userMap UserMap, userIndexes IndexTypeMap, ticketMap TicketMap, ticketIndexes IndexTypeMap) error {
	orgMapBytes, err := json.Marshal(orgMap)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./.cache/orgMap.json", orgMapBytes, 0777)
	if err != nil {
		return err
	}

	userMapBytes, err := json.Marshal(userMap)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./.cache/userMap.json", userMapBytes, 0777)
	if err != nil {
		return err
	}

	ticketMapBytes, err := json.Marshal(ticketMap)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./.cache/ticketMap.json", ticketMapBytes, 0777)
	if err != nil {
		return err
	}

	orgIndexesBytes, err := json.Marshal(orgIndexes)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./.cache/orgIndexes.json", orgIndexesBytes, 0777)
	if err != nil {
		return err
	}

	userIndexesBytes, err := json.Marshal(userIndexes)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./.cache/userIndexes.json", userIndexesBytes, 0777)
	if err != nil {
		return err
	}

	ticketIndexesBytes, err := json.Marshal(ticketIndexes)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./.cache/ticketIndexes.json", ticketIndexesBytes, 0777)
	if err != nil {
		return err
	}

	return nil
}
