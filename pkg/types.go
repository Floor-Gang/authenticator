package pkg

// Registration.

// Feature to register.
type Feature struct {
	Name          string       // ie. ".admin"
	Description   string       // "For managing administrator"
	Commands      []SubCommand // Optionally empty
	CommandPrefix string       // Optionally empty
}

// Commands that come with this feature.
type SubCommand struct {
	Name        string   // ie. "add"
	Description string   // ie. "This is for adding administrator roles"
	Example     []string // ie. [".admin", "add", "...role ID's"]
}

// Response from the server while registering.
type RegisterResponse struct {
	Serving string // The ID of the Guild being served
	Token   string // The Discord bot access token
}

// Authentication

// AuthArgs structure.
type AuthArgs struct{ MemberID string }

// AuthResponse structure.
type AuthResponse struct {
	Role    string
	IsAdmin bool
}
