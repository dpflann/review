# Given A (unsorted array of ints) and i (index into array)
# Partition the array into [ elements less than A, elements equal to A, elements greater than A ]

def partition(A, i):
  pivot = A[i]
  # move all pivots to front
  i = 0
  j = len(A) - 1
  while (i < j):
    if A[j] == pivot:
      temp = A[i]
      A[i] = A[j]
      A[j] = temp
      i += 1
    j -= 1
  last_pivot = i - 1
  # move all values less than pivot to front
  j = len(A) - 1
  while (i < j):
    if (A[j] < pivot):
      temp = A[i]
      A[i] = A[j]
      A[j] = temp
      i += 1
    j -= 1
  last_less_than = i
  # move pivots to middle
  for i in range(0, last_pivot + 1):
    temp = A[i]
    A[i] = A[last_less_than]
    A[last_less_than] = temp
    last_less_than -= 1

  return A


def swap(a, b):
  temp = a
  a = b
  b = temp


A = [1, 2, 3, 4, 5]
print(partition(A, 2))
A = [5, 6, 3, 1, 2, 3, 3]
print(partition(A, 2))
