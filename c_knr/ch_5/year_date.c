#include <stdio.h>

void month_day(int year, int yearday, int *pmonth, int *pday);
int year_day(int year, int month, int day);
char *month_name(int n);

static char daytab[2][13] = {
  {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}, 
  {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}};


int year_day(int year, int month, int day)
{
  int i, leap;

  leap = year % 4 == 0 && year % 100 != 0 || year % 400 == 0;
  
  if (leap && month == 2 && day > 29)
    return;
  if (day > 31)
    return;
  if (month < 1 && month > 12)
    return;
  if (daytab[leap][month] < day)
    return;

  for (i = 1; i < month; i++)
    day += *((*(daytab + leap)) + i);
  return day;
}

void month_day(int year, int yearday, int *pmonth, int *pday)
{
  int i, leap;
  leap = year % 4 == 0 && year % 100 != 0 || year % 400 == 0;

  if (yearday > 366)
    return;
  if (yearday > 365 && !leap)
    return;

  for (i = 1;  yearday > *(*(daytab + leap) + i); i++)
    yearday -= *(*(daytab + leap) + i);
  *pmonth = i;
  *pday = yearday;
}

char *month_name(int n)
{
  static char *name[] = {
    "Illegal", 
    "January", "February", "March",
    "April", "May", "June",
    "July", "August", "September",
    "October", "November", "December" };

  return (n < 1 || n > 12) ? name[0] : name[n];
}

main()
{
  printf("Year day for 1/1/2000: %d\n", year_day(2000, 1, 1));
  printf("Year day for 12/31/2001: %d\n", year_day(2001, 12, 31));
  printf("Year day for 11/14/2004: %d\n", year_day(2004, 11, 14));
  printf("Year day for 12/31/2004: %d\n", year_day(2004, 12, 31));

  int day, month;
  month_day(2004, 365, &month, &day);
  printf("month_day(2004, 365) --> %d / %d\n", month, day);
  printf("month %d --> %s\n", month, month_name(month));

}// end main
