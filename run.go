package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)



func main() {
	db := initDatabase()
	defer db.Close() // close connection after server stopped 
	engine := gin.Default()
	initRoutings(engine)
	// testDbWorking(db)
	engine.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}
