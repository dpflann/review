#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include "push_and_pop.c"
#include <math.h>

#define MAXOP 100 //max size of operand or operator
#define NUMBER '0' //signal that a number was found

enum special_ops { POW = 1, SIN, EXP };

int getch(void);
void ungetch(int);

/* getop: get next operator or number operand */
int getop(char s[])
{
  int i, c;

  while ((s[0] = c = getch()) == ' ' || c == '\t')
    ;
  s[1] = '\0';
  if (!isdigit(c) && c != '.' && c != '-' && !(c == 'p' || c == 'e' || c == 's'))
    return c; //not a number
  i = 0;
  if (c == 'p')
    return process_special_op(c, "pow");
  if (c == 'e')
    return process_special_op(c, "exp");
  if (c == 's')
    return process_special_op(c, "sin");
  if (c == '-')
  {
    int d = c;
    c = getch();
    if (!isdigit(c))
    {
      ungetch(c);
      return d;
    }
    s[++i] = c;
  }
  if (isdigit(c)) //integer part
    while (isdigit(s[++i] = c = getch()))
      ;
  if (c == '.') //fraction part
    while (isdigit(s[++i] = c = getch()))
      ;

  s[i] = '\0';
  if (c != EOF)
    ungetch(c);
  printf("%s\n", s);
  return NUMBER;
}

#define BUFSIZE 100

char buf[BUFSIZE]; //buffer for ungetch
int bufp = 0; //next free position in buf

int getch(void)
{
  return (bufp > 0) ? buf[--bufp] : getchar();
}

void ungetch(int c)
{
  if (bufp >= BUFSIZE)
    printf("ungetch: too many characters\n");
  else
    buf[bufp++] = c;
}

int process_special_op(int c, char op_name[])
{
  int op_1 = c;
  if (c == op_name[0])
  {
    c = getch();
    if (c == op_name[1])
    {
      int op_2 = c;
      c = getch();
      if (c == op_name[2])
      {
        return op_constant(op_name);
      }
      else
      {
        ungetch(c);
        return op_2;
      }
    }
    else
    {
      ungetch(c);
      return op_1;
    }
  }
  return op_1;
}

int op_constant(char s[])
{
  if (s == "pow")
    return POW;
  if (s == "exp")
    return EXP;
  if (s == "sin")
    return SIN;
}

/* reverse polish calculator */
main()
{
  int type;
  double op2;
  char s[MAXOP];

  while ((type = getop(s)) != EOF)
  {
    switch (type)
    {
      case NUMBER:
        push(atof(s));
        break;
      case '+':
        push(pop() + pop());
        break;
      case '*':
        push(pop() * pop());
        break;
      /* - and / are not commutative like * and +, therefore must maintain strict order on
       * operands because the order of pop() in pop() op pop() evaluation is not defined */
      case '-':
        op2 = pop();
        push(pop() - op2);
        break;
      case '/':
        op2 = pop();
        if (op2 != 0.0)
          push(pop() / op2);
        else
          printf("ERROR: zero divisor\n");
        break;
      case '%':
        op2 = pop();
        push((int) pop() % (int) op2);
        break;
      case POW:
        op2 = pop();
        push(pow(pop(), op2));
        break;
      case SIN:
        push(sin(pop()));
        break;
      case EXP:
        push(exp(pop()));
        break;
      case '\n':
        printf("\t%.8g\n", pop());
        break;
      default:
        printf("ERROR: uknown command %s\n", s);
        break;
    }// end switch (type)
  }

  return 0;
}// end main

