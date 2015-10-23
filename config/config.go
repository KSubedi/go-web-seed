package config

//Set the configuration values on this file

// Get gets you the value for the key
func Get(t string) string {
	switch t {
	case "WEBSERVER_PORT":
		return "8000"
	case "SESSION_HASH":
		return "98zxc79*Z&XC(879"
	case "SESSION_NAME":
		return "master-session"
	default:
		return ""
	}
}
