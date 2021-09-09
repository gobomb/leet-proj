package leetcode

/**
 * 690. Employee Importance

 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	var rs int
	for _, e := range employees {
		if e.Id == id {
			rs += e.Importance

			for _, s := range e.Subordinates {
				rs += getImportance(employees, s)
			}
			break
		}
	}
	return rs
}
