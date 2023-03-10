package methods

type SendingCategory struct {
	BasisAPI BasisAPIInterface
}

// SendMessage is designed to send a text message to a personal or group chat
func (c SendingCategory) SendMessage(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "SendMessage", parameters, "")
}

// SendButtons is designed to send a message with buttons to a personal or group chat
func (c SendingCategory) SendButtons(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "SendButtons", parameters, "")
}

// SendTemplateButtons is designed to send a message with interactive buttons from the list of templates in a personal or group chat
func (c SendingCategory) SendTemplateButtons(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "sendTemplateButtons", parameters, "")
}

// SendListMessage is designed to send a message with a selection button from a list of values to a personal or group chat
func (c SendingCategory) SendListMessage(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "SendListMessage", parameters, "")
}

// SendFileByUpload is designed to send a file loaded through a form (form-data)
func (c SendingCategory) SendFileByUpload(filePath string, parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "SendFileByUpload", parameters, filePath)
}

// SendFileByUrl is designed to send a file downloaded via a link
func (c SendingCategory) SendFileByUrl(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "SendFileByUrl", parameters, "")
}

// SendLocation is designed to send a geolocation message
func (c SendingCategory) SendLocation(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "sendLocation", parameters, "")
}

// SendContact is for sending a message with a contact
func (c SendingCategory) SendContact(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "sendContact", parameters, "")
}

// SendLink is designed to send a message with a link that will add an image preview, title and description
func (c SendingCategory) SendLink(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "sendLink", parameters, "")
}

// ForwardMessages is designed for forwarding messages to a personal or group chat
func (c SendingCategory) ForwardMessages(chatId, chatIdFrom string, messages []string) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "forwardMessages", map[string]interface{}{
		"chatId":     chatId,
		"chatIdFrom": chatIdFrom,
		"messages":   messages,
	}, "")
}
