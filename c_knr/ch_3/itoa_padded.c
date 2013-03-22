#include <stdio.h>
#include <limits.h>
#include <string.h>
#include "reverse.c"

void itoa_padded(int n, char s[], int minimum_width)
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
  if (strlen(s) < minimum_width)
  {
    int blanks = minimum_width - strlen(s);
    for (i; i < minimum_width - 1; i++)
      s[i] = ' ';
    s[i] = '\0';
  }
  reverse(s);
} // end itoa


main()
{
  char s[100];
  itoa_padded(1234, s, 7);
  printf("itoa(%d, s, 7) --> (%s)\n", 1234, s);

  itoa_padded(1234, s, 4);
  printf("itoa(%d, s, 4) --> (%s)\n", 1234, s);
} // end main
