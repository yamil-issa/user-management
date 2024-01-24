package main

import (
	"database/sql"
	"example/web-service-gin/pkg/router"
	"example/web-service-gin/pkg/user"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const dataSourceName = "root:@tcp(localhost:3306)/user_management"

func main() {
	// database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Connected to database!")

	// user service
	userService := user.NewUserService(db)

	r := router.SetupRouter(userService)

	// Start the HTTP server
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
