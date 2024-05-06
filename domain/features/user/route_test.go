package user_test

import (
	"clean-architecture/domain/domainif"
	"clean-architecture/domain/models"
	"clean-architecture/mocks"
	mockdomainif "clean-architecture/mocks/clean-architecture/domain/domainif"
	mockinterfaces "clean-architecture/mocks/clean-architecture/pkg/interfaces"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/interfaces"
	"clean-architecture/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/mock"
	"go.uber.org/fx"
)

var _ = Describe("User Route Tests", func() {
	var (
		t                    GinkgoTInterface
		router               infrastructure.Router
		authMiddleware       *mockinterfaces.MockAuthMiddleware
		mockUserService      *mockdomainif.MockUserService
		paginationMiddleware *mockinterfaces.MockPaginationMiddleware
	)

	setupDI := func() {
		err := mocks.DI(t,
			fx.Populate(&router),
			utils.FxReplaceAs(mockUserService, new(domainif.UserService)),
			utils.FxReplaceAs(authMiddleware, new(interfaces.AuthMiddleware)),
			utils.FxReplaceAs(paginationMiddleware, new(interfaces.PaginationMiddleware)),
		)
		if err != nil {
			t.Error(err)
		}
	}

	BeforeEach(func() {
		t = GinkgoT()
		authMiddleware = mockinterfaces.NewMockAuthMiddleware(t)
		authMiddleware.EXPECT().HandleAuthWithRole(mock.Anything).Return(mocks.MockAuthSuccessHandler)
		paginationMiddleware = mockinterfaces.NewMockPaginationMiddleware(t)
		paginationMiddleware.EXPECT().Handle().Return(mocks.MockPaginationHandler)
		mockUserService = mockdomainif.NewMockUserService(t)

	})

	It("should return users in data and pagination struct with total users count", func() {
		users := []models.User{
			{Name: "John Doe", Email: "doe@test.com", Age: 30},
		}
		count := int64(1)
		mockUserService.EXPECT().GetAllUser().Return(&users, count, nil)

		setupDI()

		userStr, _ := json.Marshal(users)
		output := fmt.Sprintf(`{"data":%s,"pagination":{"count":%v,"has_next":false}}`, userStr, count)
		apitest.
			New().
			Handler(router).
			Get("/api/user").
			Expect(t).
			Status(http.StatusOK).
			Body(output).
			End()
	})

	It("should return error, if service returns error", func() {
		mockUserService.EXPECT().GetAllUser().Return(nil, 0, errors.New("error"))

		setupDI()

		apitest.
			New().
			Handler(router).
			Get("/api/user").
			Expect(t).
			Status(http.StatusInternalServerError).
			End()
	})

})
