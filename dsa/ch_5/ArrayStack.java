/**
 * Stack using an array, totally lame, but pedagogical
 */

public class ArrayStack<E> {
  protected int capacity;
  public static final int CAPACITY = 1000;
  protected E S[];
  protected int top = -1;

  public ArrayStack() {
    this(CAPACITY);
  }

  public ArrayStack(int capacity) {
    this.capacity = capacity;
    S = (E[]) new Object[capacity];
  }

  public int size() {
    return top + 1;
  }

  public boolean isEmpty() {
    return (top < 0);
  }

  public void push(E element) {
    if (top == capacity - 1)
      System.out.println("Unable to add element, stack is full");
    else
      S[++top] = element;
  }

  public E pop() {
    if (isEmpty()) {
      System.out.println("Returning null, nothing to pop because stack is empty");
      return null;
    }
    else {
      // allow garbage collection on references
      E element = S[top];
      S[top--] = null;
      return element;
    }
  }

  public E top() {
    if (isEmpty()) {
      System.out.println("Returning null, nothing to pop because stack is empty");
      return null;
    }
    else
      return S[top];
  }

  public String toString() {
    if (isEmpty()) {
      System.out.println("Stack is empty");
      return "[]";
    }
    else {
      String str = "TOP\n[" + S[top];
      for (int i = top-1; i >= 0; i--) 
        str += "\n " + S[i];
      return str + "]\nBOTTOM";
    }
  }

  public static void main(String[] args) {
    ArrayStack<Integer> stack = new ArrayStack<Integer>();
    System.out.println("Current stack: " + stack + "\nPush 4 numbers");
    stack.push(7);
    stack.push(8);
    stack.push(11);
    stack.push(5);
    System.out.println(stack + "\nNow remove, I mean, pop! 2 items!");
    stack.pop();
    stack.pop();
    System.out.println(stack + "\nPush and Pop, then what's on top");
    stack.push(15);
    stack.push(99);
    stack.push(111);
    stack.push(1);
    stack.pop();
    stack.pop();
    System.out.println(stack.top() + " is on top.");
    System.out.println(stack);
    System.out.println(stack.pop() + " popped");
    System.out.println(stack); 
    System.out.println(stack.pop() + " popped");
    System.out.println(stack);
    System.out.println(stack.pop() + " popped");
    System.out.println(stack); 
    System.out.println(stack.pop() + " popped");
    System.out.println(stack);
    System.out.println(stack.pop() + " popped");
    System.out.println(stack); 
  }


} //end ArrayStack
