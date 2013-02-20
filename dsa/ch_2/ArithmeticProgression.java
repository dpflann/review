
/**
 * Class for representing arithmetic progressions
 */

public class ArithmeticProgression extends Progression {

  /** Increment */
  protected long increment;

  // firstValue and currentValue inherited

  /** Default constructor, provides 1 as default increment */
  ArithmeticProgression() {
    this(1);
  }

  /** Alternate constructor, allows input for increment */
  ArithmeticProgression(long increment) {
    this.increment = increment;
  }

  /** Advances progression by adding the increment to the current value
   * @return next value of the progression
   */
  protected long nextValue() {
    currentValue += increment;

    return currentValue;
  }

  // firstValue and printProgression inherited from Progression


}//end ArithmeticProgression class

