/**
 * Quick-Sort
 */
import java.util.Random;
import java.util.ArrayList;

public class QuickSort {

  public static void quickSort(ArrayList<Integer> in)
  {
    Random random = new Random();
    if (in.size() <= 1)
      return;

    int pivot = in.get(random.nextInt(in.size()));

    ArrayList<Integer> lessThan = new ArrayList<Integer>();
    ArrayList<Integer> equalTo = new ArrayList<Integer>();
    ArrayList<Integer> greaterThan = new ArrayList<Integer>();

    while (!in.isEmpty())
    {
      if (in.get(in.size() - 1) < pivot)
        lessThan.add(in.remove(in.size() - 1));
      else if (in.get(in.size() - 1) == pivot)
        equalTo.add(in.remove(in.size() - 1));
      else
        greaterThan.add(in.remove(in.size() - 1));
    }

    quickSort(lessThan);
    quickSort(greaterThan);

    while (!lessThan.isEmpty())
      in.add(lessThan.remove(0));
    while (!equalTo.isEmpty())
      in.add(equalTo.remove(0));
    while (!greaterThan.isEmpty())
      in.add(greaterThan.remove(0));

    return;
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

    System.out.println("\nQuick-Sort begins");

    QuickSort.quickSort(in);

    System.out.println("\nThe Sorted List");

    for (int i = 0; i < 25; i++)
      System.out.print(in.get(i) + " ");
    System.out.println("\nFinished");
  }
}
