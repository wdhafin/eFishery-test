package entity

//User is entity user to store user data
type User struct {
	ID        uint64 `json:"id"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	Timestamp int64  `json:"timestamp"`
}
