package repository_test

import (
	"prototype2/domain"
	"prototype2/repository"
	post "prototype2/repository/post"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (postRepo domain.PostRepository, mock sqlmock.Sqlmock, teardown func()) {
	db, mock, _ := repository.NewDBMock(t)

	postRepo = post.NewPostRepository(db)

	return postRepo, mock, func() {
		db.Close()
	}
}

func TestPostRepository_FindAll(t *testing.T) {
	postRepo, mock, teardown := setup(t)
	defer teardown()

	scenarios := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, post []domain.Post, err error)
	}{
		"When the SQL is normal": {
			arrange: func(t *testing.T) {
				post := domain.Post{ID: 1, Title: "Test Title", Text: "Test Text"}

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "text"}).
						AddRow(post.ID, post.Title, post.Text))
			},
			assert: func(t *testing.T, post []domain.Post, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 1, len(post))
			},
		},
		"In case of error SQL": {
			arrange: func(t *testing.T) {
				post := domain.Post{ID: 1, Title: "Test Title", Text: "Test Text"}

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "text"}).
						AddRow(post.ID, post.Title, post.Text))
			},
			assert: func(t *testing.T, post []domain.Post, err error) {
				// assert.NotEmpty(t, err)
			},
		},
	}

	for k, s := range scenarios {
		t.Run(k, func(t *testing.T) {
			s.arrange(t)

			post, err := postRepo.FindAll()

			s.assert(t, post, err)
		})
	}
}

// func TestPostRepository_Save(t *testing.T) {
// 	postRepo, mock, teardown := setup(t)
// 	defer teardown()

// 	post := &domain.Post{
// 		ID:    123,
// 		Title: "Test Title",
// 		Text:  "Test Text",
// 	}

// 	mock.ExpectBegin()
// 	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO posts SET id=\\? , title=\\? , text=\\?`)).
// 		WithArgs(post.ID, post.Title, post.Text).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "text"}).
// 			AddRow(post.ID, post.Title, post.Text))
// 	mock.ExpectCommit()

// 	result, err := postRepo.Save(post)

// 	assert.NoError(t, err)
// 	assert.Equal(t, int64(123), result.ID)
// }
