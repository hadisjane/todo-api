package db

import (
	"fmt"
	"log"
)

func InitMigrations() error {
	db := GetDB()
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	var err error

	// Сначала создаем таблицу пользователей
	usersTable := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
	log.Println("Creating users table if not exists...")

	// Создаем таблицу пользователей
	if _, err := db.Exec(usersTable); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Создаем функцию для генерации последовательностей для каждого пользователя
	if _, err := db.Exec(`
	CREATE OR REPLACE FUNCTION create_user_sequence_if_not_exists(user_id INT) RETURNS VOID AS $$
	BEGIN
		EXECUTE format('CREATE SEQUENCE IF NOT EXISTS task_id_seq_%s MINVALUE 1 START WITH 1', user_id);
	END;
	$$ LANGUAGE plpgsql;`); err != nil {
		return fmt.Errorf("failed to create sequence function: %w", err)
	}

	// Создаем функцию для получения следующего ID задачи для пользователя
	if _, err := db.Exec(`
	CREATE OR REPLACE FUNCTION next_task_id(user_id INT) RETURNS INT AS $$
	DECLARE
		next_id INT;
	BEGIN
		PERFORM create_user_sequence_if_not_exists(user_id);
		EXECUTE format('SELECT nextval(''task_id_seq_%s'')', user_id) INTO next_id;
		RETURN next_id;
	END;
	$$ LANGUAGE plpgsql;`); err != nil {
		return fmt.Errorf("failed to create next_task_id function: %w", err)
	}

	// Затем создаем таблицу задач с составным первичным ключом
	tasksTable := `CREATE TABLE IF NOT EXISTS tasks (
		id INT NOT NULL,
		user_id INT REFERENCES users(id) NOT NULL,
		title VARCHAR(255) NOT NULL,
		done BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (user_id, id)
	)`

	log.Println("Creating tasks table if not exists...")
	// Создаем таблицу задач
	if _, err := db.Exec(tasksTable); err != nil {
		return fmt.Errorf("failed to create tasks table: %w", err)
	}

	log.Println("Database migrations completed successfully")

	// Create sequences for users who don't have one yet
	_, err = db.Exec(`
	DO $$
	DECLARE
		r RECORD;
	BEGIN
		FOR r IN SELECT id FROM users LOOP
			-- Only create sequence if it doesn't exist
			EXECUTE format('CREATE SEQUENCE IF NOT EXISTS task_id_seq_%s', r.id);
		END LOOP;
	END $$;`)
	if err != nil {
		log.Printf("Warning: failed to create user sequences: %v", err)
	}
	return nil
}