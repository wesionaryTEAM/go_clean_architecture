package user_test

import (
	"clean-architecture/domain/domainif"
	"clean-architecture/domain/models"
	"clean-architecture/mocks"
	mockdomainif "clean-architecture/mocks/clean-architecture/domain/domainif"
	"clean-architecture/pkg/types"
	"clean-architecture/pkg/utils"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var _ = Describe("User Service Tests", func() {
	var (
		t               GinkgoTInterface
		mockUserService *mockdomainif.MockUserService
		mockUID         uuid.UUID
	)

	setupDI := func() {
		err := mocks.DI(t,
			utils.FxReplaceAs(mockUserService, new(domainif.UserService)),
		)
		if err != nil {
			t.Error(err)
		}
	}

	BeforeEach(func() {
		t = GinkgoT()
		mockUserService = mockdomainif.NewMockUserService(t)
		mockUID, _ = uuid.NewRandom()
	})

	Describe("Getting a user", func() {
		It("should return user if user exists", func() {
			user := models.User{}
			mockUserService.EXPECT().GetOneUser(types.BinaryUUID(mockUID)).Return(user, nil)

			setupDI()

			res, err := mockUserService.GetOneUser(types.BinaryUUID(mockUID))

			Expect(res).To(Equal(user))
			Expect(err).To(BeNil())
		})

		It("should return error if user does not exist", func() {
			user := models.User{}
			mockUserService.EXPECT().GetOneUser(mock.Anything).Return(user, gorm.ErrRecordNotFound)
			setupDI()

			res, err := mockUserService.GetOneUser(types.BinaryUUID(mockUID))

			Expect(err).To(Equal(gorm.ErrRecordNotFound))
			Expect(res.ID).To(Equal(types.BinaryUUID(uuid.Nil)))

		})
	})

	Describe("Getting all users", func() {
		It("should return all users", func() {
			users := []models.User{}
			mockUserService.EXPECT().GetAllUser().Return(&users, int64(0), nil)

			setupDI()

			usersRes, count, err := mockUserService.GetAllUser()

			Expect(usersRes).To(Equal(&users))
			Expect(count).To(Equal(int64(0)))
			Expect(err).To(BeNil())
		})
	})

	Describe("Updating a user", func() {
		It("should update the user", func() {
			user := models.User{}
			mockUserService.EXPECT().UpdateUser(&user).Return(nil)

			setupDI()

			err := mockUserService.UpdateUser(&user)

			Expect(err).To(BeNil())
		})
	})

	Describe("Creating a user", func() {
		It("should create the user", func() {
			user := models.User{
				Name:       "Jhon Doe",
				Email:      "jhon.doe@test.com",
				Age:        25,
				ProfilePic: "https://www.profilepic.url.com",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}
			mockUserService.EXPECT().Create(&user).Return(nil)

			setupDI()

			err := mockUserService.Create(&user)

			Expect(err).To(BeNil())
		})
	})

	Describe("Deleting a user", func() {
		It("should delete the user", func() {
			mockUserService.EXPECT().DeleteUser(types.BinaryUUID(mockUID)).Return(nil)

			setupDI()

			err := mockUserService.DeleteUser(types.BinaryUUID(mockUID))

			Expect(err).To(BeNil())
		})
	})

})
