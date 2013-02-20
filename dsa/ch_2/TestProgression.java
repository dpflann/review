/**
 * Class for testing the numerous progression classes
 */

public class TestProgression {

  public static void main(String[] args) {
    Progression progression;

    // Test ArithmeticProgression
    System.out.println("Arithmetic progression with default increment:");
    progression = new ArithmeticProgression();
    progression.printProgression(10);
    System.out.println("Arithmetic progression with increment 5:");
    progression = new ArithmeticProgression(5);
    progression.printProgression(10);

    // Test GeometricProgression
    System.out.println("Geometric progression with default base:");
    progression = new GeometricProgression();
    progression.printProgression(10);
    System.out.println("Geometric progression with base 3:");
    progression = new GeometricProgression(3);
    progression.printProgression(10);

    // Test FibonacciProgression
    System.out.println("Fibonacci progression with default start values:");
    progression = new FibonacciProgression();
    progression.printProgression(10);
    System.out.println("Fibonacci progression with start values 4 and 6:");
    progression = new FibonacciProgression(4, 6);
    progression.printProgression(10);
  }

} // end TestProgression
