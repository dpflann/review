#include <stdio.h>
#define MAX 100

int vals[MAX];
int sp = -1;

int pop()
{
  if (sp < 0)
    printf("stack is empty\n");
  else
    return vals[sp--];
}

void push(int value)
{
  if (sp == MAX - 1)
    printf("stack is full\n");
  else
    vals[++sp] = value;
}

void operate(int op)
{
  int a = pop();
  int b = pop();
  switch (op)
  {
    case '+':
      push(a + b);
      printf("%d + %d = %d\n", a, b, a + b);
      break;
    case '-':
      push(a - b);
      printf("%d - %d = %d\n", a, b, a - b);
      break;
    case '*':
      push(a * b);
      printf("%d * %d = %d\n", a, b, a * b);
      break;
    case '/':
      push(a / b);
      printf("%d / %d = %d\n", a, b, a / b);
      break;
  }
}

main(int argc, char *argv[])
{
  while (--argc > 0)
  {
    char *s = *++argv;
    int value = 0;
    int valid = 1;
    int op = 0;
    while (*s)
    {
      if (!isdigit(*s) && !(*s == '+' || *s == '-' || *s == '/' || *s == '*'))
      {
        valid = 0;
        break;
      }
      if (isdigit(*s))
        value = value * 10 + (*s - '0');
      else
      {
        operate(*s);
        op = 1;
      }
      s++;
    }
    if (valid && !op)
      push(value);
    else if (!op)
    {
      printf("Invalid input %s\n", *argv);
      break;
    }
  }
}// end main
