package message

func Message(status bool, code, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "code": code, "message": message}
}
