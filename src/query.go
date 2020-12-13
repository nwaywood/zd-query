package main

func getUserOrgAndTickets(user *user, orgMap OrgMap, ticketMap TicketMap, submitterIDMap IndexMap) (*organization, []*ticket) {
	// get all tickets submitted by user
	ticketIDs := submitterIDMap[user.ID.String()]
	var tickets []*ticket
	for _, id := range ticketIDs {
		tickets = append(tickets, ticketMap[id])
	}

	return orgMap[user.OrganizationID.String()], tickets
}
