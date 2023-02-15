package methods

type QueuesCategory struct {
	BasisAPI BasisAPIInterface
}

// ShowMessagesQueue is designed to get the list of messages that are in the queue to be sent
func (c QueuesCategory) ShowMessagesQueue() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "ShowMessagesQueue", nil, "")
}

// ClearMessagesQueue is designed to clear the queue of messages to be sent
func (c QueuesCategory) ClearMessagesQueue() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "ClearMessagesQueue", nil, "")
}
