/**
 * Class representing a geometric progression, subcalss of Progression
 */

public class GeometricProgression extends Progression {

  /** Base */
  protected long base;

  // inherits firstValue and currentValue from Progression

  /** Default constructor, default base is 2 */
  GeometricProgression() {
    this(2);
  }

  /** Alternate constructor, allows input for base */
  GeometricProgression(long base) {
    this.base = base;
    firstValue = 1;
    currentValue = firstValue;
  }

  /** Advances progression by multiplyin the base and the current value
   * @return next value of the progression
   */
  protected long nextValue() {
    currentValue *= base;
    return currentValue;
  }

  // inherits firstValue() and printProgression() from Progression

} //end GeometricProgression class
