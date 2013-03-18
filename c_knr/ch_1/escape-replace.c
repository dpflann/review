#include <stdio.h>

/* Output the input, but replace each tab by '\t', each backspace by '\b', and each backslash by \\ */
main()
{
  int c;
  while ((c = getchar()) != EOF) {
    if (c == '\t')
      printf("\\t");
    if (c == '\b')
      printf("\\b");
    if (c == '\\')
      printf("\\\\");
    else
      putchar(c);
  }
}


