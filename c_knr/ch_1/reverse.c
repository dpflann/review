#include <stdio.h>

/* pedantic revere string, using '\0' terminator concept to determine length */

void reverse(char str[]);

main()
{
  char s[] = "reverse this but not racecar";
  printf("%s\n", s);
  reverse(s);
  printf("%s\n", s);
}


/* function to reverse a string */
void reverse(char str[])
{
  int length, i, j;
  for (i = 0; str[i] != '\0'; i++);
  length = i;

  for (i = 0, j = length - 1; i < j; i++, j--)
  {
    char temp = str[i];
    str[i] = str[j];
    str[j] = temp;
  }
}

