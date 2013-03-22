#include <stdio.h>

/* count the number of 1 bits in a given number */
int bitcount(int x);

int bitcount(int x)
{
  int count = 0;
  while (x)
  {
    count++;
    x &= (x - 1);
  }
  return count;
} // end bitcount


main()
{
  printf("bitcount(%x) = %d\n", 0x8, bitcount(0x8));
  printf("bitcount(%x) = %d\n", 0x7, bitcount(0x7));
  printf("bitcount(%x) = %d\n", 0xA, bitcount(0xA));
  printf("bitcount(%d) = %d\n", -7, bitcount(-7));
} // end main()
