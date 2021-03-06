// types.go contains all the domain types used by the application
package main

import "encoding/json"

type organization struct {
	// json.Number to cleanse input data
	// https://stackoverflow.com/questions/24480835/decoding-json-int-into-string
	ID            json.Number `json:"_id"`
	URL           string      `json:"url"`
	ExternalID    string      `json:"external_id"`
	Name          string      `json:"name"`
	DomainNames   []string    `json:"domain_names"`
	CreatedAt     string      `json:"created_at"`
	Details       string      `json:"details"`
	SharedTickets bool        `json:"shared_tickets"`
	Tags          []string    `json:"tags"`
}

type user struct {
	ID             json.Number `json:"_id"`
	URL            string      `json:"url"`
	ExternalID     string      `json:"external_id"`
	Name           string      `json:"name"`
	Alias          string      `json:"alias"`
	CreatedAt      string      `json:"created_at"`
	Active         bool        `json:"active"`
	Verified       bool        `json:"verified"`
	Shared         bool        `json:"shared"`
	Locale         string      `json:"locale"`
	Timezone       string      `json:"timezone"`
	LastLoginAt    string      `json:"last_login_at"`
	Email          string      `json:"email"`
	Phone          string      `json:"phone"`
	Signature      string      `json:"signature"`
	OrganizationID json.Number `json:"organization_id"`
	Tags           []string    `json:"tags"`
	Suspended      bool        `json:"suspended"`
	Role           string      `json:"role"`
}

type ticket struct {
	ID             string      `json:"_id"`
	URL            string      `json:"url"`
	ExternalID     string      `json:"external_id"`
	CreatedAt      string      `json:"created_at"`
	Type           string      `json:"type"`
	Subject        string      `json:"subject"`
	Description    string      `json:"description"`
	Priority       string      `json:"priority"`
	Status         string      `json:"status"`
	SubmitterID    json.Number `json:"submitter_id"`
	AssigneeID     json.Number `json:"assignee_id"`
	OrganizationID json.Number `json:"organization_id"`
	Tags           []string    `json:"tags"`
	HasIncidents   bool        `json:"has_incidents"`
	DueAt          string      `json:"due_at"`
	Via            string      `json:"via"`
}
