#include <stdio.h>

/* remove comments from a C file */

main()
{
  int c;
  while ((c = getchar()) != EOF)
  {
    if (c == '/')
    {
      // line comment: //...
      if ((c = getchar()) == '/')
      {
        while ((c = gethchar()) != '\n')
            ;
      }
      // beginning of block comment /*... which must be closed by */
      else if (c == '*')
      {
        while ((c = getchar()) != '*' || c != '\n')
            ;
        if (c == '*' && (c = getchar()) == '/')
          ;
       }
      else
      {
        putchar(c);
      }
    }
    else
      putchar(c);
  }
} // end main
