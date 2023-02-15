package methods

type BasisAPIInterface interface {
	Request(method, APIMethod string, data map[string]interface{}, filePath string) (map[string]interface{}, error)
}
