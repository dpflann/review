/**
 * Simple array based Queue class
 */

public class ArrayQueue<E> {
  private E[] array;
  private static final int CAPACITY = 1000;
  private int capacity;
  private int front = 0;
  private int back = 0;
  private int size = 0;

  public ArrayQueue() {
    this(CAPACITY);
  }

  public ArrayQueue(int capacity) {
    array = (E[]) new Object[capacity];
    this.capacity = capacity;
  }

  public int size() {
    return size;
  }

  public boolean isEmpty() {
    return size == 0;
  }

  public void enqueue(E element) {
    if (isEmpty()) {
      front = 0;
      back = 0;
      array[front] = element;
    }
    else {
      if (size == capacity) {
        int newCapacity = capacity * 2;

        E[] newArray = (E[]) new Object[newCapacity];
        int i = front;
        int newIndex = 0;
        for (; i != back; i = (i + 1) % capacity) {
          newArray[newIndex] = array[i];
          newIndex++;
        }
        newArray[newIndex] = array[back];
        capacity *= 2;
        back = newIndex;
        array = newArray;
      }
      System.out.println("Back: " + back);
      back = (back + 1) % capacity;
      array[back] = element;
    }
    size++;
  }

  public E dequeue() {
    if (isEmpty()) {
      System.out.println("Queue is empty. Returning null");
      size--;
      return null;
    }
    else {
      E element = array[front];
      array[front] = null;
      front = (front + 1) % capacity;
      size--;
      return element;
    }
  }

  public E getFront() {
    if (isEmpty()) {
      System.out.println("Queue is empty. Returning null");
      return null;
    }
    else {
      E element = array[front];
      return element;
    }
  }

  public String toString() {
    if (isEmpty()) {
      System.out.println("Queue is empty.");
      return "[]";
    }
    else {
      String q = "front --> [";
      if (size == 1) {
        q = "[" + array[front] + "]";
      }
      else {
        q += array[front];
        for (int i = front + 1; i != back; i = (i + 1) % capacity) {
          if (array[i] != null)
            q += ", " + array[i];
        }
        q += ", " + array[back] + "]";
      }
      return q;
    }
  }

  public static void main(String[] args) {
    ArrayQueue<Integer> queue = new ArrayQueue<Integer>(3);
    System.out.println(queue);
    queue.enqueue(9);
    queue.enqueue(234);
    queue.enqueue(-1);
    queue.enqueue(12);
    queue.enqueue(44);
    System.out.println("Is queue empty? " + queue.isEmpty());
    System.out.println(queue + ", " + queue.size() + " nodes");
    System.out.println(queue.dequeue() + " dequeued");
    System.out.println(queue.dequeue() + " dequeued");
    System.out.println(queue.dequeue() + " dequeued");
    System.out.println(queue + ", " + queue.size() + " nodes");
    System.out.println("Front = " + queue.getFront());
    System.out.println(queue.dequeue() + " dequeued");
    System.out.println(queue.dequeue() + " dequeued");
    System.out.println(queue.dequeue() + " dequeued");
  }


} //end arrayQueue
