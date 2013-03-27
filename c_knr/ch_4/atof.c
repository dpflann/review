#include <ctype.h>
#include <stdio.h>
#include <math.h>

/* atof: convert string s to double */
double atof(char s[])
{
  double val, power, exp, exp_sign;
  int i, sign;
  
  for (i = 0; isspace(s[i]); i++) /* skip whitespace */
    ;
  sign = (s[i] == '-') ? -1 : 1;
  if (s[i] == '+' || s[i] == '-')
    i++;
  for (val = 0.0; isdigit(s[i]); i++)
    val = 10.0 * val + (s[i] - '0');
  if (s[i] == '.')
    i++;
  for (power = 1.0; isdigit(s[i]); i++)
  {
    val = 10.0 * val + (s[i] - '0');
    power *= 10.0;
  }
  if (s[i] == 'e' || s[i] == 'E')
    i++;
  exp_sign = (s[i] == '-') ? -1 : 1;
  if (s[i] == '+' || s[i] == '-')
    i++;
  for (exp = 0.0; isdigit(s[i]); i++)
    exp = 10.0 * exp + (s[i] - '0');

  return sign * val / power * pow(10, exp_sign * exp);
}// end atof

main()
{
  printf("atof(\"-913.4\") = %.3f\n", atof("-913.4"));
  printf("atof(\"-913.4e-2\") = %.3f\n", atof("-913.4e-2"));
  printf("atof(\"-913.4E10\") = %.3f\n", atof("-913.4E10"));
}// end main

