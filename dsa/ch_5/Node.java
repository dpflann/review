/**
 * Generic Node
 */

public class Node<E> {

  private E element;
  private Node<E> next;

  public Node() {
    this(null, null);
  }

  public Node(E element, Node<E> node) {
    this.element = element;
    next = node;
  }

  public Node<E> getNext() {
    return next;
  }
  
  public void setNext(Node<E> node) {
    next = node;
  }

  public E getElement() {
    return element;
  }

  public void setElement(E element) {
    this.element = element;
  }

  public String toString() {
    if (element != null)
      return element.toString();
    else
      return "";
  }

}//end Node
