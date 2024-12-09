package http_test

import (
	"encoding/json"
	"example.com/boiletplate/internal/user/repository"
)

import (
	"example.com/boiletplate/ent"
	mockqueue "example.com/boiletplate/infrastructure/queue/mock"
	"example.com/boiletplate/testing"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = FDescribe("Most outer container", Ordered, ContinueOnFailure, func() {
	var (
		testServer     *testing.TestServer
		client         *ent.Client
		mockPublisher  *mockqueue.MockIPublisher
		user           *testing.UserFixture
		userRepository *repository.Repository
	)
	BeforeAll(func() {
		testServer = testing.NewTestServer(GinkgoT())
		client = testServer.Client
		mockPublisher = testServer.MockQueue

		user = testing.GenerateUser("+33602222639")
		userRepository = repository.NewUserRepository(client)

		testServer.CreateManyUsers([]*testing.UserFixture{user})
	})
	AfterAll(func() {
		testServer.Close()
	})
	Describe("POST CREATE A USER", func() {

		Context("When a user registers with correct informations", func() {
			It("should create a user", func() {

				mockPublisher.EXPECT().PushMessage(gomock.Eq([]byte(`{"type":"CreatedUserSuccess","payload":{"phoneNumber":"+33602222632"}}`))).Return(nil)
				// Arrange: Prepare invalid JSON request
				phoneNumber := "+33602222632"
				reqBody := fmt.Sprintf(`{"phoneNumber": "%s"}`, phoneNumber) // Empty phone number
				req := httptest.NewRequest(http.MethodPost, "/api/v1/user/", strings.NewReader(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()

				fmt.Printf("test server: %v", testServer)

				// Act: Send request to the test server
				testServer.Gin.ServeHTTP(resp, req)

				user, _ := userRepository.GetOneByPhoneNumber(phoneNumber)

				// Assert: Verify response
				Expect(user.PhoneNumber).To(Equal(phoneNumber))

				Expect(resp.Code).To(Equal(http.StatusCreated))
			})
		})
	})

	Describe("GET A USER BY PHONE NUMBER", func() {
		It("Should return a user", func() {
			getUrl := fmt.Sprintf("/api/v1/user/by-phone/%s", user.PhoneNumber)
			req := httptest.NewRequest(http.MethodGet, getUrl, nil)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			testServer.Gin.ServeHTTP(resp, req)
			jsonUser, _ := json.Marshal(user)

			Expect(resp.Body.String()).To(Equal(jsonUser))
		})
	})
})
