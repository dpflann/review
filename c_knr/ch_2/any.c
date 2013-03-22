#include <stdio.h>
#include <string.h>

int any(char s1[], char s2[]);

int any(char s1[], char s2[])
{
  int present, location, i, j;
  present = location = 0;

  for (i = 0; s1[i] != '\0'; i++)
  {
    if (present)
      break;
    for (j = 0; s2[j] != '\0'; j++)
    {
      if (s1[i] == s2[j])
      {
        location = i;
        present = 1;
        break;
      }
    }
  }

  if (present)
    return location;
  else
    return -1;
} // end any

main()
{
  char s[] = "abcdefg";
  char t[] = "fhi";
  char u[] = "zzzzzzz";

  printf("s = %s, t = %s\n", s, t);
  int s_in_t = any(s, t);
  (s_in_t > -1) ? printf("%c from %s in %s\n", s[s_in_t], s, t) : printf("No char from s is in t\n");


  printf("s = %s, u = %s\n", s, u);
  int s_in_u = any(s, u);
  (s_in_u > -1) ? printf("%c from %s in %s\n", s[s_in_u], s, u) : printf("No char from s is in t\n");


} // end main
