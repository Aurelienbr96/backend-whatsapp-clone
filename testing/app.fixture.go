package testing

import (
	"context"
	"example.com/boiletplate/ent"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	mock_queue "example.com/boiletplate/infrastructure/queue/mock"
	mockpackage "example.com/boiletplate/infrastructure/upload-blob/mock"
	"example.com/boiletplate/internal/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
)

type TestServer struct {
	Gin       *gin.Engine
	Client    *ent.Client
	MockQueue *mock_queue.MockIPublisher
}

func NewTestServer(t FullGinkgoTInterface) *TestServer {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		fmt.Println("Error opening connection to sqlite", err)
		panic(err)
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}
	// Mock dependencies
	// Initialize mocks
	otpHandler := &otphandler.MockOTPHandler{}
	mockPublisher := mock_queue.NewMockIPublisher(ctrl)

	// Configure mock expectations if any
	mockBlock := mockpackage.NewMockBlobAdapter()

	// Initialize server with mocked dependencies
	srv := server.NewServer(client, otpHandler, mockPublisher)
	server.NewHandlers(srv, mockPublisher, otpHandler, mockBlock) // Initialize routes

	return &TestServer{
		Gin:       srv.Gin,
		Client:    client,
		MockQueue: mockPublisher,
	}
}

func (t *TestServer) CreateManyUsers(users []*UserFixture) {
	var usersToCreate []*ent.UserCreate
	for _, u := range users {
		usersToCreate = append(usersToCreate, t.Client.User.Create().SetPhoneNumber(u.PhoneNumber))
	}
	t.Client.User.CreateBulk(usersToCreate...).Save(context.Background())
}
