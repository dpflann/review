/**
 * Class for objects that can be sold, packed, and shipped
 */

public class BoxedItem implements Sellable, Transportable {
  
  private String description;
  private int price;
  private int weight;
  private boolean isHazardous;
  private int height = 0;
  private int width = 0;
  private int depth = 0;

  public BoxedItem(String aDescription, int aPrice, int aWeight, boolean isItHazardous) {
    description = aDescription;
    price = aPrice;
    weight = aWeight;
    isHazardous = isItHazardous;
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

  public int weight() {
    return weight;
  }

  public boolean isHazardous() {
    return isHazardous;
  }

  public int insuredValue() {
    return price * 2;
  }

  public void setBox(int aHeight, int aWidth, int aDepth) {
    height = aHeight;
    width = aWidth;
    depth = aDepth;
  }

} // end BoxedItem class
