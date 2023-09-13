#include <iostream>
#include <vector>
#include <string>

class Task {
private:
    std::string description;
    bool completed;

public:
    Task(const std::string& desc) : description(desc), completed(false) {}

    void markCompleted() {
        completed = true;
    }

    bool isCompleted() const {
        return completed;
    }

    std::string getDescription() const {
        return description;
    }
};

class ToDoList {
private:
    std::vector<Task> tasks;

public:
    void addTask(const std::string& description) {
        tasks.push_back(Task(description));
    }

    void markTaskCompleted(int index) {
        if (index >= 0 && index < tasks.size()) {
            tasks[index].markCompleted();
        }
    }

    void displayTasks() const {
        std::cout << "To-Do List:" << std::endl;
        for (size_t i = 0; i < tasks.size(); ++i) {
            std::cout << (i + 1) << ". ";
            if (tasks[i].isCompleted()) {
                std::cout << "[X] ";
            } else {
                std::cout << "[ ] ";
            }
            std::cout << tasks[i].getDescription() << std::endl;
        }
    }
};

int main() {
    ToDoList toDoList;

    while (true) {
        std::cout << "Menu:" << std::endl;
        std::cout << "1. Add Task" << std::endl;
        std::cout << "2. Mark Task as Completed" << std::endl;
        std::cout << "3. Display Tasks" << std::endl;
        std::cout << "4. Exit" << std::endl;

        int choice;
        std::cout << "Enter your choice: ";
        std::cin >> choice;

        if (choice == 1) {
            std::string description;
            std::cout << "Enter task description: ";
            std::cin.ignore(); // Clear the newline character from the buffer
            std::getline(std::cin, description);
            toDoList.addTask(description);
        } else if (choice == 2) {
            int index;
            std::cout << "Enter task index to mark as completed: ";
            std::cin >> index;
            toDoList.markTaskCompleted(index - 1);
        } else if (choice == 3) {
            toDoList.displayTasks();
        } else if (choice == 4) {
            break;
        } else {
            std::cout << "Invalid choice. Please try again." << std::endl;
        }
    }

    return 0;
}
