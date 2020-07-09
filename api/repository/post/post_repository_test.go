package repository_test

import (
	"prototype2/api/repository"
	post "prototype2/api/repository/post"
	"prototype2/domain"
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

// TestPostRepository_FindAll : test case for FindAll method
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

// TestPostRepository_Save : test case for Save method
func TestPostRepository_Save(t *testing.T) {
	postRepo, mock, teardown := setup(t)
	defer teardown()

	post := &domain.Post{
		ID:    123,
		Title: "Test Title",
		Text:  "Test Text",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT *").
		WithArgs(post.ID, post.Title, post.Text).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := postRepo.Save(post)

	assert.Nil(t, err)
	assert.Equal(t, post, result)
}

// TestPostRepository_Delete : test case for Delete method
func TestPostRepository_Delete(t *testing.T) {
	postRepo, mock, teardown := setup(t)
	defer teardown()

	post := &domain.Post{
		ID:    123,
		Title: "Test Title",
		Text:  "Test Text",
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE *").
		WithArgs(post.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := postRepo.Delete(post)

	assert.Nil(t, err)
}
