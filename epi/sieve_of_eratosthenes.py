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

def alternate_sieve(n):
  numbers = []
  for i in range(2, n + 1):
    j = 0
    add = True
    while (j < len(numbers)):
      if i % numbers[j] == 0:
        add = False
        break
      j += 1
    if (add):
      numbers.append(i)
  return numbers


#### TESTING
print("n = 10")
print(sieve_of_eratosthenes(10))

def test_sieve():
  assert(sieve_of_eratosthenes(2) == [2])
  assert(sieve_of_eratosthenes(3) == [2,3])
  assert(sieve_of_eratosthenes(10) == [2,3,5,7])
  print("Tests pass")

def test_alternate_sieve():
  assert(alternate_sieve(2) == [2])
  assert(alternate_sieve(3) == [2,3])
  assert(alternate_sieve(10) == [2,3,5,7])
  print("Tests pass")

print("Testing Sieve of Eratosthenes")
test_sieve()
print("Testing Alternate Sieve")
test_alternate_sieve()

### If you want to use commandline arguments for the sieve
if (len(sys.argv) == 2):
  try:
    n = int(sys.argv[1])
    print("Using passed argument %d" % n)
    print(sieve_of_eratosthenes(n))
  except:
    print("Error: Improper arguments passed");

