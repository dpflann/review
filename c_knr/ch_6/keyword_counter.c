#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include "ch.c"

#define MAXWORD 100
#define NKEYS (sizeof keytab / sizeof(struct key))

struct key {
  char *word;
  int count;
} keytab[] =
  { {"auto", 0},
    {"break", 0},
    {"case", 0},
    {"char", 0},
    /* ... */
    {"void", 0},
    {"volatile", 0},
    {"while", 0}
  };

int get_word(char *word, int);
int bin_search(char *word, struct key tab[], int n);


/* count C keywords */
int main()
{
  int n;
  char word[MAXWORD];

  while (get_word(word, MAXWORD) != 'c')
  {
    if (isalpha(word[0]))
      if ((n = bin_search(word, keytab, NKEYS)) >= 0)
        keytab[n].count++;
  }
  for (n = 0; n < NKEYS; n++)
    if (keytab[n].count > 0)
      printf("%4d %s\n", keytab[n].count, keytab[n].word);

  return 0;
}// end main

/* bin_search: binary search, find word in tab[0]..tab[n-1] */
int bin_search(char *word, struct key tab[], int n)
{
  int cond;
  int low, high, mid;

  low = 0;
  high = n - 1;
  while (low <= high)
  {
    mid = (low + high) / 2;
    if ((cond = strcmp(word, tab[mid].word)) < 0)
      high = mid - 1;
    else if (cond > 0)
      low = mid + 1;
    else
      return mid;
  }
  return -1;
}// end bin_search

/* getword: get next word or character from input */
int get_word(char *word, int lim)
{
  int c, getch(void);
  void ungetch(int);
  char *w = word;
  
  while (isspace(c = getch()))
    ;
  if (c != EOF)
    *w++ = c;
  if (!isalpha(c))
  {
    *w = '\0';
    return c;
  }
  for (; --lim > 0; w++)
    if (!isalnum(*w = getch()))
    {
      ungetch(*w);
      break;
    }
  *w = '\0';

  return word[0];
}// end get_word
