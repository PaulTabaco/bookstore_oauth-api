package app

import (
	"fmt"

	"github.com/PaulTabaco/bookstore_oauth-api/src/clients/cassandra"
	"github.com/PaulTabaco/bookstore_oauth-api/src/domain/access_token"
	"github.com/PaulTabaco/bookstore_oauth-api/src/http"
	"github.com/PaulTabaco/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	checkCassandraAtStart()

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8181")
}

func checkCassandraAtStart() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	fmt.Println("*** Cassandra availability checked at start")
}
