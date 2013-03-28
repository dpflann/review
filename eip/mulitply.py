import math

def multiply(x, y):
  if y == 0 or x == 0:
    return 0
  elif y == 1:
    return x
  elif y % 2 == 0:
    return 2 * multiply(x, y / 2)
  else:
    return 2 * multiply(x, math.floor(y / 2)) + x

print "m(3, 4) = ", multiply(3,  4)
print "m(3, 5) = ", multiply(3,  5)
print "m(3, 0) = ", multiply(3,  0)
print "m(3, 101) = ", multiply(3,  101)
print "m(0, 4) = ", multiply(0,  4)

