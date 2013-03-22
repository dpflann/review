# Euclidean Algorithm for gcd(a, b) two positive integers
# a = b*q + r, gcd(a,b) = gcd(b,r)

def euclidean_algorithm(a, b):
  r = 0
  while (b != 0):
    r = a % b
    a = b
    b = r
  return a

print "gcd(5, 7) --> ", euclidean_algorithm(5, 7)
print "gcd(5, 17) --> ", euclidean_algorithm(5, 17)
print "gcd(35, 7) --> ", euclidean_algorithm(35, 7)
print "gcd(334145, 71231231) --> ", euclidean_algorithm(334145, 71231231)
print "gcd(3333, 33) --> ", euclidean_algorithm(3333, 33)
