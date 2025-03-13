package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"project/controllers"
	"project/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ensure upload directory exists
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	r := gin.Default()

	// Apply authentication middleware
	r.Use(middlewares.AuthMiddleware())

	// Routes
	r.POST("/upload", controllers.UploadExcel)              // Upload an Excel file
	r.GET("/process-template", controllers.ProcessTemplate) // Fill template dynamically

	// Run server
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
