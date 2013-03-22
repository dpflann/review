#include <stdio.h>
#include <string.h>

void squeeze_delete(char s1[], char s2[]);

void squeeze_delete(char s1[], char s2[])
{
  int i, j, write_index;
  write_index = 0;
  for (i = 0; s1[i] != '\0'; i++)
  {
    int write;
    write = 1;
    for (j = 0; s2[j] != '\0'; j++)
    {
      // iterate through, checking s1[i] against elements of s2
      // if s1[i] occurs in s2[j], do not write
      // if s1[i] is not in s2[j], finish checking, then write s1[i]
      if (s1[i] == s2[j])
        write = 0;
    }
    if (write)
    {
        s1[write_index] = s1[i];
        write_index++;
    }
  }
  s1[write_index] = '\0';
} // end squeeze_delete


main()
{
  char s[] = "abcde";
  char t[] = "acfghi";

  printf("len(s) = %ld\n", strlen(s));
  printf("s = %s, t = %s\n", s, t);
  squeeze_delete(s, t);
  printf("s = %s, t = %s\n", s, t);
  printf("len(s) = %ld\n", strlen(s));
} // end main
