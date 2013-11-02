# convert a number represented by s in base b1 into a number in base b2, return a string representation

def base_converter(b1, s1, b2):
  digits = [ str(n) for n in range(0,10) ] + ['A', 'B', 'C', 'D', 'E', 'F']
  # 2 <= b1, b1 <= 16
  n1 = 0
  for l in s1:
    d = int(l)
    n1 = n1 * b1 + d

  s2 = ''
  while (n1):
    r = n1 % b2
    s2 = digits[r] + s2
    n1 = n1 / b2

  return s2

def test_base_converter():
  assert(base_converter(2, "1101", 8) == "15")
  print(base_converter(2, "1101", 8))
  assert(base_converter(2, "1101", 4) == "31")
  print(base_converter(2, "1101", 4))
  assert(base_converter(2, "1101", 16) == "D")
  print(base_converter(2, "1101", 16))
  print("Pass")

test_base_converter()
