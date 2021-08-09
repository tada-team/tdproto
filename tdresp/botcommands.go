package tdresp

// Bot commands list
type BotCommands []BotCommand

// Bot commands information
type BotCommand struct {
	// What should be inserted to the chat
	Key string `json:"key"`

	// What should be visible by user
	Title string `json:"title"`

	// Command options, if any
	Args []string `json:"args"`
}
