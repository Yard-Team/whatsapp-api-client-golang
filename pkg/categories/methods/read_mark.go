package methods

type ReadMarkCategory struct {
	BasisAPI BasisAPIInterface
}

// ReadChat is designed to mark chat messages as read
func (c ReadMarkCategory) ReadChat(parameters map[string]interface{}) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "ReadChat", parameters, "")
}
