/**
 * Find the kth smallest term
 */

import java.util.Random;
import java.util.ArrayList;

public class QuickSelect
{
  public static int quickSelect(ArrayList<Integer> array, int k)
  {
    Random random = new Random();

    if (array.size() == 1)
      return array.get(0);

    ArrayList<Integer> lessThan = new ArrayList<Integer>();
    ArrayList<Integer> equalTo = new ArrayList<Integer>();
    ArrayList<Integer> greaterThan = new ArrayList<Integer>();

    int pivot = array.get(random.nextInt(array.size() - 1));

    while (!array.isEmpty())
    {
      if (array.get(0) < pivot)
        lessThan.add(array.remove(0));
      else if (array.get(0) == pivot)
        equalTo.add(array.remove(0));
      else
        greaterThan.add(array.remove(0));
    }
    if (k <= lessThan.size())
      return quickSelect(lessThan, k);
    else if (k <= (lessThan.size() + equalTo.size()))
      return pivot;
    else
      return quickSelect(greaterThan, k - lessThan.size() - equalTo.size());
  }

  public static void main(String[] args)
  {
    int k = 1;

    if (args.length == 1)
      k = Integer.parseInt(args[0]);

    ArrayList<Integer> in = new ArrayList<Integer>();
    Random random = new Random();

    for (int i = 0; i < 25; i++)
      in.add(random.nextInt(50));

    System.out.println("The List");

    for (int i = 0; i < 25; i++)
      System.out.print(in.get(i) + " ");

    System.out.println("\nQuick Select begins for k = " + k);
    int answer = QuickSelect.quickSelect(in, k);

    System.out.println("kth smallest item = " + answer);
    System.out.println("\nFinished");
  }

}
