package domain

// Post : Database model for post
type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// PostService : represent the post's services
type PostService interface {
	Validate(post *Post) error
	Create(post *Post) (*Post, error)
	FindAll() ([]Post, error)
	GetByID(id int64) (*Post, error)
	Delete(id int64) error
}

// PostRepository : represent the post's repository contract
type PostRepository interface {
	Save(post *Post) (*Post, error)
	FindAll() ([]Post, error)
	FindByID(id int64) (*Post, error)
	Delete(post *Post) error
	Migrate() error
}
