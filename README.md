# rago
Relational Algebra implementation example in Go

## Example output

```
( employee.id )( employee.name )( employee.dept )
| 1 || Alice || 1 |
| 2 || Bob || 2 |
| 3 || Connor || 1 |
| 4 || David || 2 |
| 5 || Ethan || 1 |
( department.id )( department.name )
| 1 || Research |
| 2 || Production |
( employee.name )( department.name )
| Alice || Research |
| Bob || Production |
| Connor || Research |
| David || Production |
| Ethan || Research |
```
