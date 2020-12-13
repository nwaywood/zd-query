// commands.go contains all the scaffold code required to construct the cli app
package main

import (
	"fmt"
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
						Name:  "id",
						Usage: "Query users by id",
						Action: func(c *cli.Context) error {
							validateArgs(c)

							// get main object (from id lookup or from index)
							user := userMap[c.Args().First()]
							if user == nil {
								fmt.Println("User not found")
								os.Exit(1)
							}
							// get secondary objects
							org, tickets := getUserOrgAndTickets(user, orgMap, ticketMap, ticketIndexes["submitter_id"])
							// call display user/org/ticket
							displayUser(user, org, tickets)
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
						Name:  "add",
						Usage: "add a new template",
						Action: func(c *cli.Context) error {
							fmt.Println("new task template: ", c.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
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
						Name:  "add",
						Usage: "add a new template",
						Action: func(c *cli.Context) error {
							fmt.Println("new task template: ", c.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
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
