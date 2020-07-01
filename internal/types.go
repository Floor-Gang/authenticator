package internal

import dg "github.com/bwmarrin/discordgo"

// AuthArgs structure.
type AuthArgs struct{ MemberID string }

// AuthResponse structure.
type AuthResponse struct {
	Role    string
	IsAdmin bool
}

// AuthServer structure.
type AuthServer struct {
	config *Config
	client *dg.Session
}
