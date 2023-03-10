package methods

type JournalsCategory struct {
	BasisAPI BasisAPIInterface
}

// GetChatHistory returns the chat message history
func (c JournalsCategory) GetChatHistory(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "GetChatHistory", parameters, "")
}

// GetMessage returns a chat message
func (c JournalsCategory) GetMessage(chatId, idMessage string) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "getMessage", map[string]interface{}{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}

// LastIncomingMessages returns the most recent incoming messages of the account
func (c JournalsCategory) LastIncomingMessages(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "lastIncomingMessages", parameters, "")
}

// LastOutgoingMessages returns the last sent messages of the account
func (c JournalsCategory) LastOutgoingMessages(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "LastOutgoingMessages", parameters, "")
}
