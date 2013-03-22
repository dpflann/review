#include <stdio.h>

/* get n bits from position p */
unsigned getbits(unsigned x, int p, int n);
/* set n bits from position p in x equalt to rightmost n bits of y */
unsigned setbits(unsigned x, int p, int n, unsigned int y);

unsigned getbits(unsigned x, int p, int n)
{
  return (x >> (p + 1 - n)) & (~(~0 << n));
} // end getbits


unsigned setbits(unsigned x, int p, int n, unsigned int y)
{
  return (x | ((y & (~(~0 << n))) << (p - n + 1)));
} // end setbits

main()
{
  unsigned int x = 0x567;
  printf("getbits(%x, 3, 4) = %x\n", x, getbits(x, 3, 4));
  printf("1 = %x, 0 = %x, ~0 = %x, (~0 << 4) = %x\n", 1, 0, ~0, ~0 << 4);
  unsigned int y = setbits(x, 7, 4, 0xff);
  printf("setbits(%x, 7, 4, %x) = %x\n", x, 0xff, y);
}
