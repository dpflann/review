#include <stdio.h>

/* replace escape characters \t and \n with visible characters, \t --> '\\', 't' */
void escape(char s[], char t[])
{
  int i, j;
  i = j = 0;
  while (s[i] != '\0')
  {
    switch(s[i])
    {
      case '\t':
        t[j] = '\\';
        t[++j] = 't';
        break;
      case '\n':
        t[j] = '\\';
        t[++j] = 'n';
        break;
      default:
        t[j] = s[i];
        break;
    } // end switch
    i++;
    j++;
  } // end while
  t[j] = '\0';
}

main()
{
  char s[] = "abcd\tcde\n";
  char t[15];
  escape(s, t);
  printf("%s\nescaped:\n%s\n", s, t);
} // end main
