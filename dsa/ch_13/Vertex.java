/**
 * Vertex in a graph
 */

public class Vertex
{
  protected String data;
  protected enum State { UNVISITED, VISITED }
  protected State state;

  public Vertex(String data)
  {
    this.data = data;
    state = State.UNVISITED;
  }

  public Vertex()
  {
    this(null);
  }

  public void setData(String data)
  {
    this.data = data;
  }

  public String getData()
  {
    return data;
  }

  public State getState()
  {
    return state;
  }

  public void setState(State state)
  {
    this.state = state;
  }

  public void visit()
  {
    this.state = State.VISITED;
  }

  public void unvisit()
  {
    this.state = State.UNVISITED;
  }

  public String toString()
  {
    return "[" + data + "]";
  }

  public boolean equals(Object o)
  {
    Vertex v = (Vertex) o;
    return data.equals(v.getData());
  }

  public static void main(String[] args)
  {
    Vertex vertex = new Vertex("Hello");
    System.out.println(vertex.state);
    vertex.visit();
    System.out.println(vertex.getState());
  }

} // end class
