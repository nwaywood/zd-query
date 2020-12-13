// caches.go contains the helper functions and data structure for
// creating the caches from the source data
//
// These caches are created to allow for O(1) search times
package main

// IndexMap is a generic type to represent the secondary
// index of a domain type's field
// The value is []string since multiple values can match an index
type IndexMap map[string][]string

// IndexTypeMap is a generic type to hold all the secondary
// indexes for a particular domain type
type IndexTypeMap map[string]IndexMap

// OrgMap is a map of all Organizations keyed to ID
type OrgMap map[string]*organization

// UserMap is a map of all Users keyed to ID
type UserMap map[string]*user

// TicketMap is a map of all Tickets keyed to ID
type TicketMap map[string]*ticket

func buildOrgCaches(orgs []*organization) (OrgMap, IndexTypeMap) {
	orgMap := make(OrgMap)
	indexTypeMap := make(IndexTypeMap)
	indexTypeMap["name"] = make(IndexMap)

	for _, org := range orgs {
		orgID := org.ID.String()
		// build orgMap
		orgMap[orgID] = org

		// build indexes
		indexTypeMap["name"][org.Name] = append(indexTypeMap["name"][org.Name], orgID)
	}
	return orgMap, indexTypeMap
}

func buildUserCaches(users []*user) (UserMap, IndexTypeMap) {
	userMap := make(UserMap)
	indexTypeMap := make(IndexTypeMap)
	indexTypeMap["organization_id"] = make(IndexMap)

	for _, user := range users {
		userID := user.ID.String()
		// build userMap
		userMap[userID] = user

		// build indexes
		indexTypeMap["organization_id"][user.OrganizationID.String()] = append(indexTypeMap["organization_id"][user.OrganizationID.String()], userID)
	}
	return userMap, indexTypeMap

}
func buildTicketCaches(tickets []*ticket) (TicketMap, IndexTypeMap) {
	ticketMap := make(TicketMap)
	indexTypeMap := make(IndexTypeMap)
	indexTypeMap["organization_id"] = make(IndexMap)
	indexTypeMap["submitter_id"] = make(IndexMap)

	for _, ticket := range tickets {
		// build ticketMap
		ticketMap[ticket.ID] = ticket

		// build indexes
		indexTypeMap["organization_id"][ticket.OrganizationID.String()] = append(indexTypeMap["organization_id"][ticket.OrganizationID.String()], ticket.ID)
		indexTypeMap["submitter_id"][ticket.SubmitterID.String()] = append(indexTypeMap["submitter_id"][ticket.SubmitterID.String()], ticket.ID)
	}
	return ticketMap, indexTypeMap
}
