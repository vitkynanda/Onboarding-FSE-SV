package helpers

import "go-api/models/dto"

func ResponseError(status string, err interface{}) dto.Response { 
	return dto.Response{
		Status: status,
		Error: err,
		Data: nil,
	}
}

func ResponseSuccess(status string, err interface{}, data interface{}) dto.Response { 
	return dto.Response{
		Status: status,
		Error: err,
		Data: data,
	}
}