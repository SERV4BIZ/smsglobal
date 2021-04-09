package smsglobal

// Factory is create new object
func Factory(Username string, Password string, SenderName string) *SMSGLOBAL {
	obj := new(SMSGLOBAL)
	obj.Username = Username
	obj.Password = Password
	obj.SenderName = SenderName
	return obj
}
