
// Package kongshared provides shared functions for kong services
package kongshared

// Messages map
var msgMap = map[string]string {
	"MSG0100": "User does not exists",
	"MSG0101": "User already exists",
	"MSG0200": "Invalid account id",
	"MSG0201": "Invalid account id",
}


// GetMsg returns the message for the given code
func GetMsg(code string) string {
	return msgMap[code]
}
