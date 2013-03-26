/**
 * Simple HashTable
 */

import java.util.ArrayList;
import java.util.Random;
import java.util.Iterator;

public class HashTable
{
  protected int capacity;
  protected ArrayList<Entry>[] buckets;
  protected int entries;
  protected double loadThreshold;

  private class Entry
  {
    String key; //made key for practice writing hashing function
    String value;

    private Entry(String k, String v)
    {
      key = k;
      value = v;
    }
    private Entry()
    {
      key = "";
      value = "";
    }
    private String getKey() { return key; }
    private String getValue() { return value; }
    public void setKey(String k) { key = k; }
    public void setValue(String v) { value = v; }
    public String toString() { return "<" + key + "-->" + value + ">"; }
  }

  public HashTable()
  {
    this(100);
  }

  public HashTable(int capacity)
  {
    this.capacity = capacity;
    buckets = new ArrayList[capacity];
    loadThreshold = .9;

    for (int i = 0; i < capacity; i++)
    {
      buckets[i] = new ArrayList<Entry>();
    }
  }

  public int size()
  {
    return entries;
  }

  public int hash(String key)
  {
    int cyclicallyHashed = cyclicHash(key);
    return compress(cyclicallyHashed);
  }

  /* may want to change compression behavior */
  public int compress(int hash)
  {
    return (hash % capacity);
  }

  public int cyclicHash(String key)
  {
    int hash = 0;
    int shift = 5;
    for (int i = 0; i < key.length(); i++)
    {
      hash = (hash << 5) | (hash >> 32 - shift);
      hash += (int) key.charAt(i);
    }
    return hash;
  }

  public double loadFactor()
  {
    return (double) entries / (double) capacity;
  }

  public void rehash()
  {
    capacity *= 2;
    ArrayList<Entry>[] oldBuckets = buckets;
    buckets = new ArrayList[capacity];
    for (int i = 0; i < capacity; i++)
    {
      buckets[i] = new ArrayList<Entry>();
    }
    for (int i = 0; i < oldBuckets.length; i++)
    {
      if (!oldBuckets[i].isEmpty())
      {
        Iterator<Entry> iterator = oldBuckets[i].iterator();
        while (iterator.hasNext())
        {
          put(iterator.next());
        }
      }
    }
  }

  public void put(Entry entry)
  {
    if (loadFactor() >= loadThreshold)
      rehash();
    int hash = hash(entry.getKey());
    buckets[hash].add(entry);
    entries++;
  }

  public void put(String key, String value)
  {
    if (loadFactor() >= loadThreshold)
      rehash();
    int hash = hash(key);
    buckets[hash].add(new Entry(key, value));
    entries++;
  }

  public Entry find(String key, String value)
  {
    int hash = hash(key);
    if (buckets[hash].isEmpty())
    {
      return null;
    }
    else
    {
      ArrayList<Entry> entries = buckets[hash];
      for (int i = 0; i < entries.size(); i++)
      {
        if (entries.get(i).getValue().compareTo(value) == 0)
        {
          return entries.get(i);
        }
      }
      return
        null;
    }
  }

  public Entry remove(String key, String value)
  {
    int hash = hash(key);
    if (buckets[hash].isEmpty())
    {
      return null;
    }
    else
    {
      ArrayList<Entry> entries = buckets[hash];
      for (int i = 0; i < entries.size(); i++)
      {
        if (entries.get(i).getValue().compareTo(value) == 0)
        {
          this.entries--;
          return entries.remove(i);
        }
      }
      return
        null;
    }
  }

  public String toString()
  {
    String s = "";
    for (int i = 0; i < capacity; i++)
    {
      if (!buckets[i].isEmpty())
      {
        s += "\n" + i + ":\n\t";
        Iterator<Entry> iterator = buckets[i].iterator();
        while (iterator.hasNext())
          s += iterator.next().toString() + "\n\t";
      }
    }
    return s;
  }

  public static void main(String[] args)
  {
    HashTable hashTable = new HashTable(10);
    System.out.println(hashTable);

    Random random  = new Random();
    StringBuilder stringBuilder = new StringBuilder();
    for (int i = 0; i < 10; i ++)
    {
      for (int j = 0; j < 5; j++)
      {
        stringBuilder.append((char) ((random.nextInt(Integer.MAX_VALUE) % 26) + 'a'));
      }
      String key = Character.toString((char) (random.nextInt(10) + 'a'));
      String value = stringBuilder.toString();
      hashTable.put(key, value) ;
      stringBuilder.delete(0, stringBuilder.length());
    }
    hashTable.put("a", "hello");
    System.out.println(hashTable);
    System.out.println("removed: " + hashTable.find("a", "hello"));
    hashTable.remove("a", "hello");
    System.out.println(hashTable);
  }
} //end HashTable
