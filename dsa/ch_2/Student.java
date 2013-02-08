/**
 * Class representing a student, implementing Person interface
 */

public class Student implements Person {
  
  private String id;
  private String name;
  private int age;

  public Student(String id, String name, int age) {
    this.id = id;
    this.name = name;
    this.age = age;
  }

  protected int studyHours() { 
    return age / 2;
  }

  public String getID() {
    return id;
  }

  public String getName() {
    return name;
  }

  public int getAge() {
    return age;
  }

  public boolean equalTo(Person other) {
    Student otherStudent = (Student) other;
    return (id.equals(otherStudent.getID()));
  }

  public String toString() {
    return "Student(ID: " + id + ", Name: " + name + ", Age: " + age +")";
  }


} // end Student class
