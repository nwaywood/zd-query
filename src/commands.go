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

							// get primary search result
							users := getUsersByIndex(c.Args().First(), userIndexes["organization_id"], userMap)
							if len(users) == 0 {
								fmt.Println("No users found")
								os.Exit(1)
							}
							// get secondary results
							for _, user := range users {
								org, tickets, err := getUserOrgAndTickets(user, orgMap, ticketMap, ticketIndexes["submitter_id"])
								if err != nil {
									log.Fatal(err)
								}
								displayUser(user, org, tickets)
							}
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

							// get primary search result
							orgs := getOrgsByIndex(c.Args().First(), orgIndexes["name"], orgMap)
							if len(orgs) == 0 {
								fmt.Println("No orgs found")
								os.Exit(1)
							}
							// get secondary results
							for _, org := range orgs {
								users, tickets := getOrgUsersAndTickets(org, userMap, ticketMap, ticketIndexes["organization_id"], userIndexes["organization_id"])
								displayOrg(org, users, tickets)
							}
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
