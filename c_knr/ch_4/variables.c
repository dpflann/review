#include <stdio.h>

#define MAXVAL 100 //maximum depth of val stack

int var;
double variables[26];

void reset_var()
{
  var = 0;
}

void set_var(double f)
{
  variables[var - 'a'] = f;
  reset_var();
}

double get_var(int c)
{
  if (c >= 'a' && c <= 'z')
  {
    if (variables[c - 'a'])
      return variables[c - 'a'];
  }
  else
    printf("Error: unknown variable\n");
}

void set_candidate_var(int c)
{
  var = c;
}

int get_candidate_var()
{
  return var;
}

int check_var(int c)
{
  if (c >= 'a' &&  c <= 'z')
    if (variables[c - 'a'])
      return variables[c - 'a'];
  else
    return 0;
}
