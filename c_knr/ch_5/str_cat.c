#include <stdio.h>

void str_cat(char *s, char *t)
{
  while (*s)
    s++;
    //;
  //--s;
  while (*s++ = *t++)
    ;
}


main()
{
  char s[] = "abc";
  char t[] = "defghi";
  printf("s:, %s, t: %s\n", s, t);
  printf("\"str_cat(s, t)\"\n");
  str_cat(s, t);
  printf("s: %s\n", s);
}
