/**
 * Doubly linked list with head and tail sentinel nodes
 * pointing to the actual head and tail of the list
 * Nodes are DNodes
 */

public class DoublyLinkedList {
  private DNode head, tail;
  private int size;

  public DoublyLinkedList() {
    head = new DNode("", null, null);
    tail = new DNode("", head, null);
    head.setNext(tail);
    size = 0;
  }

  public boolean isEmpty() {
    return size == 0;
  }

  public int size() {
    return size;
  }

  public DNode getFirst() throws IllegalStateException {
    if (isEmpty()) throw new IllegalStateException("List is empty.");
    return head.getNext();
  }


  public DNode getLast() throws IllegalStateException {
    if (isEmpty()) throw new IllegalStateException("List is empty.");
    return tail.getPrevious();
  }

  public void addBefore(DNode beforeNode, DNode addNode) {
    DNode previousNode = beforeNode.getPrevious();
    beforeNode.setPrevious(addNode);
    addNode.setPrevious(previousNode);
    addNode.setNext(beforeNode);
    previousNode.setNext(addNode);
    size++;
  }

  public void addAfter(DNode afterNode, DNode addNode) {
    DNode nextNode = afterNode.getNext();
    afterNode.setNext(addNode);
    addNode.setPrevious(afterNode);
    addNode.setNext(nextNode);
    nextNode.setPrevious(addNode);
    size++;
  }

  public void addFirst(DNode node) {
    addAfter(head, node);
  }

  public void addLast(DNode node) {
    addBefore(tail, node);
  }

  public void remove(DNode removeNode) {
    DNode previous = removeNode.getPrevious();
    DNode next = removeNode.getNext();
    previous.setNext(next);
    next.setPrevious(previous);
    removeNode.setNext(null);
    removeNode.setPrevious(null);
    size--;
  }

  public String contains(String data) {
    if (isEmpty())
      return data + " not contained in list. List is emptpy";

    DNode currentNode = head;
    while (currentNode != tail) {
      if (currentNode.getData() == data)
        return currentNode.getData() + " is in the list.";
      currentNode = currentNode.getNext();
    }
    return data + " not contained in list.";
  }

  public DNode findData(String data) {
    if (isEmpty())
      return null;
    else {
      DNode currentNode = head;
      while (currentNode != tail) {
        if (currentNode.getData() == data)
          return currentNode;
        currentNode = currentNode.getNext();
      }
      return null;
    }
  }

  public String toString() {
    if (isEmpty())
      return "List is empty.";

    DNode currentNode = head.getNext();
    String s = "";
    while (currentNode.getNext() != tail) {
      s += currentNode.toString() + ", ";
      currentNode = currentNode.getNext();
    }
      s += currentNode.toString();

    return s;
  }

  public boolean hasPrevious(DNode node) {
    return node != head;
  }

  public boolean hasNext(DNode node) {
    return node != tail;
  }

  public static void main(String[] args) {
    DoublyLinkedList dll = new DoublyLinkedList();

    System.out.println(dll);
    System.out.println("Add some nodes");
    dll.addFirst(new DNode("B", null, null));
    dll.addLast(new DNode("Z", null, null));
    dll.addAfter(dll.getFirst(), new DNode("C", null, null));
    dll.addBefore(dll.getFirst(), new DNode("A", null, null));
    dll.addBefore(dll.getLast(), new DNode("Y", null, null));

    System.out.println(dll);
    System.out.println("Find some nodes: A then T");
    System.out.println(dll.contains("A"));
    System.out.println(dll.contains("T"));

    System.out.println("\n" + dll);
    System.out.println("Remove some nodes");
    dll.remove(dll.getFirst());
    dll.remove(dll.getLast());
    dll.remove(dll.findData("C"));
    System.out.println(dll);
  }

} // end DLL class


