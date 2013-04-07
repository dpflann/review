#include <stdio.h>

struct point
{
  int x;
  int y;
};

struct string
{
  int len;
  char *str;
};

struct point make_point(int x, int y)
{
  struct point temp;
  temp.x = x;
  temp.y = y;

  return temp;
}// end make_point

main()
{
  #define SIZE 5
  struct point point_1 = { 5, 10 };
  struct point *p = &point_1;
  struct point points[SIZE];
  struct point* pps = points;
  int i;

  struct {
    int len;
    char *str;
  } *a;

  struct string s = { 6, "hello" };

  struct
  {
    int len;
    char *str;
  } ss;

  ss.len = 6;
  ss.str = "hello";

  /* causes a misaligned address error, because a points to nothing */
  /*
  a->len = 6;
  a->str = "hello";
  */

  a = &ss; //warning incompatible pointer type

  printf("Point1: (x, y) --> (%d, %d)\n", point_1.x, point_1.y);
  printf("Point1: (x, y) --> (%d, %d)\n", p->x, p->y);

  printf("playing around with pointers and structs\n");

  for (i = 0; i < SIZE; i++)
  {
    points[i] = make_point(i, i + 10);
  }

  printf("(%d, %d)\n", (*pps).x, (*pps).y);
  printf("(%d, %d)\n", (*++pps).x, (*++pps).y);
  printf("(%d, %d)\n", ++(*pps).x, ++(*pps).y);

  printf("a points to \nstruct {\n\t int len;\n\tchar *str;\n\t} *a;\n");
  printf("a->len: %d, a->str: %s\n", (*a).len, (*a).str);

}//end main
