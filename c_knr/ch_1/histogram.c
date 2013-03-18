#include <stdio.h>

/* histogram for word lengths from input */

#define SIZE 25

main()
{
  int lengths[SIZE], length, i, c; // index 0 = length 1, index 1 = length 2, index 23 = length 24, index 24 = length > 24
  i = length = 0;

  for (i = 0; i < SIZE; i++)
    lengths[i] = 0;

  while ((c = getchar()) != EOF) // using command line for example purposes to stop program
  {
    if (c == '\\')
      break;
    // if char is whitespace, at word stopping place, reset length counter
    if (c == ' ' || c == '\n' || c == '\b' || c == '\t')
    {
      if (length > 0) {
        if (length > 24)
          lengths[24]++;
        else
          lengths[length - 1]++;
      }
      length = 0;
    }
    else {
      length++;
    }
  }

  for (i = 0; i < SIZE; i++) {
    int j = 0;
    if (i == 24)
      printf(">24:");
    else
      printf("%d:", i);
    for (j = 0; j < lengths[i]; j++)
      printf("-");
    printf("\n");
  }
}// end main

