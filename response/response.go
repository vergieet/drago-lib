package response

func Success(data string) string{
	return "{\"status\":1,\"data\":" + data + "}"
}

func Fail(message string) string{
	return "{\"status\":0,\"message\":\"" + message + "\"}"
}