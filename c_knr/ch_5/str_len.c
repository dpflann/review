#include <stdio.h>

int str_len(char *s)
{
  int n = 0;
  while (*s++ != '\0')
    n++;
  return n;
}

void str_cpy(char *t, char *s)
{
  while ((*t++ = *s++)) // comparison to '\0' unncessary
    ;
}

int str_cmp(char *s, char *t)
{
  for (; *s == *t; s++, t++)
    if (*s == '\0')
      return 0;
  return *s - *t;
}

main()
{
  char *s;// = "hello";
  if (s == 0)
  {
    printf("unequal\n");
    s = "hello";
  }
  printf("hello: %d\n", str_len(s));
  char *t = s + 1;
  char *a = "bye";
  char *b = a + 2;
  if (t > s)
    printf("t > s\n");
  if (a > t)
    printf("a > t\n");
  char x[] = "hello";
  char *y = "abcdefgh";
  str_cpy(x, y);
  printf("x: %s\n", x);
  a = "hello";
  b = "hell";
  printf("hello vs. hell: %s\n", (str_cmp(a, b) > 0) ? "greater" : ((str_cmp(a, b) == 0) ? "equal" : "less than"));

}
