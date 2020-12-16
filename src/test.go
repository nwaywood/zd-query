// test.go contains values used by multiple test files
package main

var orgs = []*organization{
	{
		ID:            "1",
		URL:           "www.zendesk.com",
		ExternalID:    "ex1",
		Name:          "zendesk",
		DomainNames:   []string{""},
		CreatedAt:     "",
		SharedTickets: false,
		Tags:          []string{"tech", "customer"},
	},
	{
		ID:            "2",
		URL:           "www.google.com",
		ExternalID:    "ex2",
		Name:          "google",
		DomainNames:   []string{""},
		CreatedAt:     "Tuesday",
		SharedTickets: false,
		Tags:          []string{},
	},
	{
		ID:            "3",
		URL:           "www.sgc.com",
		ExternalID:    "ex3",
		Name:          "stargate command",
		DomainNames:   []string{""},
		CreatedAt:     "",
		SharedTickets: true,
		Tags:          []string{"space", "tech"},
	},
}

var users = []*user{
	{
		ID:             "101",
		URL:            "www.stargate.com",
		ExternalID:     "ex101",
		Name:           "Samantha Carter",
		Active:         true,
		OrganizationID: "3",
		Tags:           []string{"space", "science"},
	},
	{
		ID:             "102",
		URL:            "www.stargate.com",
		ExternalID:     "ex102",
		Name:           "Jack O'Neill",
		Active:         true,
		OrganizationID: "3",
		Tags:           []string{"space", "fishing"},
	},
}

var tickets = []*ticket{
	{
		ID:             "t1",
		URL:            "www.stargate.com",
		ExternalID:     "ext1",
		OrganizationID: "3",
		Priority:       "high",
		Description:    "capacitors have overloaded!",
		Tags:           []string{"space", "science"},
		SubmitterID:    "101",
	},
	{
		ID:             "t2",
		URL:            "www.stargate.com",
		ExternalID:     "ext2",
		Priority:       "low",
		Description:    "go on fishing trip",
		OrganizationID: "3",
		SubmitterID:    "102",
		Tags:           []string{"space", "fishing"},
	},
	{
		ID:             "t3",
		URL:            "www.stargate.com",
		ExternalID:     "ext3",
		Priority:       "high",
		Description:    "become general",
		OrganizationID: "3",
		SubmitterID:    "102",
		Tags:           []string{},
	},
}

var expectedOrgMap = OrgMap{"1": orgs[0], "2": orgs[1], "3": orgs[2]}

var expectedOrgIndexMap = IndexTypeMap{
	"created_at":     IndexMap{"": []string{"1", "3"}, "Tuesday": []string{"2"}},
	"details":        IndexMap{"": []string{"1", "2", "3"}},
	"domain_names":   IndexMap{"": []string{"1", "2", "3"}},
	"external_id":    IndexMap{"ex1": []string{"1"}, "ex2": []string{"2"}, "ex3": []string{"3"}},
	"name":           IndexMap{"stargate command": []string{"3"}, "google": []string{"2"}, "zendesk": []string{"1"}},
	"shared_tickets": IndexMap{"false": []string{"1", "2"}, "true": []string{"3"}},
	"tags":           IndexMap{"space": []string{"3"}, "customer": []string{"1"}, "tech": []string{"1", "3"}},
	"url":            IndexMap{"www.sgc.com": []string{"3"}, "www.google.com": []string{"2"}, "www.zendesk.com": []string{"1"}},
}

var expectedUserMap = UserMap{"101": users[0], "102": users[1]}

var expectedUserIndexMap = IndexTypeMap{
	"active":          IndexMap{"true": []string{"101", "102"}},
	"alias":           IndexMap{"": []string{"101", "102"}},
	"created_at":      IndexMap{"": []string{"101", "102"}},
	"email":           IndexMap{"": []string{"101", "102"}},
	"external_id":     IndexMap{"ex101": []string{"101"}, "ex102": []string{"102"}},
	"last_login_at":   IndexMap{"": []string{"101", "102"}},
	"locale":          IndexMap{"": []string{"101", "102"}},
	"name":            IndexMap{"Jack O'Neill": []string{"102"}, "Samantha Carter": []string{"101"}},
	"organization_id": IndexMap{"3": []string{"101", "102"}},
	"phone":           IndexMap{"": []string{"101", "102"}},
	"role":            IndexMap{"": []string{"101", "102"}},
	"shared":          IndexMap{"false": []string{"101", "102"}},
	"signature":       IndexMap{"": []string{"101", "102"}},
	"suspended":       IndexMap{"false": []string{"101", "102"}},
	"tags":            IndexMap{"fishing": []string{"102"}, "science": []string{"101"}, "space": []string{"101", "102"}},
	"timezone":        IndexMap{"": []string{"101", "102"}},
	"url":             IndexMap{"www.stargate.com": []string{"101", "102"}},
	"verified":        IndexMap{"false": []string{"101", "102"}},
}

var expectedTicketMap = TicketMap{"t1": tickets[0], "t2": tickets[1], "t3": tickets[2]}

var expectedTicketIndexMap = IndexTypeMap{
	"assignee_id":     IndexMap{"": []string{"t1", "t2", "t3"}},
	"created_at":      IndexMap{"": []string{"t1", "t2", "t3"}},
	"description":     IndexMap{"become general": []string{"t3"}, "capacitors have overloaded!": []string{"t1"}, "go on fishing trip": []string{"t2"}},
	"due_at":          IndexMap{"": []string{"t1", "t2", "t3"}},
	"external_id":     IndexMap{"ext1": []string{"t1"}, "ext2": []string{"t2"}, "ext3": []string{"t3"}},
	"has_incidents":   IndexMap{"false": []string{"t1", "t2", "t3"}},
	"organization_id": IndexMap{"3": []string{"t1", "t2", "t3"}},
	"priority":        IndexMap{"high": []string{"t1", "t3"}, "low": []string{"t2"}},
	"status":          IndexMap{"": []string{"t1", "t2", "t3"}},
	"subject":         IndexMap{"": []string{"t1", "t2", "t3"}},
	"submitter_id":    IndexMap{"101": []string{"t1"}, "102": []string{"t2", "t3"}},
	"tags":            IndexMap{"fishing": []string{"t2"}, "science": []string{"t1"}, "space": []string{"t1", "t2"}},
	"url":             IndexMap{"www.stargate.com": []string{"t1", "t2", "t3"}},
	"via":             IndexMap{"": []string{"t1", "t2", "t3"}}}
