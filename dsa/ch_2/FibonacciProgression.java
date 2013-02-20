/**
 * Class representing a Fibonacci progression, subclass of Progression
 */

public class FibonacciProgression extends Progression {
  /** Previous value */
  long previousValue;

  // inherits firstValue and currentValue

  /** Default constructor, setting 0 and 1 as first two values */
  FibonacciProgression() {
    this(0, 1);
  }


  /** Alternate constructor, providing the first and second values */
  FibonacciProgression(long value1, long value2) {
    firstValue = value1;
    previousValue = value2 - value1; //ficitious value preceding the first
  }

  /** Advances the progression by adding the previous value to the current value
   * @return next value of the progression
   */
  protected long nextValue() {
    long temp = previousValue;
    previousValue = currentValue;
    currentValue += temp;

    return currentValue;
  }

  // inherits firstValue() and printProgression() from Progression


} // end FibonacciProgression
