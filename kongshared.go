
// Package kongshared provides shared functions for kong services
package kongshared

// Messages map
var sysMsgMap = map[string]string {
	"MSG0100": "User does not exists",
	"MSG0101": "User already exists",
	"MSG0200": "Invalid account id",
	"MSG0201": "Invalid account id",
}


// GetMsg returns the system message for the given code
func GetSysMsg(code string) string {
	return sysMsgMap[code]
}
