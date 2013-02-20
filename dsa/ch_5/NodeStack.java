/**
 * Stack using nodes instead of an array
 */

public class NodeStack<E> {
  private int size;
  private Node<E> top;

  public NodeStack() {
    size = 0;
    top = null;
  }

  public int size() {
    return size;
  }

  public boolean isEmpty() {
    return size == 0;
  }

  public void push(E element) {
    if (isEmpty()) {
      Node node = new Node<E>(element, null);
      top = node;
    }
    else {
      Node node = new Node<E>(element, top);
      top = node;
    }
    size++;
  }

  public E pop() {
    if (isEmpty()) {
      System.out.println("Stack is empty, returning null");
      return null;
    }
    else {
      E element = top.getElement();
      Node remove = top;
      top = top.getNext();
      remove.setNext(null);
      size--;

      return element;
    }
  }

  public E top() {
    if (isEmpty()) {
      System.out.println("Stack is empty, returning null");
      return null;
    }
    else {
     return top.getElement();
    }
  }

  public String toString() {
    if (isEmpty()) {
      System.out.println("Stack is empty");
      return "[]";
    }
    else {
      String str = "[" + top.getElement();
      Node<E> currentNode = top.getNext();
      while (currentNode != null) {
        str += "\n " + currentNode.getElement();
        currentNode = currentNode.getNext();
      }
      return str + "]";
    }
  }

  public static void main(String[] args) {
    NodeStack<Integer> stack = new NodeStack<Integer>();
    System.out.println(stack + "\nPush");
    stack.push(5);
    stack.push(77);
    stack.push(-1);
    stack.push(1000);
    System.out.println("Size: " + stack.size() + "\n" + stack);
    int size = stack.size();
    for (int i = 1; i <= size;i++)
      System.out.println("Popped: " + stack.pop());
    System.out.println(stack);    
  }

}//end NodeStack
