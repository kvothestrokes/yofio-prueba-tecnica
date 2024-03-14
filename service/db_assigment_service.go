package service

import (
	"database/sql"
	"fmt"
	"yofio-prueba-tecnica/config"
)

type CreditAssignerDB interface {
	SaveCreditAssignation(invest_amount int32, success int) error
	GetTotalAssignations() (int, error)
	GetTotalSuccessAssignations() (int, error)
	GetTotalFailedAssignations() (int, error)
	GetAverageSuccessAssignations() (float32, error)
	GetAverageFailedAssignations() (float32, error)
}

type creditAssignerDB struct {
	db *sql.DB
}

func NewCreditAssignerDB() CreditAssignerDB {
	return &creditAssignerDB{
		db: config.SetupDB(),
	}
}

func (c *creditAssignerDB) SaveCreditAssignation(invest_amount int32, success int) error {

	fmt.Println("Saving credit assignation")
	_, err := c.db.Exec("INSERT INTO credit_assigments_history (invest_amount, success) VALUES (?, ?)", invest_amount, success)

	if err != nil {
		fmt.Println("Error saving credit assignation ", err)
		return err
	}

	return nil

}

func (c *creditAssignerDB) GetTotalAssignations() (int, error) {

	rows, err := c.db.Query("SELECT COUNT(*) FROM credit_assigments_history")

	if err != nil {
		return 0, err
	}

	var total int
	for rows.Next() {
		rows.Scan(&total)
	}

	return total, nil

}

func (c *creditAssignerDB) GetTotalSuccessAssignations() (int, error) {

	rows, err := c.db.Query("SELECT COUNT(*) FROM credit_assigments_history WHERE success = 1")

	if err != nil {
		return 0, err
	}

	var total int
	for rows.Next() {
		rows.Scan(&total)
	}

	return total, nil

}

func (c *creditAssignerDB) GetTotalFailedAssignations() (int, error) {

	rows, err := c.db.Query("SELECT COUNT(*) FROM credit_assigments_history WHERE success = 0")

	if err != nil {
		return 0, err
	}

	var total int
	for rows.Next() {
		rows.Scan(&total)
	}

	return total, nil

}

func (c *creditAssignerDB) GetAverageSuccessAssignations() (float32, error) {

	rows, err := c.db.Query("SELECT AVG(invest_amount) FROM credit_assigments_history WHERE success = 1")

	if err != nil {
		return 0, err
	}

	var average float32
	for rows.Next() {
		rows.Scan(&average)
	}

	return average, nil

}

func (c *creditAssignerDB) GetAverageFailedAssignations() (float32, error) {

	rows, err := c.db.Query("SELECT AVG(invest_amount) FROM credit_assigments_history WHERE success = 0")

	if err != nil {
		return 0, err
	}

	var average float32
	for rows.Next() {
		rows.Scan(&average)
	}

	return average, nil

}
