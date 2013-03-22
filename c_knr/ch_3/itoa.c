#include <stdio.h>
#include <limits.h>
#include <string.h>
#include "reverse.c"

void itoa(int n, char s[])
{
  int i, sign;

  sign = (n > 0) ? 1 : -1;
  i = 0;
  do
  {
    s[i++] = (n % 10 > 0) ? n % 10 + '0' : -(n % 10) + '0';
  } while ((n /= 10) != 0);

  if (sign < 0)
    s[i++] = '-';
  s[i] = '\0';
  reverse(s);
} // end itoa


main()
{
  //NB: when C calculates the remainder, e.g. -123 % 10, the value returned
  //is negative, same with division. This is against my understanding of 
  //division, where -123 = a * 10 + r, where r > 0, hence -123 = -13 * 10 + 7, a = -13, r = 7. Python, for example, maintains the fact that r > 0.
  char s[100];
  itoa(INT_MIN, s);
  printf("itoa(%d) --> %s\n", INT_MIN, s);
  // Checking C's division and modulus behavior
  printf("-123 mod 10: %d, -123 div 10: %d\n", -123 % 10, -123 / 10);
} // end main
