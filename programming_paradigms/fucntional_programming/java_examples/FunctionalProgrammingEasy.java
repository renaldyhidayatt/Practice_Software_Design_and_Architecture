package programming_paradigms.fucntional_programming.java_examples;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class FunctionalProgrammingEasy {
    public static void main(String[] args) {
        List<Integer> numbers = Arrays.asList(1, 2, 3, 4, 5);

        List<Integer> squaredNumbers = numbers.stream().map(n -> n * n).collect(Collectors.toList());

        System.out.println("Original numbers: " + numbers);
        System.out.println("Squared numbers: " + squaredNumbers);
    }
}
