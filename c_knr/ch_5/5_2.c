#include <stdio.h>
#include <ctype.h>
#include "ch.c"

/* get_float: get next float from input into *pn */

int get_float(float *fn)
{
  int c, sign;
  while (isspace(c = getch()))
      ;
  if (!isdigit(c) && c != EOF && c != '-' && c != '+')
  {
    ungetch(c);
    return 0;
  }

  sign = (c == '-') ? -1 : 1;
  if (c == '+' || c == '-')
    c = getch();
  if (!isdigit(c))
  {
    ungetch(c);
    return 0;
  }
  for (*fn = 0; isdigit(c); c = getch())
    *fn = 10 * *fn + (c - '0');
  if (c == '.')
    c = getch();
  float divider = 1.0;
  for (*fn; isdigit(c); c = getch(), divider *= 10)
    *fn = 10 * *fn + (c - '0');
  *fn = *fn * sign / divider;
  if (c != EOF)
    ungetch(c);
  return c;
} //end getfloat

main()
{
  float fn;
  if (get_float(&fn))
    printf("%f\n", fn);
}



