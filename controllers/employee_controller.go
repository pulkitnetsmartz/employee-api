package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"employee/models"
)


func GetEmployeeDetailsFromDB(c *gin.Context, db *sql.DB) {
    // Query the database to fetch employee details
    rows, err := db.Query("SELECT employee_id, employee_name, role_id, project_id FROM employees")
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch employee details"})
        return
    }
    defer rows.Close()

    // Initialize a slice to store employee details
    var employees []models.Employee_detail

    // Iterate through the rows returned by the query
    for rows.Next() {
        var employee models.Employee_detail

        // Scan the values from the row into variables
        if err := rows.Scan(&employee.EmployeeId, &employee.Name, &employee.Role, &employee.ProjectID); err != nil {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse employee details"})
            return
        }

        // Append the employee to the slice
        employees = append(employees, employee)
    }

    // Check for errors during row iteration
    if err := rows.Err(); err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to iterate over employee details"})
        return
    }

    // Return the fetched employee details
    c.IndentedJSON(http.StatusOK, employees)
}

func AddEmployeeToDB(c *gin.Context, db *sql.DB) {
	var newEmployee models.Employee_detail

	// Bind the new employee details from JSON request body
	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	// If the database schema doesn't have a `project_name` column, remove it from the INSERT statement.
	// Update the query based on your actual schema.
	result, err := db.Exec("INSERT INTO employees (employee_id, employee_name, role_id) VALUES (?, ?, ?)",
		newEmployee.EmployeeId, newEmployee.Name, newEmployee.Role)
	if err != nil {
		fmt.Println("Error inserting employee:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to add employee"})
		return
	}

	// Get the ID of the newly created employee
	employeeID, err := result.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve employee ID"})
		return
	}

	// Set the employee ID in the response
	newEmployee.EmployeeId = int(employeeID)

	// Return the created employee details
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func UpdateEmployeeProjectDB(c *gin.Context, db *sql.DB) {
	var updatePayload struct {
		EmployeeId int `json:"employee_id"`
		ProjectID  int `json:"project_id"`
	}

	// Bind the update payload from JSON request body
	if err := c.BindJSON(&updatePayload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	// Update the project_id for the specified employee in the database
	_, err := db.Exec("UPDATE employees SET project_id = ? WHERE employee_id = ?",
		updatePayload.ProjectID, updatePayload.EmployeeId)
	if err != nil {
		fmt.Println("Error updating employee project:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update employee project"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Employee project updated successfully"})
}

func DeleteEmployeeFromDB(c *gin.Context, db *sql.DB) {
	// Get the employee ID from the request parameters
	employeeID := c.Param("id")

	// Check if the employee ID is empty
	if employeeID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Employee ID is required"})
		return
	}

	// Convert the employee ID to an integer
	id, err := strconv.Atoi(employeeID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid employee ID format"})
		return
	}

	// Execute the delete query
	result, err := db.Exec("DELETE FROM employees WHERE employee_id = ?", id)
	if err != nil {
		fmt.Println("Error deleting employee:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete employee"})
		return
	}

	// Check if any rows were affected
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

func GetProjectsFromDB(c *gin.Context, db *sql.DB) {
	// Query the database to fetch project details
	rows, err := db.Query("SELECT project_id, project_name FROM projects ORDER BY project_id ASC")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch project details"})
		return
	}
	defer rows.Close()

	// Initialize a slice to store project details
	var projects []models.Project_detail

	// Iterate through the rows returned by the query
	for rows.Next() {
		var project models.Project_detail

		// Scan the values from the row into variables
		if err := rows.Scan(&project.ProjectID, &project.ProjectName); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse project details"})
			return
		}

		// Append the project to the slice
		projects = append(projects, project)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to iterate over project details"})
		return
	}

	// Return the fetched project details
	c.IndentedJSON(http.StatusOK, projects)
}

func AddProjectToDB(c *gin.Context, db *sql.DB) {
	var newProject models.Project_detail

	// Bind the new project details from JSON request body
	if err := c.BindJSON(&newProject); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	// Insert the new project into the database
	result, err := db.Exec("INSERT INTO projects (project_name) VALUES (?)", newProject.ProjectName)
	if err != nil {
		fmt.Println("Error inserting project:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to add project"})
		return
	}

	// Get the ID of the newly created project
	projectID, err := result.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve project ID"})
		return
	}

	// Set the project ID in the response
	newProject.ProjectID = int(projectID)

	// Return the created project details
	c.IndentedJSON(http.StatusCreated, newProject)
}
