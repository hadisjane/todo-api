package repositories

import (
	"path/filepath"
	"runtime"

	"TodoApp/errs"
	"TodoApp/models"
	"encoding/json"
	"os"
	"strings"
	"time"
)

var (
	lastID int
	tasks []models.Task
	storageFile string
)

func init() {
	// Get the directory where the current file is located
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	storageFile = filepath.Join(dir, "storage.json")
}


func init() {
	var err error
	tasks, err = loadData()
	if err != nil {
		tasks = []models.Task{} // Initialize empty tasks slice if loading fails
		return
	}

	// Find the highest ID
	for _, t := range tasks {
		if t.ID > lastID {
			lastID = t.ID
		}
	}
}

func saveData() error {
	file, err := os.Create(storageFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print JSON
	return encoder.Encode(tasks)
}

func loadData() ([]models.Task, error) {
	// Check if file exists
	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		// If file doesn't exist, return empty slice
		return []models.Task{}, nil
	}

	file, err := os.Open(storageFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []models.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func CreateTask(title string, done bool) (*models.Task, error) {
	for _, t := range tasks {
		if strings.EqualFold(t.Title, title) {
			return nil, errs.ErrTaskAlreadyExists
		}
	}

	lastID++

	task := models.Task{
		ID:        lastID,
		Title:     title,
		Done:      done,
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, task)
	if err := saveData(); err != nil {
		return nil, err
	}
	return &task, nil
}

func ListTasks() []models.Task {
	return tasks
}

func GetTask(id int) (*models.Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errs.ErrTaskNotFound
}

func CompleteTask(id int) (*models.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			if tasks[i].Done {
				return nil, errs.ErrTaskAlreadyCompleted
			}
			tasks[i].Done = true
			if err := saveData(); err != nil {
				return nil, err
			}
			return &tasks[i], nil
		}
	}
	return nil, errs.ErrTaskNotFound
}

func DeleteTask(id int) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := saveData(); err != nil {
				return err
			}
			return nil
		}
	}
	return errs.ErrTaskNotFound
}
