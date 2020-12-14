// caches.go contains the helper functions and data structure for
// creating the caches from the source data
//
// These caches are created to allow for O(1) search times
package main

import "strconv"

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
	// initialise index maps
	indexTypeMap := make(IndexTypeMap)
	indexTypeMap["organization_id"] = make(IndexMap)
	indexTypeMap["url"] = make(IndexMap)
	indexTypeMap["external_id"] = make(IndexMap)
	indexTypeMap["name"] = make(IndexMap)
	indexTypeMap["alias"] = make(IndexMap)
	indexTypeMap["active"] = make(IndexMap)
	indexTypeMap["verified"] = make(IndexMap)
	indexTypeMap["shared"] = make(IndexMap)
	indexTypeMap["locale"] = make(IndexMap)
	indexTypeMap["timezone"] = make(IndexMap)
	indexTypeMap["email"] = make(IndexMap)
	indexTypeMap["phone"] = make(IndexMap)
	indexTypeMap["signature"] = make(IndexMap)
	indexTypeMap["tags"] = make(IndexMap)
	indexTypeMap["suspended"] = make(IndexMap)
	indexTypeMap["role"] = make(IndexMap)
	indexTypeMap["created_at"] = make(IndexMap)
	indexTypeMap["last_login_at"] = make(IndexMap)

	for _, user := range users {
		userID := user.ID.String()
		// build userMap
		userMap[userID] = user

		// build indexes
		indexTypeMap["organization_id"][user.OrganizationID.String()] = append(indexTypeMap["organization_id"][user.OrganizationID.String()], userID)
		indexTypeMap["url"][user.URL] = append(indexTypeMap["url"][user.URL], userID)
		indexTypeMap["external_id"][user.ExternalID] = append(indexTypeMap["external_id"][user.ExternalID], userID)
		indexTypeMap["name"][user.Name] = append(indexTypeMap["name"][user.Name], userID)
		indexTypeMap["alias"][user.Alias] = append(indexTypeMap["alias"][user.Alias], userID)
		indexTypeMap["active"][strconv.FormatBool(user.Active)] = append(indexTypeMap["active"][strconv.FormatBool(user.Active)], userID)
		indexTypeMap["verified"][strconv.FormatBool(user.Verified)] = append(indexTypeMap["verified"][strconv.FormatBool(user.Verified)], userID)
		indexTypeMap["shared"][strconv.FormatBool(user.Shared)] = append(indexTypeMap["shared"][strconv.FormatBool(user.Shared)], userID)
		indexTypeMap["locale"][user.Locale] = append(indexTypeMap["locale"][user.Locale], userID)
		indexTypeMap["timezone"][user.Timezone] = append(indexTypeMap["timezone"][user.Timezone], userID)
		indexTypeMap["email"][user.Email] = append(indexTypeMap["email"][user.Email], userID)
		indexTypeMap["phone"][user.Phone] = append(indexTypeMap["phone"][user.Phone], userID)
		indexTypeMap["signature"][user.Signature] = append(indexTypeMap["signature"][user.Signature], userID)
		indexTypeMap["suspended"][strconv.FormatBool(user.Suspended)] = append(indexTypeMap["suspended"][strconv.FormatBool(user.Suspended)], userID)
		indexTypeMap["role"][user.Role] = append(indexTypeMap["role"][user.Role], userID)
		indexTypeMap["created_at"][user.CreatedAt] = append(indexTypeMap["created_at"][user.CreatedAt], userID)
		indexTypeMap["last_login_at"][user.LastLoginAt] = append(indexTypeMap["last_login_at"][user.LastLoginAt], userID)
		for _, tag := range user.Tags {
			indexTypeMap["tags"][tag] = append(indexTypeMap["tags"][tag], userID)
		}
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
