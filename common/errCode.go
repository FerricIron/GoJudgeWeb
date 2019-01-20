package common

const (
	Success                   	= 0

	DataBaseUnavaliable		  	= 2
	JSONComponentUnavaliable	= 3

	TokenComponentUnavaliable 	= 300
	UserExist                 	= 10100
	AddUserError              	= 10101
	UserNotExist              	= 10102

	TokenNotExist 				= 10300
	PermissionDenied			= 10301


	InvalidParams 				= 10400
	InvalidForm   				= 10401
)

const (
	JudgeFinished = iota
	AcceptCode
	WrongAnwser
	ComplierError
	TimeLimitError
	ComlierTimeLimitError
	MemoryLimitError
	OuputLimitError
	RunTimeError
	OtherError = -1
)
