# x is an element of S_k, where elements of S_k have k 1's in binary expansion
# Find y, where y is an element of S_k, k = x's k, where | y - x | is the minimum
# x,y : 64b unsigned int
def minimum_weight(x):
  i = 0
  n = x
  while (n):
    if (n ^ 0x3 == 0 or n ^ 0x3 == 0x3):
      i += 1
      n >>= 1
    else:
      break
  return x ^ (0x3 << i)

def min_weight(x):
  for i in range(0, 63):
    if (((x >> i) & 1) ^ ((x >> (i + 1)) & 1)):
      x ^= (1 << i) | (1 << (i + 1))
      return x

def test():
  assert(minimum_weight(5) == 6)
  assert(minimum_weight(14) == 13)
  assert(minimum_weight(22) == 21)
print("Passed")

def test2():
  assert(min_weight(5) == 6)
  assert(min_weight(14) == 13)
  assert(min_weight(22) == 21)
print("Passed")

test()
test2()
