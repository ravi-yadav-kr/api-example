package store

const (
	insertQuery = `INSERT INTO employee (id,name, age, dept_name, phone_number, address, joining_date)
                   VALUES ($1, $2, $3, $4, $5, $6,$7);`
	getQuery = `SELECT id, name, age, dept_name, phone_number, address, joining_date FROM employee;`

	updateQuery = `UPDATE employee
                   SET name = $1, 
                       age = $2, 
                       dept_name = $3, 
                       phone_number = $4, 
                       address = $5, 
                       joining_date = $6
                     WHERE id = $7;`
	deleteQuery = `DELETE FROM employee WHERE id = $1;`
)
