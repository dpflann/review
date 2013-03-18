#include <stdio.h>

/* 
 * print Fahrenheit-Celsius table for fahr = 0, 20, ..., 300
 */

main() 
{
  /*
   * Declarations: <type> <variable name>
   */
  float fahr, celsius;
  int lower, upper, step;

  /*
   * Assignment statements
   */
  lower = 0; // lower limit of temperature table
  upper = 300; // upper limit
  step = 20; // step size

  celsius = lower;
  printf("Convering Celsius to Fahrenheit\n%3c\t%6c\n%s", 'C', 'F', "---------------\n");
  while (celsius <= upper) {
    // decimal point in constant indicates floating point
    /*
     * If an arithmetic operator has integer operands, an integer operation is performed
     * If an arithmetic operator has one floating-point operand and one integer operand, the integer
     * operand will be converted to a floating-point before the operation is done. 
     */
    fahr = (9.0 / 5.0) * (celsius) + 32.0; // integer division truncates
    // print fahr in a field that is 3 digits wide, celsisu in a field that is 6 digits wide
    printf("%3.0f\t%6.1f\n", celsius, fahr); //printf is not actually part of the C language, ANSI standard defines this
    celsius = celsius + step;
  }
}


