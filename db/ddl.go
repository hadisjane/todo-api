package db

func InitMigrations() error {
	tasksTable := `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		done BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(tasksTable)
	if err != nil {
		return err
	}
	return nil
}

func Seed() error {
	_, err := db.Exec(`
		INSERT INTO tasks (title, done)
		VALUES ($1, $2)`, "task 1", true)
	return err
}
