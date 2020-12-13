package main

import "fmt"

// displayUser prints a human readable format of user
func displayUser(user *user, org *organization, tickets []*ticket) {
	fmt.Println("====")
	fmt.Println("USER")
	fmt.Println("====")
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
	fmt.Println("User Organization:")
	fmt.Println("------------------")
	fmt.Printf("id = %s, name = %s\n", org.ID, org.Name)
	fmt.Println("User Tickets:")
	fmt.Println("-------------")
	if len(tickets) == 0 {
		fmt.Println("No tickets for User")
	} else {
		for _, t := range tickets {
			fmt.Printf("- %s\n", t.Subject)
		}
	}
}
