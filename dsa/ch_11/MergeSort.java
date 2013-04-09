/**
 * MergeSort
 */

import java.util.ArrayList;
import java.util.Random;


public class MergeSort
{

  public static void mergeSort(ArrayList<Integer> in)
  {
    int n = in.size();

    if (n<= 1)
      return;

    ArrayList<Integer> in1 = new ArrayList<Integer>();
    ArrayList<Integer> in2 = new ArrayList<Integer>();

    int i = 0;

    while (i < n / 2)
    {
      in1.add(in.remove(0));
      i++;
    }
    while (!in.isEmpty())
    {
      in2.add(in.remove(0));
    }

    mergeSort(in1);
    mergeSort(in2);
    merge(in1, in2, in);
  }

  public static void merge(ArrayList<Integer> in1, ArrayList<Integer> in2, ArrayList<Integer> in)
  {
    while (!in1.isEmpty() && !in2.isEmpty())
    {
      if (in1.get(0) <= in2.get(0))
        in.add(in1.remove(0));
      else
        in.add(in2.remove(0));
    }

    while (!in1.isEmpty())
      in.add(in1.remove(0));

    while (!in2.isEmpty())
      in.add(in2.remove(0));
  }

  public static void main(String[] args)
  {
    ArrayList<Integer> in = new ArrayList<Integer>();
    Random random = new Random();

    for (int i = 0; i < 25; i++)
      in.add(random.nextInt(50));

    System.out.println("The List");

    for (int i = 0; i < 25; i++)
      System.out.print(in.get(i) + " ");

    System.out.println("\nMerge-Sort begins");

    MergeSort.mergeSort(in);

    System.out.println("\nThe Sorted List");

    for (int i = 0; i < 25; i++)
      System.out.print(in.get(i) + " ");
    System.out.println("\nFinished");
  }

}
