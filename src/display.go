// display.go contains functions for pretty printing results
// from queries in a human readable format
package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// displayUser prints a human readable format of user with its
// associated org and tickets
func displayUser(user *user, org *organization, tickets []*ticket) {
	color.Green("====")
	color.Green("USER")
	color.Green("====")
	fmt.Printf("id = %s, name = %s, alias = %s\n", user.ID, user.Name, user.Alias)
	fmt.Printf("url = %s\n", user.URL)
	fmt.Printf("external id = %s\n", user.ExternalID)
	fmt.Printf("created at = %s\n", user.CreatedAt)
	fmt.Printf("active = %v, verified = %v, shared = %v\n", user.Active, user.Verified, user.Shared)
	fmt.Printf("locale = %s, timezone = %s\n", user.Locale, user.Timezone)
	fmt.Printf("last login at = %s\n", user.LastLoginAt)
	fmt.Printf("email = %s\n", user.Email)
	fmt.Printf("phone = %s\n", user.Phone)
	fmt.Printf("signature = %s\n", user.Signature)
	fmt.Printf("tags = %v\n", user.Tags)
	fmt.Printf("suspended = %v, role = %s\n", user.Suspended, user.Role)
	fmt.Println("ORGANIZATION:")
	fmt.Printf("id = %s, name = %s\n", org.ID, org.Name)
	fmt.Println("TICKETS:")
	if len(tickets) == 0 {
		fmt.Println("No Tickets for User")
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Status", "Subject"})
		for _, t := range tickets {
			table.Append([]string{t.Status, t.Subject})
		}
		table.Render()
	}
	fmt.Println("")
}

// displayOrg prints a human readable format of org with its
// associated users and tickets
func displayOrg(org *organization, users []*user, tickets []*ticket) {
	color.Green("============")
	color.Green("ORGANIZATION")
	color.Green("============")
	fmt.Printf("id = %s, name = %s\n", org.ID, org.Name)
	fmt.Printf("details = %s\n", org.Details)
	fmt.Printf("url = %s\n", org.URL)
	fmt.Printf("external id = %s\n", org.ExternalID)
	fmt.Printf("created at = %s\n", org.CreatedAt)
	fmt.Printf("shared tickets = %v\n", org.SharedTickets)
	fmt.Printf("domain names = %s\n", org.DomainNames)
	fmt.Printf("tags = %s\n", org.Tags)
	fmt.Println("USERS:")
	if len(users) == 0 {
		fmt.Println("No Users for Organization")
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name"})
		for _, u := range users {
			table.Append([]string{u.ID.String(), u.Name})
		}
		table.Render()
	}
	fmt.Println("TICKETS:")
	if len(tickets) == 0 {
		fmt.Println("No Tickets for Organization")
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Status", "Subject"})
		for _, t := range tickets {
			table.Append([]string{t.Status, t.Subject})
		}
		table.Render()
	}
	fmt.Println("")
}

// displayTicket prints a human readable format of ticket with its
// associated org and user
func displayTicket(ticket *ticket, org *organization, user *user) {
	color.Green("======")
	color.Green("TICKET")
	color.Green("======")
	fmt.Printf("id = %s, type = %s\n", ticket.ID, ticket.Type)
	fmt.Printf("status = %s, priority = %s\n", ticket.Status, ticket.Priority)
	fmt.Printf("subject = %s\n", ticket.Subject)
	fmt.Printf("description = %s\n", ticket.Description)
	fmt.Printf("Submitter id = %s, assignee_id = %s\n", ticket.SubmitterID, ticket.AssigneeID)
	fmt.Printf("url = %s\n", ticket.URL)
	fmt.Printf("external id = %s\n", ticket.ExternalID)
	fmt.Printf("created at = %s\n", ticket.CreatedAt)
	fmt.Printf("due at = %s\n", ticket.DueAt)
	fmt.Printf("has incidents = %v\n", ticket.HasIncidents)
	fmt.Printf("tags = %v\n", ticket.Tags)
	fmt.Println("ORGANIZATION:")
	fmt.Printf("id = %s, name = %s\n", org.ID, org.Name)
	fmt.Println("USER:")
	fmt.Printf("id = %s, name = %s\n", user.ID, user.Name)
	fmt.Println("")
}
