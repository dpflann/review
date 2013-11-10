# multiply two unsigned integers using only:
# assignment and bitwise operators

def simple_multiply(x, y):
  "This functions multiplies two unsigned integers using only assignment and bitwise operators."
  value = 0
  while (y):
    if (y & 0x1) == 0x1:
      value |= x
    x <<= 1
    y >>= 1
  return value


def test(function):
  print("Testing %s" % (function.func_name))
  assert(function(5, 6) == 30)
  assert(function(6, 1) == 6)
  assert(function(17, 2) == 34)
  print("Test pass")

test(simple_multiply)
