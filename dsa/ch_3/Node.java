/**
 * Node class, used for now in LinkedLists
 * Contains a String as data
 */

public class Node {
  private String data;
  private Node next;

  public Node(String data, Node node) {
    this.data = data;
    next = node;
  }

  public String getData() {
    return data;
  }

  public Node getNext() {
    return next;
  }

  public void setData(String newData) {
    data = newData;
  }

  public void setNext(Node newNext) {
    next = newNext;
  }

  public String toString() {
    return "[data = " + data + "]";
  }
}

