/**
 * Edge List representation of a graph
 */

import java.util.ArrayList;
import java.util.Random;
import java.util.LinkedList;

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

  public void unvisitAllEdges()
  {
    for (Edge edge : edges)
      edge.unvisit();
  }

  public void DFS(Vertex startVertex)
  {
    unvisitAllVertices();
    unvisitAllEdges();
    DFShelper(startVertex);
  }

  public void DFShelper(Vertex startVertex)
  {
    startVertex.visit();

    ArrayList<Vertex> adjacentVertices = new ArrayList<Vertex>();

    for (Edge edge : incidentEdges(startVertex))
    {
      if (edge.isUnvisited())
      {
        edge.visit();
        Vertex adjacentVertex = edge.getOpposite(startVertex);
        if (adjacentVertex.isUnvisited())
        {
          edge.setAsDiscoveryEdge();
          DFShelper(adjacentVertex);
        }
        else
        {
          edge.setAsBackEdge();
        }
      }
    }
  }

  public void BFS(Vertex startVertex)
  {
    unvisitAllVertices();
    unvisitAllEdges();
    BFShelper(startVertex);
  }

  public void BFShelper(Vertex startVertex)
  {
    LinkedList<Vertex> level = new LinkedList<Vertex>();
    level.add(startVertex);
    while (!level.isEmpty())
    {
      Vertex currentVertex = level.remove();
      for (Edge edge : incidentEdges(currentVertex))
      {
        if (edge.isUnvisited())
        {
          edge.visit();
          Vertex adjacentVertex = edge.getOpposite(currentVertex);

          if (adjacentVertex.isUnvisited())
          {
            adjacentVertex.visit();
            edge.setAsDiscoveryEdge();
            level.add(adjacentVertex);
          }
          else
          {
            edge.setAsBackEdge();
          }
        }
      }
    }
  }

  // similar to Erdos-Renyi graph construction
  public static EdgeList makeSimpleGraph(int numberOfVertices, float probabilityOfConnection)
  {
    System.out.println("**********\nCreating a graph with " + numberOfVertices + " vertices and probability of creating an edge between vertices " + probabilityOfConnection + "\n**********");

    ArrayList<Vertex> vertices = new ArrayList<Vertex>();
    ArrayList<Edge> edges = new ArrayList<Edge>();

    //make vertices
    for (int i = 0; i < numberOfVertices; i++)
    {
      char c = (char) (i + 65);
      Vertex vertex = new Vertex(Character.toString(c));
      System.out.println("Made vertex with data = " + vertex.getData());
      vertices.add(vertex);
    }

    //make edges
    Random random = new Random();

    for (Vertex vertex : vertices)
    {
      System.out.println("Adding edges");
      while (random.nextFloat() <= probabilityOfConnection)
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
    return new EdgeList(vertices, edges);
  }

  public static void main(String[] args)
  {

    EdgeList edgeList = makeSimpleGraph(6, .7f);

    for (Edge edge : edgeList.edges)
    {
      System.out.println(edge);
    }

    edgeList.DFS(edgeList.vertices.get(0));
    System.out.println("*************\nPerformed DFS\n*************");

    for (Edge edge : edgeList.edges)
    {
      if (edge.isDiscovery())
        System.out.println(edge);
    }

    System.out.println("Visited, hence reachable vertices");
    for (Vertex vertex : edgeList.vertices)
    {
      if (!vertex.isUnvisited())
        System.out.println(vertex);
    }
    System.out.println("Unvisited, hence unreachable vertices");
    for (Vertex vertex : edgeList.vertices)
    {
      if (vertex.isUnvisited())
        System.out.println(vertex);
    }

    edgeList.BFS(edgeList.vertices.get(0));
    System.out.println("*************\nPerformed BFS\n*************");

    for (Edge edge : edgeList.edges)
    {
      if (edge.isDiscovery())
        System.out.println(edge);
    }

    System.out.println("Visited, hence reachable vertices");
    for (Vertex vertex : edgeList.vertices)
    {
      if (!vertex.isUnvisited())
        System.out.println(vertex);
    }
    System.out.println("Unvisited, hence unreachable vertices");
    for (Vertex vertex : edgeList.vertices)
    {
      if (vertex.isUnvisited())
        System.out.println(vertex);
    }
  }
} // end class
