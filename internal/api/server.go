package api

import (
	"fmt"
	db "simplebank/internal/repository"
	"simplebank/internal/token"
	"simplebank/util"

	"aidanwoods.dev/go-paseto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config util.Config
	store  db.Store
	token  token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	key, err := paseto.V4SymmetricKeyFromHex(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot assign key from hex %v", err)
	}

	tokenMaker, err := token.NewPasetoMaker(key, config.TokenImplicit)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %v", err)
	}

	server := &Server{
		config: config,
		store:  store,
		token:  tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)

	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccounts)
	router.POST("/transfers", server.createTransfer)
	router.POST("/users", server.createUser)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
