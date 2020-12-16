package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildOrgCaches(t *testing.T) {
	orgMap, indexMap := buildOrgCaches(orgs)

	assert.Equal(t, expectedOrgMap, orgMap)
	assert.Equal(t, expectedOrgIndexMap, indexMap)
}

func TestBuildUserCaches(t *testing.T) {
	userMap, indexMap := buildUserCaches(users)

	assert.Equal(t, expectedUserMap, userMap)
	assert.Equal(t, expectedUserIndexMap, indexMap)
}

func TestBuildTicketCaches(t *testing.T) {
	ticketMap, indexMap := buildTicketCaches(tickets)

	assert.Equal(t, expectedTicketMap, ticketMap)
	assert.Equal(t, expectedTicketIndexMap, indexMap)
}
