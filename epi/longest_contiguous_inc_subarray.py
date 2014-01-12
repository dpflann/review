# Implement an algorithm that takes as input an array A of n elements,
# and returns the beginning and ending indices of a longest increasing,
# contiguous subarray of A.

def longest_contiguous_increasing_subarray(A):
  i, j = 0, 0
  longest = (i, j)
  for n in range(0, len(A) - 1):
    if A[n] < A[n + 1]:
      j = n + 1
      longest = (i, j)
    else:
      i = j = n + 1
  return longest


lcis = longest_contiguous_increasing_subarray

def test_lcis():
  assert(lcis([ 1, 2, 3 ]) == (0, 2))
  assert(lcis([ 3, 2, 1 ]) == (0, 0))
  assert(lcis([ 1, 2, 3, 1, 2, 1, 5, 6, 7, 8 ]) == (5, 9))
  print "LCIS passes"

test_lcis()

