#include <stdio.h>

/* repalces spaces by tabs + spaces */
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
    if (c == ' ')
    {
      // this the beginning of a sequence of blanks of size at least 1 blank
      int blanks = 1;
      while ((c = getchar()) == ' ' && characters_in_column <= n)
      {
        ++blanks;
        ++characters_in_column;
      }
      while ((blanks / 5) > 0)
      {
        putchar('\t');
        blanks = blanks / 5;
      }
      int i;
      for (i = 0; i < blanks; i++)
        putchar(' ');
      if (characters_in_column == n)
      {
        characters_in_column = 0;
        ++column;
      }
      else if (c != ' ')
        putchar(c);
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

