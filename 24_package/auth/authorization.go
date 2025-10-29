package auth

// for example if we have some sort of implementation that we dont want to expose
// than we name it starting with non capital first latter

func extractSession() string {
	return "loggedins"
}

func GetSession() string {
	return extractSession()
}
