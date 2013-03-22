#include <stdio.h>

/* invert n bits starting at position p, that is if bit = 1 --> 0, if bit = 0 --> 1 */
int invert_bits(int x, int p, int n);

int invert_bits(int x, int p, int n)
{
  return x ^ (~(~0 << n) << (p - n + 1));
} // end invert_bits

main()
{
  printf("invert(%x, p = 3, n = 2) = %x\n", 0x6, invert_bits(0x6, 3, 2));
} // end main
