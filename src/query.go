package main

// getUserOrgAndTickets returns the organization that a user belongs to and
// all of the users tickets
func getUserOrgAndTickets(user *user, orgMap OrgMap, ticketMap TicketMap, submitterIDMap IndexMap) (*organization, []*ticket) {
	// get all tickets submitted by user
	ticketIDs := submitterIDMap[user.ID.String()]
	var tickets []*ticket
	for _, id := range ticketIDs {
		tickets = append(tickets, ticketMap[id])
	}

	return orgMap[user.OrganizationID.String()], tickets
}

func getUsersByIndex(key string, indexMap IndexMap, userMap UserMap) []*user {
	ids := indexMap[key]
	var users []*user
	for _, id := range ids {
		users = append(users, userMap[id])
	}
	return users
}
