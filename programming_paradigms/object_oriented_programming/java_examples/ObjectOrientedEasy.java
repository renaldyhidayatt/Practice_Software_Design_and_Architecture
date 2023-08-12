package programming_paradigms.object_oriented_programming.java_examples;

class Car {
    String brand;
    String model;
    int year;

    Car(String brand, String model, int year) {
        this.brand = brand;
        this.model = model;
        this.year = year;
    }

    void displayInfo() {
        System.out.println("Brand: " + brand + ", Model: " + model + ", Year: " + year);
    }
}

public class ObjectOrientedEasy {
    public static void main(String[] args) {
        Car car = new Car("Toyota", "Camry", 2022);
        car.displayInfo();
    }
}
