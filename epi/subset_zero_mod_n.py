import itertools as it

def powerset(iterable):
   "powerset([1,2,3]) --> () (1,) (2,) (3,) (1,2) (1,3) (2,3) (1,2,3)"
   s = list(iterable)
   return it.chain.from_iterable(it.combinations(s, r) for r in range(len(s)+1))


def find_indicies(A):
  n = len(A)
  for i in range(0, n):
    mod_A = A[i] % n
    if mod_A == 0:
      return [i]
    else:
      A[i] = mod_A

  # find subsets of indicies
  subset_indicies = list(powerset(range(0, n)))
  for s in subset_indicies:
    if len(s) != 0:
      subset_sum = reduce(lambda x, y: x + y, [A[i] for i in s])
      if subset_sum % n == 0:
        return list(s)

  return []

def test_find_indicies():
  A = [1, 2, 3]
  assert(find_indicies(A) == [2])
  A = [ 429, 334, 62, 711, 704, 763, 98, 733, 721, 995 ]
  assert(find_indicies(A) == [ 0, 3 ])
  print("find_indicies passes")

test_find_indicies()

