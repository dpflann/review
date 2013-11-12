# x,y are positive integer: perform x/y with (+, -, *)

def integer_division(x, y):
  d = 0
  while (x >= y):
    d += 1
    x -= y

  return d

def test_integer_division():
  assert(integer_division(6, 3) == 2)
  assert(integer_division(1, 5) == 0)
  assert(integer_division(17, 1) == 17)
  assert(integer_division(15, 5) == 3)
  print("Test Pass")

test_integer_division()
