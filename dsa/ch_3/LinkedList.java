/**
 * LinkedList class; singly linked here for practice
 */


public class LinkedList {
  protected Node head;
  protected Node tail;
  protected long size;

  public LinkedList() {
    head = null;
    tail = null;
    size = 0;
  }

  public void addFirst(Node node) {
    if (head == null && size == 0) {
      head = node;
      tail = node;
      size++;
    }
    else {
      node.setNext(head);
      head = node;
      size++;
    }
  }

  public void addLast(Node node) {
    node.setNext(null);
    tail.setNext(node);
    tail = node;
    size++;
  }

  public void removeFirst() {
    if (head == null)
      return;
    else if (size == 1) {
      head = null;
      tail = null;
      size--;
    }
    else {
      Node removed = head;
      head = removed.getNext();
      removed.setNext(null);
      size--;
    }
  }

  public String toString() {
    String s = "";

    if (size > 0) {
      Node currentNode = head;
      while (currentNode.getNext() != null) {
        s += currentNode + ", ";
        currentNode = currentNode.getNext();
      }
      s += currentNode;
    }
    else {
      s = "Empty list";
    }
    return s;
  }

  public long getSize() {
    return size;
  }

  public static void main(String[] args) {
    LinkedList ll = new LinkedList();

    System.out.println("List: " + ll + "\nAdd to the front");

    ll.addFirst(new Node("A", null));
    ll.addFirst(new Node("B", null));
    ll.addFirst(new Node("C", null));

    System.out.println("List: " + ll + "\nAdd to the back");

    ll.addLast(new Node("X", null));
    ll.addLast(new Node("Y", null));
    ll.addLast(new Node("Z", null));

    System.out.println("List: " + ll + "\nRemove some nodes");

    ll.removeFirst();
    ll.removeFirst();

    System.out.println("List: " + ll + "\nRemove them all");

    while (ll.getSize() > 0)
      ll.removeFirst();

    System.out.println("List: " + ll);
  }
}
