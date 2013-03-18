#include <stdio.h>

/* histogram for word lengths from input */

#define SIZE 26

main()
{
  int alphabet[SIZE], i, c; // for alphabet: index 0 = a ... index 25 = z, just lowercase
  for (i = 0; i < SIZE; i++)
    alphabet[i] = 0;

  while ((c = getchar()) != EOF) // using command line for example purposes to stop program
  {
    if (c == '\\')
      break;
    switch(c)
    {
      case 'a':
        alphabet[c - 'a']++;
        break;
      case 'b':
        alphabet[c - 'a']++;
        break;
      case 'c':
        alphabet[c - 'a']++;
        break;
      case 'd':
        alphabet[c - 'a']++;
        break;
      case 'e':
        alphabet[c - 'a']++;
        break;
      case 'f':
        alphabet[c - 'a']++;
        break;
      case 'g':
        alphabet[c - 'a']++;
        break;
      case 'h':
        alphabet[c - 'a']++;
        break;
      case 'i':
        alphabet[c - 'a']++;
        break;
      case 'j':
        alphabet[c - 'a']++;
        break;
      case 'k':
        alphabet[c - 'a']++;
        break;
      case 'l':
        alphabet[c - 'a']++;
        break;
      case 'm':
        alphabet[c - 'a']++;
        break;
      case 'n':
        alphabet[c - 'a']++;
        break;
      case 'o':
        alphabet[c - 'a']++;
        break;
      case 'p':
        alphabet[c - 'a']++;
        break;
      case 'q':
        alphabet[c - 'a']++;
        break;
      case 'r':
        alphabet[c - 'a']++;
        break;
      case 's':
        alphabet[c - 'a']++;
        break;
      case 't':
        alphabet[c - 'a']++;
        break;
      case 'u':
        alphabet[c - 'a']++;
        break;
      case 'v':
        alphabet[c - 'a']++;
        break;
      case 'w':
        alphabet[c - 'a']++;
        break;
      case 'x':
        alphabet[c - 'a']++;
        break;
      case 'y':
        alphabet[c - 'a']++;
        break;
      case 'z':
        alphabet[c - 'a']++;
        break;
    }
  }

  for (i = 0; i < SIZE; i++)
  {
    printf("%c: %d\n", 'a' + i, alphabet[i]);
  }
}// end main


