package console

import (
	"TodoApp/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// CREATE TASK
func CreateTask() {
	fmt.Print("📝 Enter task title: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	task, err := services.CreateTask(title, false)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	fmt.Printf("✅ Task created: ID=%d, Title='%s'\n", task.ID, task.Title)
}


// GET TASK
func GetTask() {
	fmt.Print("🔍 Enter task ID: ")
	var input string
	fmt.Scanln(&input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("🚫 Invalid ID")
		return
	}

	task, err := services.GetTask(id)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	status := "❌ Not done"
	if task.Done {
		status = "✅ Done"
	}

	fmt.Printf("\n📄 Task Details:\n")
	fmt.Printf("🆔 ID:        %d\n", task.ID)
	fmt.Printf("📌 Title:     %s\n", task.Title)
	fmt.Printf("📅 Created:   %s\n", task.CreatedAt.Format("02 Jan 2006 15:04"))
	fmt.Printf("📍 Status:    %s\n\n", status)
}


// DELETE TASK
func DeleteTask() {
	fmt.Print("🗑️ Enter task ID to delete: ")
	var input string
	fmt.Scanln(&input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("🚫 Invalid ID")
		return
	}

	if err := services.DeleteTask(id); err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	fmt.Println("🗑️ Task deleted successfully")
}

// COMPLETE TASK
func CompleteTask() {
	fmt.Print("✅ Enter task ID to mark as complete: ")
	var input string
	fmt.Scanln(&input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("🚫 Invalid ID")
		return
	}

	task, err := services.CompleteTask(id)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	fmt.Printf("🎉 Task '%s' marked as completed!\n", task.Title)
}

// LIST TASKS
func ListTasks() {
	tasks := services.ListTasks()
	fmt.Println("\n🗂️ Task List:")
	for _, task := range tasks {
		status := "❌ Not Completed"
		if task.Done {
			status = "✅ Completed"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, status)
	}
	fmt.Println()
}
