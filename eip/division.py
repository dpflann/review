# Let n and d be positive integers. How many positive integers not exceeding n are divisible by d?

import pkg_resources
pkg_resources.require("matplotlib")
from matplotlib.pylab import *

def determine_divisible(n, d):
  return [number for number in range(1, n + 1) if number % d == 0]

def test_n(n):
  print "N=",n
  for d in range(1, n + 1):
   print "\td=",d, " # divisible=",len(determine_divisible(n, d))

#for n in range(10, 50):
  #test_n(n)



matplotlib.plot([1,2,3])
show()
