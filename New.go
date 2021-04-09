package smsglobal

// New is same Factory function
func New(Username string, Password string, SenderName string) *SMSGLOBAL {
	return Factory(Username, Password, SenderName)
}
