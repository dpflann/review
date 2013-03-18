#include <stdio.h>


int power(int base, int exponent); // function prototype, parameter names are optional in the prototype

int power(int base, int exponent) // base and exponent are parameters or 'formal arguments'
{
  // arguments are passed 'by value' as temporary variables local to the function
  // a calling function cannot directly alter a variable
  // unless passed a reference/address aka pointer with an appropriately delcard parameter
  int i, p;
  p = 1;
  for (i = 1; i <= exponent; i++)
  {
    p = p * base;
  }

  return p;
}

int power2(int base, int exponent) // this function uses call 'by value' to reduce the number of local variables used
{
  int p;
  for (p = 1; exponent > 0; exponent--)
    p = p * base;
  return p;
}

/* test power function */
main()
{

  int i;
  for (i = 0; i < 10; ++i)
    printf("2^%d = %d, 3^%d = %d\n", i, power(2, i), i, power2(3, i)); //2,3, and i are arguments or 'actual arguments'

  return 0; // main can return to the program that called it, which is essentially the environment in which the program was executed, can return status to said environment
}//end main
