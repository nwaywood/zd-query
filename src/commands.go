// commands.go contains all the scaffold code required to construct the cli app
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func buildCLIApp(orgMap OrgMap, orgIndexes IndexTypeMap, userMap UserMap, userIndexes IndexTypeMap, ticketMap TicketMap, ticketIndexes IndexTypeMap) *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
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

							// get primary search result
							tickets := getTicketsByIndex(c.Args().First(), ticketIndexes["organization_id"], ticketMap)
							if len(tickets) == 0 {
								fmt.Println("No Tickets found")
								os.Exit(1)
							}
							// get secondary results
							for _, ticket := range tickets {
								org, user, err := getTicketOrgAndUser(ticket, orgMap, userMap)
								if err != nil {
									log.Fatal(err)
								}
								displayTicket(ticket, org, user)
							}
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

// queryUserAndPrintResults is a helper function used by `users` queries to print results
func queryUserAndPrintResults(query string, verifiedIndexMap IndexMap, userMap UserMap, orgMap OrgMap, ticketMap TicketMap, submitterIDIndexMap IndexMap) {
	// get primary search result
	users := getUsersByIndex(query, verifiedIndexMap, userMap)
	if len(users) == 0 {
		fmt.Println("No users found")
		os.Exit(1)
	}
	// get secondary results
	// TODO print errors at the end
	for _, user := range users {
		org, tickets, err := getUserOrgAndTickets(user, orgMap, ticketMap, submitterIDIndexMap)
		if err != nil {
			fmt.Println(err)
		} else {
			displayUser(user, org, tickets)
		}
	}
}

func queryOrgAndPrintResults(query string, orgNameIndexMap IndexMap, orgMap OrgMap, userMap UserMap, ticketMap TicketMap, ticketOrgIDMap IndexMap, userOrgIDMap IndexMap) {
	// get primary search result
	orgs := getOrgsByIndex(query, orgNameIndexMap, orgMap)
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
