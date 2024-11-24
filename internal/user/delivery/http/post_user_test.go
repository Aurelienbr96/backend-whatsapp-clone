package http_test

import (
	"example.com/boiletplate/testing"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = FDescribe("Most outer container", Ordered, ContinueOnFailure, func() {
	var (
		testServer *gin.Engine
	)
	BeforeAll(func() {
		testServer = testing.NewTestServer(GinkgoT())
	})
	Describe("POST CREATE A USER", func() {

		Context("When a user registers with incorrect information", func() {
			It("should return a 400 error", func() {
				// Arrange: Prepare invalid JSON request
				reqBody := `{"phoneNumber": ""}` // Empty phone number
				req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(reqBody))
				req.Header.Set("Content-Type", "application/json")
				resp := httptest.NewRecorder()

				// Act: Send request to the test server
				testServer.ServeHTTP(resp, req)

				// Assert: Verify response
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				Expect(resp.Body.String()).To(ContainSubstring("Request body cannot be empty"))
			})
		})
	})
})
