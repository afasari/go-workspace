package model

type Todo struct {
	Id   uint64 `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func CreateTodo(todo string) error {
	statement := `insert into todos(todo, done) values($1, $2)`

	_, err := db.Query(statement, todo, false)
	return err
}

func GetAllTodos() ([]Todo, error) {
	todos := []Todo{}

	statement := `select id, todo, done from todos;`

	rows, err := db.Query(statement)
	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.Id, &todo.Todo, &todo.Done)
		if err != nil {
			return todos, err
		}

		todos = append(todos, todo)
	}

	return todos, err
}

func GetTodo(id uint64) (Todo, error) {
	statement := `select todo, done from todos where id=$1;`

	todo := Todo{}
	todo.Id = id

	row, err := db.Query(statement, id)
	if err != nil {
		return todo, err
	}

	for row.Next() {
		err := row.Scan(&todo.Todo, &todo.Done)
		if err != nil {
			return todo, err
		}
	}

	return todo, err
}

func MarkDone(id uint64) error {
	todo, err := GetTodo(id)
	if err != nil {
		return err
	}

	statement := `update todos set done=$2 where id=$1;`
	_, err = db.Query(statement, id, !todo.Done)

	return err
}

func Delete(id uint64) error {
	statement := `delete from todos where id=$1;`

	_, err := db.Exec(statement, id)
	return err
}
