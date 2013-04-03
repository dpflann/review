#include <stdio.h>
#include <string.h>

int str_end(char *s, char *t)
{
  int l_s = strlen(s);
  int l_t = strlen(t);
  s = s + ((l_s) - l_t);
  if (*s == *t)
    return 1;
  else
    return 0;
}

main()
{
  printf("hello vs bye: %d\n", str_end("hello", "bye"));
  printf("hello vs lo: %d\n", str_end("hello", "lo"));
  
}
