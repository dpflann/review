#include <stdio.h>

/* copy input to output; 1st version */

main()
{
  int c; // use int because E0F is larger than a char, c must be big enough to hold any value that getchar(..) returns
  /*
  c = getchar();
  while (c != EOF) {
    putchar(c);
    c = getchar();
  }*/

  // more concisely
  // Parantheses are required because '!=' has higher precedence than '=' (assignment)
  c = getchar();
  printf("getchar() != EOF: %d\n", c != EOF);
  printf("Value of EOF: %d\n", EOF);
  while ((c = getchar()) != EOF) {
    putchar(c);
  }

}
