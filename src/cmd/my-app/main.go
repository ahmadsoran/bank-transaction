package main

import (
	"github.com/gin-gonic/gin"
	"main.go/src/cmd/my-app/migration"
	"main.go/src/cmd/my-app/routes"
	"main.go/src/internal/constant"
	conf "main.go/src/internal/db"
)

func init() {
	constant.InitENVS()
	conf.ConnectDB()
}
func main() {
	r := gin.Default()
	migration.Migration()
	routes.Routes(r)
	r.Run("localhost:6969")
}
