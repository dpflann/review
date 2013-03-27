#include <stdio.h>

#define MAXVAL 100 //maximum depth of val stack

int sp = 0;
double stack[MAXVAL];

/* push: push f onto stack */
void push(double f)
{
  if (sp < MAXVAL)
    stack[sp++] = f;
  else
    printf("Stack is full\n");
}// end push

/* pop: pop f off stack */
double pop()
{
  if (sp > 0)
    return stack[--sp];
  else
  {
    printf("ERROR: stack empty\n");
    return 0.0;
  }
}

void print_top()
{
  if (sp > 0)
    printf("top: %f\n", stack[sp - 1]);
  else
    printf("ERROR: stack empty\n");
}

void duplicate_top()
{
  if (sp > 0)
  {
    double dupe = stack[sp - 1];
    push(dupe);
  }
  else
    printf("ERROR: stack empty\n");
}

void swap_top_two()
{
  if (sp > 1)
  {
    double top = pop();
    double bottom = pop();
    push(top);
    push(bottom);
  }
  else
    printf("ERROR: stack empty\n");
}

void clear_stack()
{
  sp = 0;
}
