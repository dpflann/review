/**
 * Edge in a graph
 * Currently, undirected
 */

import java.util.ArrayList;

public class Edge
{
  protected Vertex endPointOne;
  protected Vertex endPointTwo;
  protected int weight;
  protected boolean isDirected; // sufficient for making a subclass?
  protected State state;
  protected EdgeType edgeType;

  // one --> two if directed
  public Edge(Vertex vertexOne, Vertex vertexTwo, int edgeWeight, boolean isDirected)
  {
    endPointOne = vertexOne;
    endPointTwo = vertexTwo;
    weight = edgeWeight;
    this.isDirected = isDirected;
    state = State.UNVISITED;
    edgeType = EdgeType.UNDETERMINED;
  }

  public Edge(Vertex vertexOne, Vertex vertexTwo)
  {
    this(vertexOne, vertexTwo, 0, false);
  }

  public Edge(Vertex vertexOne, Vertex vertexTwo, boolean isDirected)
  {
    this(vertexOne, vertexTwo, 0, isDirected);
  }

  public Edge(Vertex vertexOne, Vertex vertexTwo, int edgeWeight)
  {
    this(vertexOne, vertexTwo, edgeWeight, false);
  }

  public Edge(int edgeWeight)
  {
    this(null, null, edgeWeight, false);
  }

  public Edge(int edgeWeight, boolean isDirected)
  {
    this(null, null, edgeWeight, isDirected);
  }

  public Edge()
  {
    this(null, null, 0, false);
  }

  public int getWeight()
  {
    return weight;
  }

  public Vertex getEndVertex()
  {
    return endPointTwo;
  }

  public Vertex getStartVertex()
  {
    return endPointOne;
  }

  public boolean isDirected()
  {
    return isDirected;
  }

  public ArrayList<Vertex> getVertices()
  {
    ArrayList<Vertex> vertices = new ArrayList<Vertex>();
    vertices.add(endPointOne);
    vertices.add(endPointTwo);

    return vertices;
  }

  public Vertex getOpposite(Vertex vertex)
  {
    if (containsVertex(vertex))
    {
      if (isStartVertex(vertex))
        return endPointTwo;
      else
        return endPointOne;
    }

    return null;
  }

  public void setWeight(int edgeWeight)
  {
    weight = edgeWeight;
  }

  public void setEndVertex(Vertex endVertex)
  {
    endPointTwo = endVertex;
  }

  public void setStartVertex(Vertex startVertex)
  {
    endPointOne = startVertex;
  }

  public void setDirected(boolean directed)
  {
    isDirected = directed;
  }

  public boolean containsVertex(Vertex vertex)
  {
    return endPointOne.getData().equals(vertex.getData()) || endPointTwo.getData().equals(vertex.getData());
  }

  public boolean isStartVertex(Vertex vertex)
  {
    return endPointOne.getData().equals(vertex.getData());
  }

  public boolean isEndVertex(Vertex vertex)
  {
    return endPointTwo.getData().equals(vertex.getData());
  }

  public void replaceVertex(Vertex vertexToReplace, Vertex replacementVertex)
  {
    if (containsVertex(vertexToReplace))
    {
      if (isStartVertex(vertexToReplace))
        endPointOne = replacementVertex;
      else
        endPointTwo = replacementVertex;
    }
  }

  public boolean isUnvisited()
  {
    return state == State.UNVISITED;
  }

  public void setAsDiscoveryEdge()
  {
    edgeType = EdgeType.DISCOVERY;
  }

  public void setAsBackEdge()
  {
    edgeType = EdgeType.BACK;
  }

  public void setAsCrossEdge()
  {
    edgeType = EdgeType.CROSS;
  }

  public void visit()
  {
    state = State.VISITED;
  }

  public void unvisit()
  {
    state = State.UNVISITED;
  }

  public boolean isDiscovery()
  {
    return edgeType == EdgeType.DISCOVERY;
  }

  public String toString()
  {
    return ("Edge: " + endPointOne + "<-->" + endPointTwo);
  }

  public boolean equals(Object o)
  {
    Edge e = (Edge) o;
    return (endPointOne.equals(e.getStartVertex()) && endPointTwo.equals(e.getEndVertex()) ||
        endPointTwo.equals(e.getStartVertex()) && endPointOne.equals(e.getEndVertex()));
  }
} // end class
