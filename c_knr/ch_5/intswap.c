#include <stdio.h>

void int_swap(int *a, int *b)
{
  int temp = *a;
  *a = *b;
  *b = temp;
}

main()
{
  int x = 1000;
  int y = -99;
  printf("x: %d, y: %d\n", x, y);
  int_swap(&x, &y);
  printf("x: %d, y: %d\n", x, y);
  

}
