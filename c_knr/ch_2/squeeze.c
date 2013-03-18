#include <stdio.h>

/* remove occurrences of char c in string s */
void squeeze(char s[],int c);

void squeeze(char s[], int c)
{
  int i, j;

  for (i = j = 0; s[i] != '\0'; i++)
    if (s[i] != c)
      s[j++] = s[i];
  s[j] = '\0';
} // end squeeze

main()
{
  char s[] = "hello";
  squeeze(s, 'l');
  printf("squeeze(\"s = hello\", \'l\') --> %s\n", s);

} // end main
