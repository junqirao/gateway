package response

import "fmt"

type JSON struct {
	Status  int         `json:"-"`              // http status code
	Code    int         `json:"code"`           // business code
	Message string      `json:"message"`        // response message
	Data    interface{} `json:"data,omitempty"` // response data
}

// Error implement of error
func (j JSON) Error() string {
	return j.Message
}

// NewJSONResp new json resp
func NewJSONResp(status, code int, message string) JSON {
	return JSON{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
func (j JSON) WithCode(code int) JSON {
	j.Code = code
	return j
}

func (j JSON) WithMessage(message string) JSON {
	j.Message = message
	return j
}

func (j JSON) WithDetail(detail interface{}) JSON {
	j.Message = fmt.Sprintf("%s: %s", j.Message, detail)
	return j
}

func (j JSON) WithData(data interface{}) JSON {
	j.Data = data
	return j
}

func (j JSON) WithStatusCode(code int) JSON {
	j.Status = code
	return j
}
