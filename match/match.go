package match

// TODO add struct that holds regexp objects, and gives references to those objects
// instead of creating new regex objects based on strings from this package

// Username returns raw string representing regex pattern to match usernames
func Username() string {
	return `^[a-zA-Z0-9_-]+`
}

// Email returns raw string representing regex pattern to match emails
func Email() string {
	return `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`
}

// Tag returns raw string representing regex pattern to match tags (e.g. "#hashtag")
func Tag() string {
	return `#[a-zA-Z0-9_]+`
}

// At returns raw string representing regex pattern to match ats (e.g. "@user")
func At() string {
	return `@[a-zA-Z0-9_-]+`
}
