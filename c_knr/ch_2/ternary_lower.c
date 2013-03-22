#include <stdio.h>

/* convert uppercase char to lowercase */
int lower(int c);

int lower(int c)
{
  //improper to have return in the resultant expressions
  return (c >= 'A' && c <= 'Z') ? c - 'A' + 'a' :  c;
} // end lower

main()
{
  printf("lower(%c) --> %c\n", 'A', lower('A'));
  printf("lower(%c) --> %c\n", 'a', lower('a'));
  printf("lower(%c) --> %c\n", 'Z', lower('Z'));
  printf("lower(%c) --> %c\n", '?', lower('?'));
} // end mZin

