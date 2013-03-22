#include <stdio.h>

/* strlen: return length of s */
int str_len(char s[]);

int str_len(char s[])
{
  int i = 0;
  while (s[i] != '\0')
    ++i;
  return i;
} // end strlen


main()
{
  char* s = "hello whirled";
  printf("strlen(\"%s\") = %d\n", s, str_len(s));
} // end main
