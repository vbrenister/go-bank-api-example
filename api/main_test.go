package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/vbrenister/go-bank-api-example/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	return NewServer(store)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}