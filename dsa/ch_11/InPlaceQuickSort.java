/**
 * In Place Quick Sort
 */

import java.util.Random;

public class InPlaceQuickSort
{
  public static void inPlaceQuickSort(int[] original, int a, int b)
  {
    if (a >= b)
      return;

    int pivot = original[b];
    int left = a;
    int right = b - 1;

    while (left <= right)
    {
      while (left <= right && original[left] <= pivot)
        left += 1;
      while (right >= left && original[right] >= pivot)
        right -= 1;
      if (left < right)
      {
        int temp = original[left];
        original[left] = original[right];
        original[right] = temp;
      }
    }

    int temp = original[left];
    original[left] = original[b];
    original[b] = temp;

    inPlaceQuickSort(original, a, left - 1);
    inPlaceQuickSort(original, left + 1, b);
  }

  public static void main(String[] args)
  {
     int[] in = new int[25];
    Random random = new Random();

    for (int i = 0; i < 25; i++)
      in[i] = random.nextInt(50);

    System.out.println("The List");

    for (int i = 0; i < 25; i++)
      System.out.print(in[i] + " ");

    System.out.println("\nQuick-Sort begins");

    InPlaceQuickSort.inPlaceQuickSort(in, 0, in.length - 1);

    System.out.println("\nThe Sorted List");

    for (int i = 0; i < 25; i++)
      System.out.print(in[i] + " ");
    System.out.println("\nFinished");
  }

}
