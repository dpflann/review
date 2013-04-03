#include <stdio.h>

void str_n_cpy(char *s, char *t, int n)
{
  while (n--)
  {
    *s++ = *t++;
  }
}

void str_n_cat(char *s, char *t, int n)
{
  while (*s)
    s++;

  while (n--)
  {
    *s++ = *t++;
  }
}

int str_n_cmp(char *s, char *t, int n)
{
  for (n; *s == *t && n > 0; s++, t++, n--)
    if (*s == '\0')
      return 0;
  return *s - *t;
}

main()
{
  char t[8] = "goodbye";
  char s[6] = "hello";
  
  char x[4] = "abc";
  char y[6] = "defgh";

  /*
   * Encountered strange string buffer,
   * string pointer behavior

    x: abc, y: defgh
    "str_n_cat(x, y, 3)"
    x: abcdeoodbye, y: defgh
    s: hello, t: eoodbye
    "str_n_cpy(s, t, 3)"
    s: eoolo, t: eoodbye

  */

  printf("x: %s, y: %s\n", x, y);
  printf("\"str_n_cat(x, y, 3)\"\n");
  str_n_cat(x, y, 2);
  printf("x: %s, y: %s\n", x, y);
  
  printf("s: %s, t: %s\n", s, t);
  str_n_cpy(s, t, 3);
  printf("\"str_n_cpy(s, t, 3)\"\n");
  printf("s: %s, t: %s\n", s, t);

  printf("hello vs. hellish, 3 -> %d\n", str_n_cmp("hello", "hellish", 3));
}
