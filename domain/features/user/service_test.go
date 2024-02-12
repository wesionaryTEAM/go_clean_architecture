package user_test

import (
	"clean-architecture/domain/domainif"
	"clean-architecture/domain/models"
	"clean-architecture/mocks"
	mockdomainif "clean-architecture/mocks/clean-architecture/domain/domainif"
	"clean-architecture/pkg/types"
	"clean-architecture/pkg/utils"
	"fmt"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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

	It("should return user if user exists", func() {
		user := models.User{
			Name:  "John Doe",
			Email: "john_doe@example.com",
			Age:   25,
		}
		binaryuuid := types.BinaryUUID(mockUID)
		user.ID = binaryuuid
		fmt.Println("binary uuid ", binaryuuid)
		mockUserService.EXPECT().GetOneUser(binaryuuid).Return(user, nil)
		fmt.Println("user ", user)

		setupDI()

		Expect(user).ToNot(BeNil())
	})

})
