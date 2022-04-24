package models

type Response struct {
	code    int
	message string
	data    interface{}
}

func getResponse(code int, message string, data interface{}) Response {
	return Response{
		code,
		message,
		data,
	}
}

func ResponseSuccess(data interface{}) Response {
	return getResponse(200, "success", data)
}
