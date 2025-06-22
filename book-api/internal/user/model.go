package user

type User struct {
	Username string `json:"username"`
	Password string `json:"password"` // for demo, plain text (use hashed in production)
	Role     string `json:"role"`     // "admin" or "user"
}

// Mocked users (replace with DB later)
var users = []User{
	{Username: "admin", Password: "admin123", Role: "admin"},
	{Username: "user", Password: "user123", Role: "user"},
}

func Authenticate(username, password string) *User {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return &u
		}
	}
	return nil
}
