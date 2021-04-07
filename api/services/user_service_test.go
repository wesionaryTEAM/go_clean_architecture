package services

import (
	"testing"

	"prototype2/models"

	"github.com/stretchr/testify/mock"
)

/**
* Mocking the repository layers
 */
type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(user *models.User) (*models.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*models.User), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]models.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]models.User), args.Error(1)
}

func (mock *MockRepository) Delete(user *models.User) error {
	args := mock.Called()
	return args.Error(1)
}

func (mock *MockRepository) Migrate() error {
	args := mock.Called()
	return args.Error(1)
}

// Write your UNIT TEST code below


/**
* Function name: Validate
* Test case: When user is nil
 */
func TestValidateEmptyUser(t *testing.T) {
	//testService := NewUserService(nil)
	//
	//err := testService.Validate(nil)
	//
	//assert.NotNil(t, err)
	//
	//assert.Equal(t, "The user is empty", err.Error())
}


/**
* Function name: Validate
* Test case: When name and email and DOB is empty
 */
func TestValidateEmptyNameField(t *testing.T) {
	//testService := NewUserService(nil)
	//scenarios := []struct {
	//	user   models.User
	//	expect string
	//}{
	//	{
	//		user:   models.User{ID: uuid.New(), Name: "", Email: "testemail@test.com", DOB: "1995-12-28"},
	//		expect: "The name of user is empty",
	//	},
	//	{
	//		user:   models.User{ID: "1", Name: "Lorem Ipsum", Email: "", DOB: "1995-12-28"},
	//		expect: "The email of user is empty",
	//	},
	//	{
	//		user:   models.User{ID: "1", Name: "Lorem Ipsum", Email: "testemail@test.com", DOB: ""},
	//		expect: "The DOB of user is empty",
	//	},
	//}
	//
	//for _, s := range scenarios {
	//	err := testService.Validate(&s.user)
	//	assert.NotNil(t, err)
	//	assert.Equal(t, s.expect, err.Error())
	//}
}


/**
* Function name: Find all
* Test case: Should return all the mocked objects
 */
func TestFindAll(t *testing.T) {
	//mockRepository := new(MockRepository)
	//
	//var identifier string = "1"
	//
	//user := models.User{ID: identifier, Name: "Binod Kafle", Email: "mebinod50@gmail.com", DOB: "1993-12-13"}
	//
	//// setting up the find all action as a mocked action
	//mockRepository.On("FindAll").Return([]models.User{user}, nil)
	//
	//// instantiating the UserService with mocked repository
	//testService := NewUserService(mockRepository)
	//
	//result, _ := testService.FindAll()
	//
	//mockRepository.AssertExpectations(t)
	//
	//assert.Equal(t, identifier, result[0].ID)
	//assert.Equal(t, "Binod Kafle", result[0].Name)
	//assert.Equal(t, "mebinod50@gmail.com", result[0].Email)
	//assert.Equal(t, "1993-12-13", result[0].DOB)
}


/**
*	Function name: ValidateAge
* Test case: Should return false when age is less than 18 and true when age is
* above 18
 */
func TestValidateAge(t *testing.T) {
	//testService := NewUserService(nil)
	//
	//// Initializing the test scenarios
	//scenarios := []struct {
	//	user   models.User
	//	expect bool
	//}{
	//	{user: models.User{ID: "1", Name: "Abcde", Email: "testemail@test.com", DOB: "1965-12-28"}, expect: true},
	//	{user: models.User{ID: "1", Name: "Lorem Ipsum", Email: "abc@abc.com", DOB: "2020-12-28"}, expect: false},
	//}
	//
	//// Testing the scenarios
	//for _, s := range scenarios {
	//	result := testService.ValidateAge(&s.user)
	//	assert.Equal(t, s.expect, result)
	//}
}


/**
*	Function name: Create
* Test case: Should assert according to the created user
 */
func TestCreate(t *testing.T) {
	//mockRepository := new(MockRepository)
	//user := models.User{Name: "Binod Kafle", Email: "mebinod50@gmail.com", DOB: "1993-12-13"}
	//
	//mockRepository.On("Save").Return(&user, nil)
	//
	//testService := NewUserService(mockRepository)
	//
	//result, err := testService.Create(&user)
	//
	//mockRepository.AssertExpectations(t)
	//
	//assert.NotNil(t, result.ID)
	//assert.Equal(t, "Binod Kafle", result.Name)
	//assert.Equal(t, "mebinod50@gmail.com", result.Email)
	//assert.Equal(t, "1993-12-13", result.DOB)
	//assert.Nil(t, err)
}
