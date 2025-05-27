package console

import (
	"fmt"
)

func ShowCommands() {
	for {
		fmt.Print(`
📋 ToDo App - Available Commands:
0️⃣  Exit
1️⃣  ➕ Add Task
2️⃣  🔍 View Task
3️⃣  ❌ Delete Task
4️⃣  ✅ Complete Task
5️⃣  📃 List All Tasks

🧠 Enter command number: `)

		var command int
		fmt.Scan(&command)

		switch command {
		case 0:
			fmt.Println("👋 Goodbye!")
			return
		case 1:
			CreateTask()
		case 2:
			GetTask()
		case 3:
			DeleteTask()
		case 4:
			CompleteTask()
		case 5:
			ListTasks()
		default:
			fmt.Println("🚫 Unknown command! Try again.")
		}
	}
}
