/**
 * ArrayList backed PriorityQueue
 */

import java.util.ArrayList;
import java.util.Random;

public class PriorityQueue
{
  protected ArrayList<Entry> entries;

  private static class Entry implements Comparable<Entry>
  {
    int key;
    String value;

    public Entry(int key, String value)
    {
      this.key = key;
      this.value = value;
    }
    public int getKey()
    {
      return key;
    }

    public String getValue()
    {
      return value;
    }

    public int compareTo(Entry otherEntry)
    {
      if (key < otherEntry.getKey())
        return -1;
      else if (key > otherEntry.getKey())
        return 1;
      else
        return 0;
    }

    public String toString()
    {
      return "[" +  key + " -> " + value + "]";
    }
  } // end Entry

  public PriorityQueue()
  {
    entries = new ArrayList<Entry>();
  }

  public int size()
  {
    return entries.size();
  }

  public boolean isEmpty()
  {
    return entries.isEmpty();
  }

  public Entry min()
  {
    if (!isEmpty())
      return entries.get(0);
    else
      return null;
  }

  public void insert(int key, String value)
  {
    Entry entry = new Entry(key, value);
    if (isEmpty())
      entries.add(entry);
    else if (entries.get(entries.size() - 1).compareTo(entry) < 0)
      entries.add(entry);
    else
    {
      int i = 0;
      while (entries.get(i).compareTo(entry) < 0)
        i++;
      // ArrayList add @ index will shift current value at index to right, and all subsequent to right
      entries.add(i, entry);
    }
  }

  public Entry removeMin()
  {
    if (!isEmpty())
      return entries.remove(0);
    else
      return null;
  }

  public void explain()
  {
    if (!isEmpty())
    {
      for (Entry entry : entries)
      {
        System.out.println(entry);
      }
    }
    else
      System.out.println("Empty!");
  }

  public static void main(String[] args)
  {
    PriorityQueue pq = new PriorityQueue();
    pq.explain();
    Random random  = new Random();
    StringBuilder stringBuilder = new StringBuilder();
    for (int i = 0; i < 10; i ++)
    {
      for (int j = 0; j < 5; j++)
      {
        stringBuilder.append((char) ((random.nextInt(Integer.MAX_VALUE) % 26) + 'a'));
      }
      pq.insert(random.nextInt(50) % 50, stringBuilder.toString()) ;
      stringBuilder.delete(0, stringBuilder.length());
    }
    pq.explain();
    System.out.println("Size: " + pq.size());
    System.out.println("Remove Min: " + pq.removeMin() + ", size = " + pq.size());
    pq.explain();
    System.out.println("Insert [34 -> \"abc\"]");
    pq.insert(34, "abcde");
    pq.explain();
    System.out.println("Min: " + pq.min());
  }
} // end PriorityQueue
