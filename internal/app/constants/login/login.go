package login

const (
	SuccessMessage         = "Welcome, %s!"
	FailMessage            = "Login Failed"
	AccountNotFoundMessage = "Account not found. Persist an Account to access the system"
	InvalidPasswordMessage = "Invalid Password"
	LogoutSuccessMessage   = "Logout Successfully."
)

// Message will be used as custom messages for login routes
type Message struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	Token   string `json:"token,omitempty"`
}
