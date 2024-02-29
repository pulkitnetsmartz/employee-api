package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"employee/controllers"
)

func SetupEmployeeRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/api/employeeDetailsFromDB", func(c *gin.Context) {
		controllers.GetEmployeeDetailsFromDB(c, db)
	})

	router.POST("/api/addEmployeeToDB", func(c *gin.Context) {
		controllers.AddEmployeeToDB(c, db)
	})

	router.PUT("/api/updateEmployeeProject", func(c *gin.Context) {
		controllers.UpdateEmployeeProjectDB(c, db)
	})

	router.DELETE("/api/employee/:id", func(c *gin.Context) {
		controllers.DeleteEmployeeFromDB(c, db)
	})

	router.GET("/api/projectDetailsFromDB", func(c *gin.Context) {
		controllers.GetProjectsFromDB(c, db)
	})

	router.POST("/api/addProjectToDB", func(c *gin.Context) {
		controllers.AddProjectToDB(c, db)
	})
}
