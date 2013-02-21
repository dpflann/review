/**
 * Simple Queue class utilizing nodes for storage
 */


public class NodeQueue<E> {
  private Node front;
  private Node back;
  private int size;

  public NodeQueue() {
    front = null;
    back = null;
    size = 0;
  }

  public int size() {
    return size;
  }

  public boolean isEmpty() {
    return size == 0;
  }

  public void enqueue(E element) {
    Node<E> node = new Node<E>(element, null);
    if (isEmpty()) {
      front = node;
      back = node;
    }
    else {
      back.setNext(node);
      back = node;
    }
    size++;
  }

  public E dequeue() {
    if (isEmpty()) {
      System.out.println("Queue is empty. Returning null");
      return null;
    }
    else {
      if (size > 1) {
        Node<E> node = front;
        front = node.getNext();
        size--;
        return node.getElement();
      }
      else {
        Node<E> node = front;
        front = null;
        back = null;
        size--;
        return node.getElement();
      }
    }
  }

  public E getFront() {
    if (isEmpty()) { 
      System.out.println("Queue is empty. Returning null");
      return null;
    }
    else {
      Node<E> node = front;
      E element = node.getElement();
      return element;
    }
  }

  public String toString() {
    if (isEmpty()) {
      System.out.println("Queue is empty.");
      return "";
    }
    else {
      String q = "front --> [";
      q += front.toString();
      Node<E> currentNode = front.getNext();
      while (currentNode != null) {
        q += ", " + currentNode.toString();
        currentNode = currentNode.getNext();
      }
      return q + "]";
    }
  }

  public static void main(String[] args) {
    NodeQueue<Integer> queue = new NodeQueue<Integer>();
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
}//end Queue
