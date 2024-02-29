package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"employee/controllers"
)

func main() {
	// Initialize database connection
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/employee")
	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	// Initialize Gin router
	router := gin.Default()

	router.GET("/api/employeeDetailsFromDB", func(c *gin.Context) {
		controllers.GetEmployeeDetailsFromDB(c, db)
	})
	// Add employee to db
	router.POST("/api/addEmployeeToDB", func(c *gin.Context) {
		controllers.AddEmployeeToDB(c, db)
	})
	// Delete employee from db
	router.DELETE("/api/employee/:id", func(c *gin.Context) {
		controllers.DeleteEmployeeFromDB(c, db)
	})
	// Get project details
	router.GET("/api/projectDetailsFromDB", func(c *gin.Context) {
		controllers.GetProjectsFromDB(c, db)
	})
	//Add new project to db
	router.POST("/api/addProjectToDB", func(c *gin.Context) {
		controllers.AddProjectToDB(c, db)
	})
	//Update project_id for employee
	router.PUT("/api/updateEmployeeProject", func(c *gin.Context) {
		controllers.UpdateEmployeeProjectDB(c, db)
	})
	
	// Run the server
	router.Run("localhost:8080")
}
