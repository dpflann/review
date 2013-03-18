#include <stdio.h>

/* Word count */

#define IN 1 /* inside a word */
#define OUT 0 /* outside a word */

/* count lines, words, and characters in input */

main()
{
  int c, nl, nw, nc, state;
  state = OUT;
  nl = nw = nc = 0; // assignment is an expression with a value, and assignments associate from right to left
  while ((c = getchar()) != EOF) {
    ++nc;
    if (c == '\n')
      ++nl;
    /* evaluation will stop as soon as the truth or falsehood is known, is c is ' ' then no need
     * to test if c == '\n', can be useful for more complex events
     * */
    if (c == ' ' || c == '\n' || c == '\t')
      state = OUT;
    else if (state = OUT) {
      state = IN;
      ++nw;
    }
  }
  printf("%d %d %d\n", nl, nw, nc);
}
