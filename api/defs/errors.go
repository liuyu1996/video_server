package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:400, Error:Err{Error:"RequestBody is not correct", ErrorCode:"001"}}
	ErrorNotAuthUser = ErrorResponse{HttpSC:401, Error:Err{Error:"User authentication failed", ErrorCode:"002"}}
	ErrorDBError = ErrorResponse{HttpSC:500, Error:Err{Error:"DB ops failed", ErrorCode:"003"}}
	ErrorInternalError = ErrorResponse{HttpSC:500, Error:Err{Error:"Internal service error", ErrorCode:"004"}}
	)