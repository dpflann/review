#include <stdio.h>
#include <stdlib.h>

main()
{
  union u_tag
  {
    int ival;
    float fval;
  } u;

  u.fval = 3.5;
  printf("%f\n", u.fval);
  u.ival = 5;
  printf("%d\n", u.ival);
}// end main
