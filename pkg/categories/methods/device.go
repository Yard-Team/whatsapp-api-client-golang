package methods

type DeviceCategory struct {
	BasisAPI BasisAPIInterface
}

// GetDeviceInfo is designed to get information about the device (phone) on which the WhatsApp Business application is running
func (c DeviceCategory) GetDeviceInfo() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "GetDeviceInfo", nil, "")
}
