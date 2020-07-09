package domain

// User : Database model for user
type User struct {
	ID       string `json:"id" gorm:"unique;not null"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password"`
	DOB      string `json:"dob"`
}

// UserService : represent the user's services
type UserService interface {
	Validate(user *User) error
	ValidateAge(user *User) bool
	Create(user *User) (*User, error)
	FindAll() ([]User, error)
}

// UserRepository : represent the user's repository contract
type UserRepository interface {
	Save(user *User) (*User, error)
	FindAll() ([]User, error)
	Delete(user *User) error
	Migrate() error
}
