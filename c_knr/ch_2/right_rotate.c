#include <stdio.h>

/* rotate x by n, shifting n lsb bits to be the msb bits, and the msb bits to be the lsb bits: 
 * 1100 rotated right by 2 --> 0011
 * 0011 rotated right by 2 --> 1100
 */
int right_rotate(int x, int n);

int right_rotate(int x, int n)
{
  return (x >> n) | ((x & (~(~0 << n))) << n);
} // end right_rotate

main()
{
  printf("right_rotate(%x, 2) --> %x\n", 0xC, right_rotate(0xC, 2));
  printf("right_rotate(%x, 2) --> %x\n", 0x3, right_rotate(0x3, 2));
} // end main
