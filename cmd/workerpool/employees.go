package main

type Employee struct {
	Name   string
	Salary float64
}

var employees = make([]Employee, 15)

func initEmployees() {
	employees[0] = Employee{
		Name:   "Иван",
		Salary: 45000,
	}

	employees[1] = Employee{
		Name:   "Василий",
		Salary: 64000,
	}

	employees[2] = Employee{
		Name:   "Сергей",
		Salary: 35000,
	}

	employees[3] = Employee{
		Name:   "Семён",
		Salary: 84000,
	}

	employees[4] = Employee{
		Name:   "Олег",
		Salary: 112500,
	}

	employees[5] = Employee{
		Name:   "Евгений",
		Salary: 75000,
	}

	employees[6] = Employee{
		Name:   "София",
		Salary: 55000,
	}

	employees[7] = Employee{
		Name:   "Мария",
		Salary: 22000,
	}

	employees[8] = Employee{
		Name:   "Наталья",
		Salary: 300000,
	}

	employees[9] = Employee{
		Name:   "Кирилл",
		Salary: 1,
	}

	employees[10] = Employee{
		Name:   "Василий",
		Salary: 23000,
	}

	employees[11] = Employee{
		Name:   "Томара",
		Salary: 89000,
	}

	employees[12] = Employee{
		Name:   "Мага",
		Salary: 135000,
	}

	employees[13] = Employee{
		Name:   "Алексей",
		Salary: 67000,
	}

	employees[14] = Employee{
		Name:   "Айгуль",
		Salary: 44000,
	}
}
