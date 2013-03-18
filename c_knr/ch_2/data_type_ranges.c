#include <stdio.h>
#include <limits.h>

main()
{

  signed char s_c;
  signed short int s_shrt;
  signed long int s_lng;

  unsigned char us_c;
  unsigned short int us_shrt;
  unsigned long int us_lng;

  // signed --> [-2^(n-1), 2^(n-1) - 1]
  // unsigned --> [0, 2^(n) - 1]

  printf("Signed Char: [%d, %d]\n", SCHAR_MIN, SCHAR_MAX);
  printf("Unsigned Char: [%d, %d]\n", 0, UCHAR_MAX);

  printf("Signed Short: [%d, %d]\n", SHRT_MIN, SHRT_MAX);
  printf("Unsigned Short: [%d, %d]\n", 0, USHRT_MAX);

  printf("Signed Long: [%ld, %ld]\n", LONG_MIN, LONG_MAX);
  printf("Unsigned Long: [%lu, %lu]\n", 0, ULONG_MAX);

} // end main
