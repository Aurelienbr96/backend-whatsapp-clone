package testing

import (
	"example.com/boiletplate/ent"
	"example.com/boiletplate/ent/enttest"
	"example.com/boiletplate/ent/migrate"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/infrastructure/queue"
	"example.com/boiletplate/internal/server"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
)

func NewTestServer(t FullGinkgoTInterface) *gin.Engine {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	// Mock dependencies
	// Initialize server with mocked dependencies
	srv := server.NewServer(client, &otphandler.MockOTPHandler{}, &queue.MockPublisher{})
	server.NewHandlers(srv, &queue.MockPublisher{}, &otphandler.MockOTPHandler{}) // Initialize routes
	testServer := srv.Gin
	return testServer
}
