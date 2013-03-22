#include <stdio.h>

// concatenate strings s and t, t appended to s
// assume's s has enough space for t
void str_cat(char s[], char t[]);

void str_cat(char s[], char t[])
{
  int i, j;
  i = j = 0;
  while (s[i] != '\0')
    i++;
  while ((s[i++] = t[j++]) != '\0')
    ;
} // end str_cat


main()
{
  char s[11] = "hello";
  char t[4] = "bye";

  printf("%s, %s\n", s, t);
  str_cat(s, t);
  printf("%s\n", s);
} // end main
