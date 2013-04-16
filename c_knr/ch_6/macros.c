#include <stdio.h>

#define floater(decimal) 111.##decimal

main()
{
  /* bit fields of given size, initialized to 0 in this case */
  struct {
    unsigned int is_keyword : 1;
    unsigned int is_static : 2;
  } flags;

  flags.is_keyword = 1;
  flags.is_static = 3;

  printf("flags: is_keyword: %x\n", flags.is_keyword);
  printf("flags: is_static: %x\n", flags.is_static);

  int x = 5;
  int y = 6;
  float d = floater(999);
  printf("d = %f\n", d);
}
