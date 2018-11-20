package application

type ApplicationMessage struct { 
	Message	string
	Code	string
	Id	string
	Name	string
}

var Messages = struct { 
	SUCCESS	ApplicationMessage
}{
	SUCCESS:	ApplicationMessage{"Ok", "200", "0", "SUCCESS", },
}