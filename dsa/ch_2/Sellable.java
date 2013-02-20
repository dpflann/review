/** Interface for objects that can be sold */

public interface Sellable {

  /** description of the object */
  public String description();

  /** list price in cents */
  public int listPrice();

  /** lowest price acceptable price in cents */
  public int lowestPrice();


} // end Sellable interface
