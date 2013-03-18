#include <stdio.h>

/*
 * Slimmed down Fahrenheit to Celsius table
 */
//Avoid using "magic numbers" use symbolic constants: #define <name> <replacement text>

#define LOWER 0 //<-- there is no semicolon
#define UPPER 300
#define STEP 20

main()
{
  int fahr;

  for (fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP)
    printf("%3d\t%6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));

  printf("\nNow, in reverse!\n");
  for (fahr = UPPER; fahr >= LOWER; fahr = fahr - STEP)
    printf("%3d\t%6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
}
