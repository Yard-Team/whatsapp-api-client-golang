package methods

type ReceivingCategory struct {
	BasisAPI BasisAPIInterface
}

// ReceiveNotification is designed to receive a single incoming notification from the notification queue
func (c ReceivingCategory) ReceiveNotification() (map[string]interface{}, error) {
	return c.BasisAPI.Request("GET", "ReceiveNotification", nil, "")
}

// DeleteNotification is designed to remove an incoming notification from the notification queue
func (c ReceivingCategory) DeleteNotification(receiptId int) (map[string]interface{}, error) {
	return c.BasisAPI.Request("DELETE", "DeleteNotification", map[string]interface{}{
		"receiptId": receiptId,
	}, "")
}

// DownloadFile is for downloading received and sent files
func (c ReceivingCategory) DownloadFile(chatId, idMessage string) (map[string]interface{}, error) {
	return c.BasisAPI.Request("POST", "downloadFile", map[string]interface{}{
		"chatId":    chatId,
		"idMessage": idMessage,
	}, "")
}
