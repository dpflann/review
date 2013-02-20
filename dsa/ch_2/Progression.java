/**
 * Class for numeric progressions, and exemplifying inheritance
 */

public class Progression {

  /** First value of the progression */
  protected long firstValue;

  /** Current value of the progression */
  protected long currentValue;

  Progression() {
    currentValue = firstValue = 0;
  }

  /** Resets progression to first value
   * @return first value
   */
    protected long firstValue() {
      currentValue = firstValue;
      return currentValue;
    }

    /** Advances progression to next value
     * @return next value in the progression
     */
    protected long nextValue() {
      return ++currentValue;
    }

    /** prints the first n values of the progression
     * @param n number of values to print
     */
    public void printProgression(int n) {
      System.out.print(firstValue());
      for (int i = 2; i <= n; i++)
        System.out.print(" " + nextValue());
      System.out.println();
    }

} //end Progression class
