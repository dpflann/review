#gcd without division, multipliation, modulus

def gcd(x, y):
  if y == 0:
    return x
  else:
    r = x
    while (r > 0):
      r = r - y
    if (r < 0):
      r = r + y
    return gcd(y, r)

def test_gcd():
  assert(gcd(1, 5) == 1)
  assert(gcd(2, 5) == 1)
  assert(gcd(15, 5) == 5)
  assert(gcd(6, 12) == 6)
  print("Test Pass")

test_gcd()
