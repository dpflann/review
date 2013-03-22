#include <stdio.h>

int atoi(char s[]);
int lower(int c);
int htoi(char s[]);

int atoi(char s[])
{
  int i, number;
  number = 0;

  for (i = 0; s[i] != '\0' && s[i] >= '0' && s[i] <= '9'; i++)
    number = 10 * number + s[i] - '0';
  return number;
} // end atoi

int lower(int c)
{
  if (c >= 'A' && c <= 'Z')
    return c - 'A' + 'a';
  else
    return c;
} // end lower

int htoi(char s[])
{
  int i, number, error;
  number = error = 0;

  if (s[0] != '0' && (s[1] != 'x' || s[1] != 'X'))
      error = 1;
  else
  {
    for (i = 2; s[i] != '\0'; i++)
    {
      if (s[i] >= '0' && s[i] <= '9')
        number = 16 * number + s[i] - '0';
      else if (s[i] >= 'a' && s[i] <= 'f')
        number = 16 * number + s[i] - 'a' + 10;
      else if (s[i] >= 'A' && s[i] <= 'F')
        number = 16 * number + s[i] - 'A' + 10;
      else
      {
        break;
        error = 1;
        number = 0;
      }
    }
  }

  if (error)
  {
    printf("Input is not in correct format\n");
  }
  return number;
} // end htoi
main()
{
  printf("atoi(\"1234\") --> %d\n", atoi("1234"));
  printf("lower(\'D\') --> %c\n", lower('D'));
  printf("htoi(\"0x1A\") --> %d\n", htoi("0x1A"));
} // end main
