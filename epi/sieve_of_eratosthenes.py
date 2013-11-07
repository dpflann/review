# find prime from 1 to n
import copy
import sys

def sieve_of_eratosthenes(n):
  numbers = [ i for i in range(2, n + 1) ]
  i = 0
  while (i < len(numbers)):
    d = numbers[i]
    clone = copy.copy(numbers)
    clone.remove(d)
    for n in clone:
      if n % d == 0:
        numbers.remove(n)
    i += 1
  return numbers


print("n = 10")
print(sieve_of_eratosthenes(10))

if (len(sys.argv) == 2):
  try:
    n = int(sys.argv[1])
    print("Using passed argument %d" % n)
    print(sieve_of_eratosthenes(n))
  except:
    print("Error: Improper arguments passed");

