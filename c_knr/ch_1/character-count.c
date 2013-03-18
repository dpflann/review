#include <stdio.h>

/* count characters in input; 1st version */

main()
{
  long nc; // long is at least 32 bits
  /*
  nc = 0;
  while (getchar() != EOF)
    ++nc;
  printf("%ld\n", nc); //ld to represent long number to be output
  */

  // as a for-loop
  double nc_d;
  for (nc_d = 0; getchar() != EOF; nc_d++)
    ; // grammatical rules of C require for-loop to have a body; hence the 'null statement' aka semicolon to satisfy. Placed on separate line for visibility
  printf("%.0f\n", nc_d);
}
