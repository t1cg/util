package user

type UserMessage struct { 
	Header	string
	Message	string
	Id	string
	Name	string
}

var Messages = struct { 
	CONNECTION_REFUSED	UserMessage
	INCORRECT_PASSWORD	UserMessage
	INVALID_ROLE	UserMessage
}{
	CONNECTION_REFUSED:	UserMessage{"Our servers are down.", "Please check back later.", "0", "CONNECTION_REFUSED", },
	INCORRECT_PASSWORD:	UserMessage{"Wrong Password.", "Try again or click Forgot password to reset it.", "1", "INCORRECT_PASSWORD", },
	INVALID_ROLE:	UserMessage{"Unauthorized", "Invalid user role.", "2", "INVALID_ROLE", },
}