/**
 *  Examing recursion through Fibonacci Sequences
 */

public class Recursion {

  public static int binaryFibonacci(int k) {
    if (k <= 1)
      return k;
    else
      return binaryFibonacci(k - 1) + binaryFibonacci(k - 2);
  }

  private static int[] linearFibonacciHelper(int k) {
    int[] tup = new int[2];
    if (k <= 1) {
      // (k, 0)
      tup[0] = k;
      tup[1] = 0;
      return tup;
    }
    else {
      // (i, j)
      tup = linearFibonacciHelper(k - 1);
      int i = tup[0];
      tup[0] = tup[0] + tup[1];
      tup[1] = i;
      // (i + j, i)
      return tup;
    }
  }

  public static int linearFibonacc(int k) {
    int[] tup = linearFibonacciHelper(k);
    return tup[0];
  }

  public static void main(String[] args) {
    System.out.println("Binary Fibonacci");
    for (int i=1; i < 10; i++) {
      System.out.println("binaryFibonacci(" + i + ") = " + Recursion.binaryFibonacci(i));
    }

    System.out.println("\nLinearFibonacci");
    for (int i=1; i < 10; i++) {
      System.out.println("linearFibonacci(" + i + ") = " + Recursion.linearFibonacc(i));
    }
  }

}
