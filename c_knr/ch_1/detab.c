#include <stdio.h>

/* replace tabs by spaces */
#define n 10 /* column width is 10 characaters */
#define COLUMNS 5

main()
{
  int c;
  int characters_in_column = 1;
  int column = 1;
  while ((c = getchar()) != EOF)
  {
    if (c == '\n' || column > 5)
    {
      putchar('\n');
      characters_in_column = 0;
      column = 1;
    }
    if (c == '\t')
    {
      int i;
      for (i = 0; i < n - characters_in_column; ++i)
      {
        putchar(' ');
      }
      ++column;
    }
    else
    {
      putchar(c);
      ++characters_in_column;
      if (characters_in_column == 10)
      {
        characters_in_column = 0;
        ++column;
      }
    }
  }
}
