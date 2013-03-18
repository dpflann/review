#include <stdio.h>
#define MAXLINE 1000 /* maximum input line size */

int get_line(char line[], int maxline);
void copy(char to[], char from[]);

/* print longest input line */
main()
{

  int len; // current line length;
  int max; // maximum length so far
  char line[MAXLINE]; // current line
  char longest[MAXLINE]; // longest line

  max = 0;
  while ((len = get_line(line, MAXLINE)) > 0)
  {
    // can print the line if greater than certain number of characters
    /*
    if (len > threshold)
      printf("%s\n", line);
    */
    if (len > max)
    {
      max = len;
      copy(longest, line);
    }
  }
  if (max > 0) // there was a line
    printf("%s\n", longest);
  return 0;
} // end main

/* get_line: read a line into s, return length */
int get_line(char s[], int limit)
{
  int c, i;

  for (i = 0; i < limit - 1 && (c = getchar()) != EOF && c != '\n'; ++i)
    s[i] = c;
  if (c == '\n')
  {
    s[i] = c;
    ++i;
  }
  s[i] = '\0';
  return i;
} // end get_line

/* copy: copy 'from' into 'to', assume enough space */
void copy(char from[], char to[])
{
  int i;

  i = 0;
  // '\0' indicates the termination of a character array: "bye" --> ['b', 'y', 'e', '\0']
  while ((to[i] = from[i]) != '\0')
    ++i;
} // end copy
