/**
 * Edge List representation of a graph
 */

import java.util.ArrayList;
import java.util.Random;

public class EdgeList extends Graph
{
  // better data structures may make life easier
  protected ArrayList<Vertex> vertices;
  protected ArrayList<Edge> edges;

  public EdgeList(ArrayList<Vertex> vertices, ArrayList<Edge> edges)
  {
    this.vertices = vertices;
    this.edges = edges;
  }

  public EdgeList()
  {
    this( new ArrayList<Vertex>(), new ArrayList<Edge>() );
  }

  public ArrayList<Vertex> vertices()
  {
    return vertices;
  }

  public ArrayList<Edge> edges()
  {
    return edges;
  }

  public Edge removeEdge(Edge edge)
  {
    if (edges.contains(edge))
    {
      Edge removedEdge = edge;
      edges.remove(edge);
      return removedEdge;
    }
    else
      return null;
  }

  public Vertex removeVertex(Vertex vertex)
  {
    if (vertices.contains(vertex))
    {
      ArrayList<Edge> incidentEdges = incidentEdges(vertex);
      for (Edge edge : incidentEdges)
      {
        edges.remove(edge);
      }
    }

    return vertex;
  }

  public ArrayList<Edge> incidentEdges(Vertex vertex)
  {
    ArrayList<Edge> incidentEdges = new ArrayList<Edge>();

    if (vertices.contains(vertex))
    {
      for (Edge edge : edges)
      {
        if (edge.containsVertex(vertex) && !(incidentEdges.contains(edge)))
            incidentEdges.add(edge);
      }
    }

    return incidentEdges;
  }

  public void replaceVertex(Vertex vertex, Vertex newVertex)
  {
    ArrayList<Edge> incidentEdges = new ArrayList<Edge>();

    if (vertices.contains(vertex))
    {
      incidentEdges = incidentEdges(vertex);
      for (Edge edge : incidentEdges)
      {
        edge.replaceVertex(vertex, newVertex);
      }
      vertices.remove(vertex);
      vertices.add(newVertex);
    }
  }
  
  public void replaceEdge(Edge edge, Edge newEdge)
  {
    if (edges.contains(edge))
    {
      newEdge.setStartVertex(edge.getStartVertex());
      newEdge.setEndVertex(edge.getEndVertex());
      edges.remove(edge);
      edges.add(newEdge);
    }
  }

  public void unvisitAllVertices()
  {
    for (Vertex vertex : vertices)
      vertex.unvisit();
  }
  
  public void DFS(Vertex startVertex)
  {
    unvisitAllVertices();
    DFShelper(startVertex);
  }

  public void DFShelper(Vertex startVertex)
  {
  }

  public void BFS(Vertex startVertex)
  {
    unvisitAllVertices();
    BFShelper(startVertex);
  }

  public void BFShelper(Vertex startVertex)
  {
  }

  public static void main(String[] args)
  {
    //make vertices
    ArrayList<Vertex> vertices = new ArrayList<Vertex>();
    ArrayList<Edge> edges = new ArrayList<Edge>();

    for (int i = 0; i < 6; i++)
    {
      char c = (char) (i + 65);
      Vertex vertex = new Vertex(Character.toString(c));
      System.out.println("Made vertex with data = " + vertex.getData());
      vertices.add(vertex);
    }

    //make edges

    float probabilityOfConnection = .3f;
    Random random = new Random();

    for (Vertex vertex : vertices)
    {
      System.out.println("Adding edges");
      while (random.nextFloat() >= probabilityOfConnection)
      {
        Vertex oppositeVertex = vertices.get(random.nextInt(vertices.size()));

        while (oppositeVertex.getData().equals(vertex.getData()))
        {
          oppositeVertex = vertices.get(random.nextInt(vertices.size()));
        }

        Edge edge = new Edge(vertex, oppositeVertex);

        if (!edges.contains(edge))
        {
          System.out.println("Edge: " + vertex.getData() + "<-->" + oppositeVertex.getData() + " + added");
          edges.add(edge);
        }
      }
    }

    EdgeList edgeList = new EdgeList(vertices, edges);
    for (Edge edge : edgeList.edges)
    {
      System.out.println(edge);
    }
  }
} // end class
