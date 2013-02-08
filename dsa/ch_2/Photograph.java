/**
 * Class for sellable Photographs, implements Sellable interface
 */

public class Photograph implements Sellable {
  
  private String description;
  private int price;
  private boolean inColor; // true if photograph is in color

  public Photograph(String aDescription, int aPrice, boolean isInColor) {
    description = aDescription;
    price = aPrice;
    inColor = isInColor;
  }

  public String description() {
    return description;
  }

  public int listPrice() {
    return price;
  }

  public int lowestPrice() {
    return price / 2;
  }

  public boolean isColor() {
    return inColor;
  }

} // end Photograph class
