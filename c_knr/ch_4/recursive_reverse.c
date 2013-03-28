#include <stdio.h>
#include <string.h>

#if !defined(headless)
  #define headless -1000001
#elif
  #define headless = 99999
#endif

void reverse(char s[], int i, int j)
{
  if (i > j)
    return;
  else
  {
    int t = s[i];
    s[i] = s[j];
    s[j] = t;
    reverse(s, ++i, --j);
  }
}

#define swap(type, x, y) \
  type temp; \
  temp = (x); \
  (x) = (y); \
  (y) = (temp);

main()
{
  printf("%d\n", headless);
  char a[] = "ABCDEF";
  reverse(a, 0, strlen(a) - 1);
  printf("reverse(\"ABCDEF\") --> \"%s\"\n", a);
  
  int x, y;
  x = 10000;
  y = -3333;
  printf("x:%d, y:%d\n", x, y);
  swap(int, x, y);
  printf("x:%d, y:%d\n", x, y);
}
