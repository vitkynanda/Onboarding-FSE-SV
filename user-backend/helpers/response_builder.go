package helpers

import "go-api/models/dto"

func ResponseError(msg string, status int) dto.Response { 
	return dto.Response{
		Message: msg,
		StatusCode: status,
		Data: nil,
	}
}

func ResponseSuccess(msg string, status int, data interface{}) dto.Response { 
	return dto.Response{
		Message: msg,
		StatusCode: status,
		Data: data,
	}
}