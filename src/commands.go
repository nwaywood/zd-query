// commands.go contains all the scaffold code required to construct the cli app
// commands and sub-commands
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

// buildCLIApp is a closure to construct the cli commands with the context of the caches
func buildCLIApp(orgMap OrgMap, orgIndexes IndexTypeMap, userMap UserMap, userIndexes IndexTypeMap, ticketMap TicketMap, ticketIndexes IndexTypeMap) *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			// create all user sub-commands
			{
				Name:  "users",
				Usage: "Query users by a particular user value",
				Subcommands: []*cli.Command{
					{
						Name:      "id",
						Usage:     "Query users by id",
						ArgsUsage: "<id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							// get primary search result
							user := userMap[c.Args().First()]
							if user == nil {
								fmt.Println("User not found")
								os.Exit(1)
							}
							// get secondary results
							org, tickets, err := getUserOrgAndTickets(user, orgMap, ticketMap, ticketIndexes["submitter_id"])
							if err != nil {
								log.Fatal(err)
							}
							displayUser(user, org, tickets)

							return nil
						},
					},
					{
						Name:      "org_id",
						Usage:     "Query users by organization_id",
						ArgsUsage: "<org_id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["organization_id"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "url",
						Usage:     "Query users by url",
						ArgsUsage: "<url>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["url"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "external_id",
						Usage:     "Query users by external_id",
						ArgsUsage: "<external_id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["external_id"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "name",
						Usage:     "Query users by name",
						ArgsUsage: "<name>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["name"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "alias",
						Usage:     "Query users by alias",
						ArgsUsage: "<alias>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["alias"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "active",
						Usage:     "Query users by active",
						ArgsUsage: "<active>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["active"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "verified",
						Usage:     "Query users by verified",
						ArgsUsage: "<verified>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["verified"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "shared",
						Usage:     "Query users by shared",
						ArgsUsage: "<shared>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["shared"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "locale",
						Usage:     "Query users by locale",
						ArgsUsage: "<locale>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["locale"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "timezone",
						Usage:     "Query users by timezone",
						ArgsUsage: "<timezone>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["timezone"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "email",
						Usage:     "Query users by email",
						ArgsUsage: "<email>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["email"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "phone",
						Usage:     "Query users by phone",
						ArgsUsage: "<phone>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["phone"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "signature",
						Usage:     "Query users by signature",
						ArgsUsage: "<signature>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["signature"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "tags",
						Usage:     "Query users by tag",
						ArgsUsage: "<tag>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["tags"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "suspended",
						Usage:     "Query users by suspended",
						ArgsUsage: "<suspended>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["suspended"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "role",
						Usage:     "Query users by role",
						ArgsUsage: "<role>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["role"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "created_at",
						Usage:     "Query users by created_at",
						ArgsUsage: "<created_at>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["created_at"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
					{
						Name:      "last_login_at",
						Usage:     "Query users by last_login_at",
						ArgsUsage: "<last_login_at>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryUserAndPrintResults(c.Args().First(), userIndexes["last_login_at"], userMap, orgMap, ticketMap, ticketIndexes["submitter_id"])

							return nil
						},
					},
				},
			},
			// create all organization sub-commands
			{
				Name:  "organizations",
				Usage: "Query organizations by a particular organization value",
				Subcommands: []*cli.Command{
					{
						Name:      "id",
						Usage:     "Query organizations by id",
						ArgsUsage: "<id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							// get primary search result
							org := orgMap[c.Args().First()]
							if org == nil {
								fmt.Println("User not found")
								os.Exit(1)
							}
							// get secondary results
							users, tickets := getOrgUsersAndTickets(org, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])
							displayOrg(org, users, tickets)

							return nil
						},
					},
					{
						Name:      "name",
						Usage:     "Query orgs by name",
						ArgsUsage: "<name>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryOrgAndPrintResults(c.Args().First(), orgIndexes["name"], orgMap, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])

							return nil
						},
					},
					{
						Name:      "url",
						Usage:     "Query orgs by url",
						ArgsUsage: "<url>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryOrgAndPrintResults(c.Args().First(), orgIndexes["url"], orgMap, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])

							return nil
						},
					},
					{
						Name:      "external_id",
						Usage:     "Query orgs by external_id",
						ArgsUsage: "<external_id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryOrgAndPrintResults(c.Args().First(), orgIndexes["external_id"], orgMap, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])

							return nil
						},
					},
					{
						Name:      "domain_names",
						Usage:     "Query orgs by domain_name",
						ArgsUsage: "<domain_name>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryOrgAndPrintResults(c.Args().First(), orgIndexes["domain_names"], orgMap, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])

							return nil
						},
					},
					{
						Name:      "created_at",
						Usage:     "Query orgs by created_at",
						ArgsUsage: "<created_at>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryOrgAndPrintResults(c.Args().First(), orgIndexes["created_at"], orgMap, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])

							return nil
						},
					},
					{
						Name:      "details",
						Usage:     "Query orgs by details",
						ArgsUsage: "<details>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryOrgAndPrintResults(c.Args().First(), orgIndexes["details"], orgMap, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])

							return nil
						},
					},
					{
						Name:      "shared_tickets",
						Usage:     "Query orgs by shared_tickets",
						ArgsUsage: "<shared_tickets>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryOrgAndPrintResults(c.Args().First(), orgIndexes["shared_tickets"], orgMap, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])

							return nil
						},
					},
					{
						Name:      "tags",
						Usage:     "Query orgs by tag",
						ArgsUsage: "<tag>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryOrgAndPrintResults(c.Args().First(), orgIndexes["tags"], orgMap, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])

							return nil
						},
					},
				},
			},
			// create all ticket sub-commands
			{
				Name:  "tickets",
				Usage: "Query tickets by a particular ticket value",
				Subcommands: []*cli.Command{
					{
						Name:      "id",
						Usage:     "Query tickets by id",
						ArgsUsage: "<id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							// get primary search result
							ticket := ticketMap[c.Args().First()]
							if ticket == nil {
								fmt.Println("Ticket not found")
								os.Exit(1)
							}
							// get secondary results
							org, user, err := getTicketOrgAndUser(ticket, orgMap, userMap)
							if err != nil {
								log.Fatal(err)
							}
							displayTicket(ticket, org, user)

							return nil
						},
					},
					{
						Name:      "org_id",
						Usage:     "Query tickets by organization_id",
						ArgsUsage: "<org_id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["organization_id"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "url",
						Usage:     "Query tickets by url",
						ArgsUsage: "<url>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["url"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "submitter_id",
						Usage:     "Query tickets by submitter_id",
						ArgsUsage: "<submitter_id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["submitter_id"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "external_id",
						Usage:     "Query tickets by external_id",
						ArgsUsage: "<external_id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["external_id"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "created_at",
						Usage:     "Query tickets by created_at",
						ArgsUsage: "<created_at>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["created_at"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "type",
						Usage:     "Query tickets by type",
						ArgsUsage: "<type>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["type"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "subject",
						Usage:     "Query tickets by subject",
						ArgsUsage: "<subject>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["subject"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "description",
						Usage:     "Query tickets by description",
						ArgsUsage: "<description>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["description"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "priority",
						Usage:     "Query tickets by priority",
						ArgsUsage: "<priority>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["priority"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "status",
						Usage:     "Query tickets by status",
						ArgsUsage: "<status>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["status"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "assignee_id",
						Usage:     "Query tickets by assignee_id",
						ArgsUsage: "<assignee_id>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["assignee_id"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "tags",
						Usage:     "Query tickets by tag",
						ArgsUsage: "<tag>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["tags"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "has_incidents",
						Usage:     "Query tickets by has_incidents",
						ArgsUsage: "<has_incidents>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["has_incidents"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "due_at",
						Usage:     "Query tickets by due_at",
						ArgsUsage: "<due_at>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["due_at"], ticketMap, orgMap, userMap)

							return nil
						},
					},
					{
						Name:      "via",
						Usage:     "Query tickets by via",
						ArgsUsage: "<via>",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							queryTicketAndPrintResults(c.Args().First(), ticketIndexes["via"], ticketMap, orgMap, userMap)

							return nil
						},
					},
				},
			},
		},
	}
}

func validateArgs(c *cli.Context) {
	invalidArgsString := "Invalid number of args. Use --help or -h for usage."
	if c.Args().Len() != 1 {
		fmt.Println(invalidArgsString)
		os.Exit(1)
	}
}

// queryUserAndPrintResults is a helper function used by `users` commands to print results to stdout
func queryUserAndPrintResults(query string, fieldIndexMap IndexMap, userMap UserMap, orgMap OrgMap, ticketMap TicketMap, submitterIDIndexMap IndexMap) {
	// get primary search result
	users := getUsersByIndex(query, fieldIndexMap, userMap)
	if len(users) == 0 {
		fmt.Println("No users found")
		os.Exit(1)
	}
	// get secondary results
	errs := []string{}
	for _, user := range users {
		org, tickets, err := getUserOrgAndTickets(user, orgMap, ticketMap, submitterIDIndexMap)
		if err != nil {
			errs = append(errs, err.Error())
		} else {
			displayUser(user, org, tickets)
		}
	}
	// print any errors encountered at the end of the output to ensure they are seen
	for _, e := range errs {
		color.Red(e)
	}
}

// queryOrgAndPrintResults is a helper function used by `organizations` commands to print results to stdout
func queryOrgAndPrintResults(query string, fieldIndexMap IndexMap, orgMap OrgMap, userMap UserMap, ticketMap TicketMap, ticketOrgIDMap IndexMap, userOrgIDMap IndexMap) {
	// get primary search result
	orgs := getOrgsByIndex(query, fieldIndexMap, orgMap)
	if len(orgs) == 0 {
		fmt.Println("No orgs found")
		os.Exit(1)
	}
	// get secondary results
	for _, org := range orgs {
		users, tickets := getOrgUsersAndTickets(org, userMap, ticketMap, ticketOrgIDMap, userOrgIDMap)
		displayOrg(org, users, tickets)
	}
}

// queryTicketAndPrintResults is a helper function used by `tickets` commands to print results to stdout
func queryTicketAndPrintResults(query string, fieldIndexMap IndexMap, ticketMap TicketMap, orgMap OrgMap, userMap UserMap) {
	// get primary search result
	tickets := getTicketsByIndex(query, fieldIndexMap, ticketMap)
	if len(tickets) == 0 {
		fmt.Println("No Tickets found")
		os.Exit(1)
	}
	// get secondary results
	errs := []string{}
	for _, ticket := range tickets {
		org, user, err := getTicketOrgAndUser(ticket, orgMap, userMap)
		if err != nil {
			errs = append(errs, err.Error())
		} else {
			displayTicket(ticket, org, user)
		}
	}
	// print any errors encountered at the end of the output to ensure they are seen
	for _, e := range errs {
		color.Red(e)
	}
}
