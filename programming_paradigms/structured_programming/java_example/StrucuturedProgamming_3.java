package java_example;

import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

interface TaskSerializable extends Serializable {
    String getTaskDetails();
}

class Task implements TaskSerializable {
    private String title;
    private boolean completed;

    public Task(String title) {
        this.title = title;
        this.completed = false;
    }

    public String getTitle() {
        return title;
    }

    public boolean isCompleted() {
        return completed;
    }

    public void markAsCompleted() {
        completed = true;
    }

    @Override
    public String getTaskDetails() {
        return "[" + (completed ? "X" : " ") + "] " + title;
    }
}

class TaskManager {
    private List<Task> tasks;
    private String dataFilePath;

    public TaskManager(String dataFilePath) {
        this.dataFilePath = dataFilePath;
        tasks = new ArrayList<>();
        loadTasks();
    }

    public void addTask(String title) {
        tasks.add(new Task(title));
        saveTasks();
    }

    public void markTaskAsCompleted(int index) {
        if (index >= 0 && index < tasks.size()) {
            tasks.get(index).markAsCompleted();
            saveTasks();
        }
    }

    public void listTasks() {
        System.out.println("To-Do List:");
        for (int i = 0; i < tasks.size(); i++) {
            System.out.println(i + ". " + tasks.get(i).getTaskDetails());
        }
    }

    private void loadTasks() {
        try (ObjectInputStream inputStream = new ObjectInputStream(new FileInputStream(dataFilePath))) {
            tasks = (List<Task>) inputStream.readObject();
        } catch (IOException | ClassNotFoundException e) {
            // If there's an error, simply start with an empty task list
            tasks = new ArrayList<>();
        }
    }

    private void saveTasks() {
        try (ObjectOutputStream outputStream = new ObjectOutputStream(new FileOutputStream(dataFilePath))) {
            outputStream.writeObject(tasks);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

public class StrucuturedProgamming_3 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        TaskManager taskManager = new TaskManager("tasks.dat");

        while (true) {
            System.out.println("Menu:");
            System.out.println("1. Add Task");
            System.out.println("2. Mark Task as Completed");
            System.out.println("3. List Tasks");
            System.out.println("4. Exit");
            System.out.print("Choose an option: ");
            int choice = scanner.nextInt();
            scanner.nextLine();

            switch (choice) {
                case 1:
                    System.out.print("Enter task title: ");
                    String title = scanner.nextLine();
                    taskManager.addTask(title);
                    break;
                case 2:
                    System.out.print("Enter task index to mark as completed: ");
                    int index = scanner.nextInt();
                    taskManager.markTaskAsCompleted(index);
                    break;
                case 3:
                    taskManager.listTasks();
                    break;
                case 4:
                    System.out.println("Exiting...");
                    scanner.close();
                    System.exit(0);
                    break;
                default:
                    System.out.println("Invalid choice. Please choose again.");
            }
        }
    }
}
