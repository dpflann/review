#include <stdio.h>
#include "reverse.c"

// convert int x into base b string representation 
void itob(int x, char s[], int b)
{
  int r, i;
  i = 0;
  while (x > 0)
  {
    r = x % b;
    r = (r >= 10 && r <= 15) ? r - 10 + 'a' : r + '0';
    x = x / b;
    s[i++] = r;
  }
  s[i] = '\0';
  reverse(s);
} // end itob

main()
{
  char s[100];
  itob(123, s, 16);
  printf("itob(123, s, 16) --> %s\n", s);

  itob(123, s, 8);
  printf("itob(123, s, 8) --> %s\n", s);

  itob(123, s, 2);
  printf("itob(123, s, 2) --> %s\n", s);
} // end main
