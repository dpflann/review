# Generate Next Lexicographic Permutation

def next_lexicographic_permutation(array):
  j = len(array) - 2 # n - 1
  # find location of first pair where a[j] < a[j+1]
  while array[j] > array[j + 1]:
    j -= 1
  
  # find next smallest integer greater than a[j] to the right of a[j]
  k = len(array) - 1
  while array[j] > array[k]:
    k -= 1

  # swap a[j] and a[k]
  temp = array[j]
  array[j] = array[k]
  array[k] = temp

  # construct next smallest
  r = len(array) - 1 # n
  s = j + 1

  while r > s:
    temp = array[r]
    array[r] = array[s]
    array[s] = temp
    r -= 1
    s += 1

  return array


def permute(array):
  import math
  current = array
  for i in range(math.factorial(len(array))):
    print(current)
    current = next_lexicographic_permutation(current)

# feed results of nlp(arr[i]) to nlp(arr[i+1])
array = [1,2,3]
permute(array)



