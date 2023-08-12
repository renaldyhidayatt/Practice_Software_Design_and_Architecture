package programming_paradigms.object_oriented_programming.java_examples;

class Student {
    private String name;
    private int age;
    private String studentId;

    public Student(String name, int age, String studentId) {
        this.name = name;
        this.age = age;
        this.studentId = studentId;
    }

    public String getName() {
        return name;
    }

    public int getAge() {
        return age;
    }

    public String getStudentId() {
        return studentId;
    }

    public void displayInfo() {
        System.out.println("Name: " + name + ", Age: " + age + ", Student ID: " + studentId);
    }
}

public class ObjectOrientedIntermediate {
    public static void main(String[] args) {
        Student student = new Student("Alice", 20, "12345");
        student.displayInfo();
    }
}
