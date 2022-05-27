package db

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/thanupakonc/pingkaii_meta/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))

	router.Use(cors.Default())

	router.GET("/province", server.listProvinces)
	router.GET("/district/:id", server.listProvinceDistricts)
	router.GET("/tambon/:id", server.listDistrictTambon)
	router.GET("/postcode/:id", server.listTambonPostcode)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
