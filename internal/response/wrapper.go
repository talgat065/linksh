package response

type HasWrapper interface {
	Wrap() Wrapper
}

type Wrapper struct {
	ErrorCode int         `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
