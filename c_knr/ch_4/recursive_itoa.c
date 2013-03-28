#include <stdio.h>

int r_itoa(int i, char a[])
{
  if (i == 0)
  {
    a[0] = 0;
    return -1;
  }
  else
  {
    int digit = i % 10;
    int index = r_itoa(i / 10, a);
    if (index == -1)
      index++;
    a[index] = digit + '0';
    return index + 1;
  }
}

main()
{
  char a[100];
  a[r_itoa(12345, a)] = '\0';
  printf("r_itoa(12345, a) -> %s\n", a);
}
