package proto

type AgentAuth struct {
	ApiKey   string
	Uuid     string
	Hostname string
	Username string
}

type UserAuth struct {
	Email    string
	Provider string
	ApiKey   string
}

type AuthResponse struct {
	Code  uint   // standard HTTP status (http://httpstatus.es/)
	Error string // empty if auth ok (Code=200)
}