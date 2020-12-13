// query.go contains all the query functions required to fulfill
// any query request.
//
// These are all pure functions so they take the data structures they
// need to perform the query as input parameters
package main

// getUserOrgAndTickets returns the organization that a user belongs to and
// all of the users tickets
// TODO error handling
func getUserOrgAndTickets(user *user, orgMap OrgMap, ticketMap TicketMap, submitterIDMap IndexMap) (*organization, []*ticket) {
	// get all tickets submitted by user
	ticketIDs := submitterIDMap[user.ID.String()]
	var tickets []*ticket
	for _, id := range ticketIDs {
		tickets = append(tickets, ticketMap[id])
	}

	return orgMap[user.OrganizationID.String()], tickets
}

// getOrgUsersAndTickets returns the users and tickets that belong
// to the org
// TODO error handling
func getOrgUsersAndTickets(org *organization, userMap UserMap, ticketMap TicketMap, ticketOrgIDMap IndexMap, userOrgIDMap IndexMap) ([]*user, []*ticket) {
	// get all the tickets of an org
	ticketIDs := ticketOrgIDMap[org.ID.String()]
	var tickets []*ticket
	for _, id := range ticketIDs {
		tickets = append(tickets, ticketMap[id])
	}
	// get all the users of an org
	userIDs := userOrgIDMap[org.ID.String()]
	var users []*user
	for _, id := range userIDs {
		users = append(users, userMap[id])
	}

	return users, tickets
}

func getTicketOrgAndUser(ticket *ticket, orgMap OrgMap, userMap UserMap) (*organization, *user) {
	return orgMap[ticket.OrganizationID.String()], userMap[ticket.SubmitterID.String()]
}

// getUsersByIndex looks up the given key in the indexMap and resolves
// all the matches to user structs using userMap
func getUsersByIndex(key string, indexMap IndexMap, userMap UserMap) []*user {
	ids := indexMap[key]
	var users []*user
	for _, id := range ids {
		users = append(users, userMap[id])
	}
	return users
}

// getOrgsByIndex looks up the given key in the indexMap and resolves
// all the matches to orgs structs using orgMap
func getOrgsByIndex(key string, indexMap IndexMap, orgMap OrgMap) []*organization {
	ids := indexMap[key]
	var orgs []*organization
	for _, id := range ids {
		orgs = append(orgs, orgMap[id])
	}
	return orgs
}

// getTicketsByIndex looks up the given key in the indexMap and resolves
// all the matches to ticket structs using ticketMap
func getTicketsByIndex(key string, indexMap IndexMap, ticketMap TicketMap) []*ticket {
	ids := indexMap[key]
	var tickets []*ticket
	for _, id := range ids {
		tickets = append(tickets, ticketMap[id])
	}
	return tickets
}
