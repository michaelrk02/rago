package main

import (
    "os"
)

func main() {
    var rel *relation

    employee := makeRelation([][]string{
        []string{"employee.id", "employee.name", "employee.dept"},
        []string{"1", "Alice", "1"},
        []string{"2", "Bob", "2"},
        []string{"3", "Connor", "1"},
        []string{"4", "David", "2"},
        []string{"5", "Ethan", "1"}})

    department := makeRelation([][]string{
        []string{"department.id", "department.name"},
        []string{"1", "Research"},
        []string{"2", "Production"}})

    employee.print(os.Stdout)
    department.print(os.Stdout)

    rel = join(employee, department, func(t *tuple) bool {
        return t.values["employee.dept"] == t.values["department.id"]
    })
    rel = projection(rel, []string{"employee.name", "department.name"})
    rel.print(os.Stdout)
}

