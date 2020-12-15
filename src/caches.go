// caches.go contains the helper functions and data structures for
// creating the caches from the source data
//
// These caches are created to allow for O(1) search times
package main

import (
	"strconv"
)

// IndexTypeMap is a generic type to hold all the secondary
// indexes for a particular domain type (user, ticket or org)
type IndexTypeMap map[string]IndexMap

// IndexMap is a generic type to represent the secondary
// index of a domain type's field (e.g. name, details etc.)
// The value is []string since multiple values can match an index
type IndexMap map[string][]string

// OrgMap is a map of all Organizations keyed by ID
type OrgMap map[string]*organization

// UserMap is a map of all Users keyed by ID
type UserMap map[string]*user

// TicketMap is a map of all Tickets keyed by ID
type TicketMap map[string]*ticket

func buildOrgCaches(orgs []*organization) (OrgMap, IndexTypeMap) {
	orgMap := make(OrgMap)
	// initialise index maps
	indexTypeMap := make(IndexTypeMap)
	indexTypeMap["name"] = make(IndexMap)
	indexTypeMap["url"] = make(IndexMap)
	indexTypeMap["external_id"] = make(IndexMap)
	indexTypeMap["domain_names"] = make(IndexMap)
	indexTypeMap["created_at"] = make(IndexMap)
	indexTypeMap["details"] = make(IndexMap)
	indexTypeMap["shared_tickets"] = make(IndexMap)
	indexTypeMap["tags"] = make(IndexMap)

	for _, org := range orgs {
		orgID := org.ID.String()
		// build orgMap
		orgMap[orgID] = org

		// build indexes
		indexTypeMap["name"][org.Name] = append(indexTypeMap["name"][org.Name], orgID)
		indexTypeMap["url"][org.URL] = append(indexTypeMap["url"][org.URL], orgID)
		indexTypeMap["external_id"][org.ExternalID] = append(indexTypeMap["external_id"][org.ExternalID], orgID)
		indexTypeMap["created_at"][org.CreatedAt] = append(indexTypeMap["created_at"][org.CreatedAt], orgID)
		indexTypeMap["details"][org.Details] = append(indexTypeMap["details"][org.Details], orgID)
		indexTypeMap["shared_tickets"][strconv.FormatBool(org.SharedTickets)] = append(indexTypeMap["shared_tickets"][strconv.FormatBool(org.SharedTickets)], orgID)
		for _, tag := range org.Tags {
			indexTypeMap["tags"][tag] = append(indexTypeMap["tags"][tag], orgID)
		}
		for _, dn := range org.DomainNames {
			indexTypeMap["domain_names"][dn] = append(indexTypeMap["domain_names"][dn], orgID)
		}
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
	// initialise index maps
	indexTypeMap := make(IndexTypeMap)
	indexTypeMap["organization_id"] = make(IndexMap)
	indexTypeMap["submitter_id"] = make(IndexMap)
	indexTypeMap["url"] = make(IndexMap)
	indexTypeMap["external_id"] = make(IndexMap)
	indexTypeMap["created_at"] = make(IndexMap)
	indexTypeMap["subject"] = make(IndexMap)
	indexTypeMap["description"] = make(IndexMap)
	indexTypeMap["priority"] = make(IndexMap)
	indexTypeMap["status"] = make(IndexMap)
	indexTypeMap["assignee_id"] = make(IndexMap)
	indexTypeMap["tags"] = make(IndexMap)
	indexTypeMap["has_incidents"] = make(IndexMap)
	indexTypeMap["due_at"] = make(IndexMap)
	indexTypeMap["via"] = make(IndexMap)

	for _, ticket := range tickets {
		// build ticketMap
		ticketMap[ticket.ID] = ticket

		// build indexes
		indexTypeMap["organization_id"][ticket.OrganizationID.String()] = append(indexTypeMap["organization_id"][ticket.OrganizationID.String()], ticket.ID)
		indexTypeMap["submitter_id"][ticket.SubmitterID.String()] = append(indexTypeMap["submitter_id"][ticket.SubmitterID.String()], ticket.ID)
		indexTypeMap["assignee_id"][ticket.AssigneeID.String()] = append(indexTypeMap["assignee_id"][ticket.AssigneeID.String()], ticket.ID)
		indexTypeMap["url"][ticket.URL] = append(indexTypeMap["url"][ticket.URL], ticket.ID)
		indexTypeMap["external_id"][ticket.ExternalID] = append(indexTypeMap["external_id"][ticket.ExternalID], ticket.ID)
		indexTypeMap["created_at"][ticket.CreatedAt] = append(indexTypeMap["created_at"][ticket.CreatedAt], ticket.ID)
		indexTypeMap["subject"][ticket.Subject] = append(indexTypeMap["subject"][ticket.Subject], ticket.ID)
		indexTypeMap["description"][ticket.Description] = append(indexTypeMap["description"][ticket.Description], ticket.ID)
		indexTypeMap["priority"][ticket.Priority] = append(indexTypeMap["priority"][ticket.Priority], ticket.ID)
		indexTypeMap["status"][ticket.Status] = append(indexTypeMap["status"][ticket.Status], ticket.ID)
		indexTypeMap["due_at"][ticket.DueAt] = append(indexTypeMap["due_at"][ticket.DueAt], ticket.ID)
		indexTypeMap["via"][ticket.Via] = append(indexTypeMap["via"][ticket.Via], ticket.ID)
		indexTypeMap["has_incidents"][strconv.FormatBool(ticket.HasIncidents)] = append(indexTypeMap["has_incidents"][strconv.FormatBool(ticket.HasIncidents)], ticket.ID)
		for _, tag := range ticket.Tags {
			indexTypeMap["tags"][tag] = append(indexTypeMap["tags"][tag], ticket.ID)
		}
	}
	return ticketMap, indexTypeMap
}
