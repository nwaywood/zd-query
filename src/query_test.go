package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserOrgAndTickets(t *testing.T) {
	// Happy case: User exists and has org
	org, ts, err := getUserOrgAndTickets(users[0], expectedOrgMap, expectedTicketMap, expectedTicketIndexMap["submitter_id"])

	expectedTickets := []*ticket{expectedTicketMap["t1"]}

	assert.NoError(t, err)
	assert.Equal(t, orgs[2], org)
	assert.Equal(t, expectedTickets, ts)

	// Unhappy case: User doesn't belong to org
	userInvalidOrg := &user{
		ID:             "105",
		Name:           "John Doe",
		OrganizationID: "5",
	}

	_, _, err = getUserOrgAndTickets(userInvalidOrg, expectedOrgMap, expectedTicketMap, expectedTicketIndexMap["submitter_id"])
	assert.Error(t, err)
}

func TestGetOrgUsersAndTickets(t *testing.T) {
	us, ts := getOrgUsersAndTickets(orgs[2], expectedUserMap, expectedTicketMap, expectedTicketIndexMap["organization_id"], expectedUserIndexMap["organization_id"])

	assert.Equal(t, tickets, ts)
	assert.Equal(t, users, us)
}

func TestGetTicketOrgAndUser(t *testing.T) {
	// Happy case: Ticket exists and has org and user
	o, u, err := getTicketOrgAndUser(tickets[0], expectedOrgMap, expectedUserMap)
	assert.NoError(t, err)
	assert.Equal(t, orgs[2], o)
	assert.Equal(t, users[0], u)

	// Unhappy case: Ticket doesn't belong to org
	_ticket := &ticket{
		ID:             "t1",
		OrganizationID: "7",
		SubmitterID:    "101",
	}
	_, _, err = getTicketOrgAndUser(_ticket, expectedOrgMap, expectedUserMap)
	assert.Error(t, err)

	// Unhappy case: Ticket doesn't belong to a user
	_ticket = &ticket{
		ID:             "t1",
		OrganizationID: "3",
		SubmitterID:    "111",
	}
	_, _, err = getTicketOrgAndUser(_ticket, expectedOrgMap, expectedUserMap)
	assert.Error(t, err)
}

func TestGetUsersByIndex(t *testing.T) {
	xs := getUsersByIndex("space", expectedUserIndexMap["tags"], expectedUserMap)

	assert.Equal(t, users, xs)
}

func TestGetTicketsByIndex(t *testing.T) {
	xs := getTicketsByIndex("high", expectedTicketIndexMap["priority"], expectedTicketMap)

	expected := []*ticket{expectedTicketMap["t1"], expectedTicketMap["t3"]}

	assert.Equal(t, expected, xs)
}

func TestGetOrgsByIndex(t *testing.T) {
	xs := getOrgsByIndex("empty", expectedOrgIndexMap["tags"], expectedOrgMap)

	expected := []*organization{}

	assert.Equal(t, expected, xs)
}
