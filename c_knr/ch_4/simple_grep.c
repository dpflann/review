#include <stdio.h>
#include <string.h>
#define MAXLINE 1000

int get_line(char line[], int max);
int strindex(char source[], char searchfor[]);
int strrindex(char s[], char t[]);

char pattern[] = "ould"; /* pattern to search for */

/* fina all lines matching pattern */
main()
{
  char line[MAXLINE];
  int found = 0;

  while (get_line(line, MAXLINE) > 0)
    if (strindex(line, pattern) >= 0)
    {
      printf("%s", line);
      found++;
    }
  return found;
}// end main

/* get_line: get line into s, return length */
int get_line(char s[], int lim)
{
  int c, i;
  i = 0;
  while (--lim > 0 && (c = getchar()) != EOF && c != '\n')
    s[i++] = c;
  if (c == '\n')
    s[i++] = '\n';
  s[i++] = '\0';
  return i;
}// end get_line

/* strindex: return index of t in s, -1 if none */
int strindex(char s[], char t[])
{
  if (strlen(s) < strlen(t))
    return -1;

  int i, j, k;
  for (i = 0; s[i] != '\0'; i++)
  {
    for (j = i, k = 0; t[k] != '\0' && s[j] == t[k]; j++, k++)
      ;
    if (k > 0 && t[k] == '\0')
    {
      printf("i: %d\n", i);      
      return i;
    }
  }
  return -1;
} // end strindex

/* strrindex: return right most index of t in s, -1 if none */
int strrindex(char s[], char t[])
{
  if (strlen(s) < strlen(t))
    return -1;

  int i, j, k;
  for (i = strlen(s) - 1; i >= 0; i--)
  {
    for (j = i, k = strlen(t) - 1; k >= 0 && s >= 0 && s[j] == t[k]; j--, k--)
      ;
    if (k < 0 && s[++j] == t[++k])
    {
      printf("i: %d\n", i);
      return i;
    }
  }
  return -1;
} // end strrindex
