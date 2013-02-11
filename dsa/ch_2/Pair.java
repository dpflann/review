/**
 * class representing a Pair of items (Key, Value)
 */

public class Pair<K, V> {
  K key;
  V value;

  /* Unsafe - the size of the objects to be used as K, V are unknown
   * at compile time, they are only determined at run time,
   * This makes assumptions about object size, and can create problems
   * It is deemd unsafe
  public Pair(K k, V v) {
    key = k;
    value = v;
  }*/

  public void set(K key, V value) {
    this.key = key;
    this.value = value;
  }

  public K getKey() {
    return key;
  }

  public V getValue() {
    return value;
  }

  public String toString() {
    return "[" + getKey() + ", " + getValue() + "]";
  }

  public static void main(String[] agrs) {

    Pair<String, Integer> pair1 = new Pair< String, Integer>();
    pair1.set(new String("height"), new Integer(36));
    System.out.println(pair1);
    Pair<Student, Double> pair2 = new Pair<Student, Double>();
    pair2.set(new Student("A5976", "Sue", 19), new Double(9.5));
    System.out.println(pair2);

    int[] x = new int[10];
    //Pair<String, Integer>[] pairs = new Pair<String, Integer>[10]; //You can't do that on television, no parameterized type arrays

  }

} // end Pair class
