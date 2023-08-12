package programming_paradigms.fucntional_programming.java_examples;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class FunctionalIntermediate {
    public static void main(String[] args) {
        List<String> names = Arrays.asList("Alice", "Bob", "Charlie", "David", "Eve");

        List<String> filteredNames = names.stream()
                .filter(name -> name.length() > 4)
                .map(String::toUpperCase)
                .collect(Collectors.toList());

        System.out.println("Original names: " + names);
        System.out.println("Filtered and uppercase names: " + filteredNames);
    }
}
