/**
 * Iterative MergeSort
 */

import java.util.Random;

public class IterativeMergeSort
{
  public static void mergeSort(int[] original)
  {
    int[] in = new int[original.length];
    System.arraycopy(original, 0, in, 0, in.length);
    int[] out = new int[original.length];
    int[] temp;

    int n = in.length;
    for (int i = 1; i < n; i *= 2)
    {
      for (int j = 0; j < n; j += 2 * i)
      {
        merge(in, out, j, i);
      }
      temp = in;
      in = out;
      out = temp;
    }
    System.arraycopy(in, 0, original, 0, in.length);
  }

  public static void merge(int[] in, int[] out, int start, int increment)
  {
    int x = start;
    int end1 = Math.min(start + increment, in.length);
    int end2 = Math.min(start + 2 * increment, in.length);
    int y = start + increment;
    int z = start;
    while ((x < end1) && (y < end2))
    {
      if (in[x] <= in[y])
        out[z++] = in[x++];
      else
        out[z++] = in[y++];
    }
    if (x < end1)
      System.arraycopy(in, x, out, z, end1 - x);
    else if (y < end2)
      System.arraycopy(in, y, out, z, end2 - y);
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

    System.out.println("\nMerge-Sort begins");

    IterativeMergeSort.mergeSort(in);

    System.out.println("\nThe Sorted List");

    for (int i = 0; i < 25; i++)
      System.out.print(in[i] + " ");
    System.out.println("\nFinished");
  }

}
