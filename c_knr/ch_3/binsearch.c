#include <stdio.h>

/* binary search for target x in array of numbers v of size n */
//int binary_search(int x, int v[], int n);
int binary_search(int x, int v[], int n)
{
  int low, mid, high;

  low = 0;
  high = n - 1;
  while (low <= high)
  {
    mid = (low + high) / 2;
    if (x < v[mid])
      high = mid - 1;
    else if (x > v[mid])
      low = mid + 1;
    else
      return mid;
  }
  return -1;
} // end binary_search

main()
{
  int v[5] = { 1, 2, 3, 4, 5 };
  printf("b-search for 1 in {1, 2, 3, 4, 5}: %d\n", binary_search(1, v, 5));
  printf("b-search for 6 in {1, 2, 3, 4, 5}: %d\n", binary_search(6, v, 5));
  printf("b-search for 4 in {1, 2, 3, 4, 5}: %d\n", binary_search(4, v, 5));
} // end main
