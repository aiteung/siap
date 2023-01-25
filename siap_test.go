package siap

import (
	"fmt"
	"log"
	"testing"
)

func TestAtapi(t *testing.T) {
	OpenDb()
	// Create employee
	createID, err := CreateEmployee("Jake", "United States")
	if err != nil {
		log.Fatal("Error creating Employee: ", err.Error())
	}
	fmt.Printf("Inserted ID: %d successfully.\n", createID)

	// Read employees
	count, err := ReadEmployees()
	if err != nil {
		log.Fatal("Error reading Employees: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

	// Update from database
	updatedRows, err := UpdateEmployee("Jake", "Poland")
	if err != nil {
		log.Fatal("Error updating Employee: ", err.Error())
	}
	fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

	// Delete from database
	deletedRows, err := DeleteEmployee("Jake")
	if err != nil {
		log.Fatal("Error deleting Employee: ", err.Error())
	}
	fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)
}
