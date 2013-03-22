#include <stdio.h>

// filling char array s with chars between beginning and end, inclusive
int expansion(char s[], int beginning, int end, int j)
{
  int c;
  for (c = beginning; c <= end; c++, j++)
    s[j] = c;
  return j;
}

/* expansion of regEx-pattern like strings */
void expand(char s1[], char s2[])
{
  int i, j;
  i = j = 0;

  while (s1[i] != '\0')
  {
    if (s1[i] >= 'a' && s1[i] <= 'z' && s1[i + 1] == '-')
    {
      j = expansion(s2, s1[i], s1[i + 2], j);
      i += 3;
    }
    else if (s1[i] >= '0' && s1[i] <= '9' && s1[i + 1] == '-')
    {
      j = expansion(s2, s1[i], s1[i + 2], j);
      i += 3;
    }
    else
    {
      s2[j] = s1[i];
      i++, j++;
    }
  } // end while
  s2[j] = '\0';
} // end expand

main()
{
  char t[50];
  expand("bbbbbbba-zccccc0-4123", t);
  printf("expand(%s) = %s\n", "bbbbbbba-zccccc0-4123", t);
}// end main
