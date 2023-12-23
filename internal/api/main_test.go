package api

import (
	"fmt"
	"os"
	db "simplebank/internal/repository"
	"simplebank/util"
	"testing"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymetricKey: "39781be0ea0773931d3ddbf7eaeb11d1e75c6c3938286328c2fbd24daaedc578",
		TokenImplicit:    "access",
		TokenDuration:    time.Minute}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server

}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}

func TestKey(t *testing.T) {
	key := paseto.NewV4SymmetricKey().ExportHex()

	fmt.Println("THIS IS THE KEY:", key)

	keybt, err := paseto.V4SymmetricKeyFromHex(key)
	if err != nil {
		fmt.Println("THIS IS THE ERROR:", err)
	}
	fmt.Println("THIS IS THE Real", keybt)

}
