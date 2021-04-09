package smsglobal

// Create is same Factory function
func Create(Username string, Password string, SenderName string) *SMSGLOBAL {
	return Factory(Username, Password, SenderName)
}
