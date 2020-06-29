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
}

// PostRepository : represent the post's repository contract
type PostRepository interface {
	Save(post *Post) (*Post, error)
	FindAll() ([]Post, error)
	Delete(post *Post) error
	Migrate() error
}
