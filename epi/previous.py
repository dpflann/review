def previously_occurred(sequence):
  for i in range(0, len(sequence)):
    for j in range(0, i):
      if sequence[i] == sequence[j]:
        return sequence[i]
  return None


print "[1,2,3,4,5,6,13,4,5] --> ", previously_occurred([1,2,3,4,5,6,13,4,5])

