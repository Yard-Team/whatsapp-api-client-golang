package api

import (
	"strconv"
	"strings"

	"github.com/Yard-Team/whatsapp-api-client-golang/pkg/categories"
)

type BasisAPI struct {
	URL              string
	IDInstance       string
	APITokenInstance string
}

func (a BasisAPI) Methods() categories.BasisAPICategories {
	return categories.BasisAPICategories{BasisAPI: a}
}

func (a BasisAPI) Request(method, APIMethod string, data map[string]interface{}, filePath string) (map[string]interface{}, error) {
	url := a.getURL(method, APIMethod, data)

	return ExecuteRequest(method, url, data, filePath)
}

func (a BasisAPI) getURL(method, APIMethod string, data map[string]interface{}) string {
	if a.URL != "" {
		return a.URL
	}

	var url strings.Builder

	url.WriteString("https://api.basis-api.com/")
	url.WriteString("waInstance")
	url.WriteString(a.IDInstance)
	url.WriteString("/")
	url.WriteString(APIMethod)
	url.WriteString("/")
	url.WriteString(a.APITokenInstance)

	if method == "DELETE" {
		url.WriteString("/")
		url.WriteString(strconv.Itoa(data["receiptId"].(int)))
	}

	return url.String()
}
