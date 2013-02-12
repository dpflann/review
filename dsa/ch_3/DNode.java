/**
 * A node presentation for Doubly Linked List.
 * Contains references to previous and next nodes and data (which in this
 * case will be String for pedagogical simplicity)
 * 
 */


public class DNode {
  private DNode previous ;
  private DNode next ;
  private String data;

  public DNode() {
    data = "";
    previous = null;
    next = null;
  }

  public DNode(String data, DNode previous, DNode next) {
    this.data = data;
    this.previous = previous;
    this.next = next;
  }

  public DNode getNext() {
    return next;
  }
  public DNode getPrevious() {
    return previous;
  }

  public String getData() {
    return data;
  }

  public void setPrevious(DNode newPrevious) {
    previous = newPrevious;
  }

  public void setNext(DNode newNext) {
    next = newNext;
  }

  public void setData(String newData) {
    data = newData;
  }

  public String toString() {
    return "(Data: " + data + ")";
  }
} // end DNode class
