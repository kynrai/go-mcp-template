package repo

import "fmt"

func Images(user string, subscribed bool) string {
	return fmt.Sprintf("images for %s, paid status %v", user, subscribed)
}

func Videos(user string, subscribed bool) string {
	return fmt.Sprintf("videos for %s, paid status %v", user, subscribed)
}
