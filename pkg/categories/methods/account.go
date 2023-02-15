package methods

type AccountCategory struct {
	BasisAPI BasisAPIInterface
}

// GetSettings is designed to get the current settings of the account
func (c AccountCategory) GetSettings() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "GetSettings", nil, "")
}

// SetSettings is for setting the account settings
func (c AccountCategory) SetSettings(parameters map[string]interface{}) (map[string]interface{}, error) {
	method := "GET"
	if parameters != nil {
		method = "POST"
	}

	return c.BasisAPI.Request(method, "SetSettings", parameters, "")
}

// SetSystemProxy is for setting up a system proxy
func (c AccountCategory) SetSystemProxy() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "SetSystemProxy", nil, "")
}

// GetStateInstance is designed to get the state of the account
func (c AccountCategory) GetStateInstance() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "getStateInstance", nil, "")
}

// GetStatusInstance is designed to get the socket connection state of the account instance with WhatsApp
func (c AccountCategory) GetStatusInstance() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "getStatusInstance", nil, "")
}

// Reboot is designed to restart the account
func (c AccountCategory) Reboot() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "Reboot", nil, "")
}

// Logout is designed to unlogin the account
func (c AccountCategory) Logout() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "Logout", nil, "")
}

// QR is designed to get a QR code
func (c AccountCategory) QR() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "qr", nil, "")
}

// SetProfilePicture is designed to set the avatar of the account
func (c AccountCategory) SetProfilePicture(filePath string) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "setProfilePicture", nil, filePath)
}
