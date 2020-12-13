// caches.go contains the helper functions and data structure for
// creating the caches from the source data
//
// These caches are created to allow for O(1) search times
package main

// IndexMap is a generic type to represent the secondary
// index of a type
type IndexMap map[string]string

// IndexTypeMap is a generic type to hold all the secondary
// indexes for a particular domain type
type IndexTypeMap map[string]IndexMap

// OrgMap is a map of all Organizations keyed to ID
type OrgMap map[string]*Organization

// UserMap is a map of all Users keyed to ID
type UserMap map[string]*User

// TicketMap is a map of all Tickets keyed to ID
type TicketMap map[string]*Ticket

func buildOrgCaches(orgs []*Organization) (OrgMap, IndexTypeMap) {
	orgMap := make(OrgMap)
	indexTypeMap := make(IndexTypeMap)

	for _, org := range orgs {
		orgID := org.ID.String()
		// build orgMap
		orgMap[orgID] = org

		// build indexes
		indexTypeMap["name"][org.Name] = orgID
	}
	return orgMap, indexTypeMap
}

func buildUserCaches(users []*User) (UserMap, IndexTypeMap) {
	userMap := make(UserMap)
	indexTypeMap := make(IndexTypeMap)

	for _, user := range users {
		userID := user.ID.String()
		// build userMap
		userMap[userID] = user

		// build indexes
		indexTypeMap["organization_id"][user.OrganizationID.String()] = userID
	}
	return userMap, indexTypeMap

}
func buildTicketCaches(tickets []*Ticket) (TicketMap, IndexTypeMap) {
	ticketMap := make(TicketMap)
	indexTypeMap := make(IndexTypeMap)

	for _, ticket := range tickets {
		// build ticketMap
		ticketMap[ticket.ID] = ticket

		// build indexes
		indexTypeMap["organization_id"][ticket.OrganizationID.String()] = ticket.ID
		indexTypeMap["submitter_id"][ticket.SubmitterID.String()] = ticket.ID
	}
	return ticketMap, indexTypeMap
}
