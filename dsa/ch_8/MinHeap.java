/**
 * MinHeap implementation backed by an ArrayList
 */

import java.util.ArrayList;
import java.util.Random;

public class MinHeap
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

  public MinHeap()
  {
    entries = new ArrayList<Entry>();
    entries.add(0, null);
  }

  public int size()
  {
    return entries.size() - 1;
  }

  public boolean isEmpty()
  {
    return size() == 0;
  }

  public int parentIndex(int i)
  {
    if (i == 1)
      return -1;
    else
      return i / 2;
  }

  public Entry parent(int i)
  {
    return entries.get(parentIndex(i));
  }

  public int leftChildIndex(int i)
  {
    int index = 2 * i;
    if (index > size())
      return -1;
    else
      return index;
  }

  public int rightChildIndex(int i)
  {
    int index = 2 * i + 1;
    if (index > size())
      return -1;
    else
      return index;
  }

  private boolean validIndex(int i)
  {
    return i > 0;
  }

  public boolean isInternal(int i)
  {
    return hasLeft(i);
  }

  public boolean hasRight(int i)
  {
    return rightChildIndex(i) <= size() && rightChildIndex(i) > 1;
  }

  public boolean hasLeft(int i)
  {
    return leftChildIndex(i) <= size() && leftChildIndex(i) > 1;
  }

  public Entry rightChild(int i)
  {
    if (validIndex(rightChildIndex(i)))
      return entries.get(rightChildIndex(i));
    else
      return null;
  }

  public Entry leftChild(int i)
  {
    if (validIndex(leftChildIndex(i)))
      return entries.get(leftChildIndex(i));
    else
      return null;
  }

  public void swap(int i, int j)
  {
    Entry temp = entries.get(i);
    entries.set(i, entries.get(j));
    entries.set(j, temp);
  }

  public void add(Entry e)
  {
    entries.add(e);
    upHeap(e, size());
  }

  public Entry removeMin()
  {
    Entry e = entries.get(1);
    entries.set(1, entries.remove(size()));
    downHeap(entries.get(1), 1);
    return e;
  }

  public void upHeap(Entry e, int i)
  {
    while (parentIndex(i) != -1 && parent(i).getKey() > e.getKey())
    {
      swap(parentIndex(i), i);
      i = parentIndex(i);
    }
  }

  public void downHeap(Entry e, int i)
  {
    int direction = 0;
    while (isInternal(i))
    {
      if (!hasRight(i))
        direction = leftChildIndex(i);
      else if (leftChild(i).getKey() <= rightChild(i).getKey())
        direction = leftChildIndex(i);
      else
        direction = rightChildIndex(i);
      swap(i, direction);
      i = direction;
    }
  }

  public Entry min()
  {
    if (!isEmpty())
      return entries.get(1);
    else
      return null;
  }

  public void explain()
  {
    if (!isEmpty())
    {
      for (int i = 1; i <= size(); i++)
      {
        System.out.println(entries.get(i));
      }
    }
    else
      System.out.println("Empty!");
  }


  public static void main(String[] args)
  {
    MinHeap mHeap = new MinHeap();
    mHeap.explain();
    Random random  = new Random();
    StringBuilder stringBuilder = new StringBuilder();
    for (int i = 0; i < 10; i ++)
    {
      for (int j = 0; j < 5; j++)
      {
        stringBuilder.append((char) ((random.nextInt(Integer.MAX_VALUE) % 26) + 'a'));
      }
      mHeap.add(new Entry(random.nextInt(50) % 50, stringBuilder.toString())) ;
      stringBuilder.delete(0, stringBuilder.length());
    }
    mHeap.explain();
    System.out.println("Size: " + mHeap.size());
    System.out.println("Remove Min: " + mHeap.removeMin() + ", size = " + mHeap.size());
    mHeap.explain();
  }
} // end Heap
