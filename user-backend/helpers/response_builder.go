package helpers

import "go-api/models/dto"

func ResponseError(status string, err interface{}, code int) dto.Response { 
	return dto.Response{
		StatusCode: code,
		Status: status,
		Error: err,
		Data: nil,
	}
}

func ResponseSuccess(status string, err interface{}, data interface{}, code int) dto.Response { 
	return dto.Response{
		StatusCode: code,
		Status: status,
		Error: err,
		Data: data,
	}
}