#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>
#include "word_count.h"
#include "ch.c"

#define MAXWORD 100

struct tnode *addtree(struct tnode *p, char *word)
{
  int cond;
  if (p == NULL)
  {
    p = talloc();
    p->word = (char *) str_dup(word);
    p->count = 1;
    p->left = p->right = NULL;
  }
  else if ((cond = strcmp(word, p->word)) == 0)
    p->count++;
  else if (cond < 0)
    p->left = addtree(p->left, word);
  else
    p->right = addtree(p->right, word);

  return p;
}


void treeprint(struct tnode *p)
{
  if (p != NULL)
  {
    treeprint(p->left);
    printf("%4d %s\n", p->count, p->word);
    treeprint(p->right);
  }
}

struct tnode *talloc(void)
{
  return (struct tnode *) malloc(sizeof(struct tnode));
}

char *str_dup(char *s)
{
  char *p;
  
  p = (char *) malloc(strlen(s) + 1);
  if (p != NULL)
    strcpy(p, s);

  return p;
}

int get_word(char *word, int lim)
{
  int c, getch(void);
  void ungetch(int);
  char *w = word;
  
  while (isspace(c = getch()))
    ;
  if (c != '-')
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

/* word frequency count */
int main()
{
  struct tnode *root;
  char word[MAXWORD];
  
  root = NULL;
  while (get_word(word, MAXWORD) != '-')
    if (isalpha(word[0]))
        root = addtree(root, word);
  treeprint(root);

  return 0;
}// end main
