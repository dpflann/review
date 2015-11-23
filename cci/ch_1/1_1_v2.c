#include <stdio.h>

int unique_composition(char* s) {
  // When working with ascii char set
  char ascii[256] = {};
  while (*s) {
    char c = *s;
    if (ascii[c] == 1) {
      return 0;
    }
    else {
      ascii[c] = 1;
    }
    s++;
  }
  return 1;
}

int main() {
  char* s_1 = "hello";
  printf(unique_composition(s_1) ? "'%s' is unique\n" : "'%s' is not unique\n", s_1);
  char* s_2 = "abcdefg";
  printf(unique_composition(s_2) ? "'%s' is unique\n" : "'%s' is not unique\n", s_2);
}
