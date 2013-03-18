#include <stdio.h>

/* count lines in input */
main()
{
  int c, nl, tabs;
  nl = tabs = 0;
  while ((c = getchar()) != EOF) {
    if (c == '\n') //'\n' is a character, "\n" is a character string, here, we're comparing chars
      ++nl;
    if (c == '\t')
      ++tabs;
  }
  printf("Newlines: %d,\t Tabs:%d\n", nl, tabs);
}

