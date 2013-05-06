/**
 * Graph class
 */

import java.util.ArrayList;

public abstract class Graph
{
  // better data structures may make life easier
  public abstract ArrayList<Vertex> vertices();

  public abstract ArrayList<Edge> edges();

  public abstract Edge removeEdge(Edge edge);

  public abstract Vertex removeVertex(Vertex vertex);

  public abstract ArrayList<Edge> incidentEdges(Vertex vertex);

  public abstract void replaceVertex(Vertex vertex, Vertex newVertex);

  public abstract void replaceEdge(Edge edge, Edge newEdge);

  public abstract void DFS(Vertex startVertex);

  public abstract void BFS(Vertex startVertex);
} // end class
