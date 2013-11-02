# Create the powerset of a set

def find_sets(s, container):
  if s not in container:
    container.append(s)
  if len(s) == 0:
    return set([])
  for e in s:
    copy = s.copy()
    copy.remove(e)
    find_sets(copy, container)
  return

def powerset(s):
  container = []
  find_sets(s, container)
  return container

resultant_sets = sorted(powerset(set(['a', 'b', 'c', 'd'])), key=len)
print("\n%s Sets\n" % len(resultant_sets))

for s in resultant_sets:
  if len(s) == 0:
    print("{}")
  else:
    print ("{%s}" % (','.join(map(str, s))))


resultant_sets = sorted(powerset(set([1,2,3,4])), key=len)
print("\n%s Sets\n" % len(resultant_sets))

for s in resultant_sets:
  if len(s) == 0:
    print("{}")
  else:
    print ("{%s}" % (','.join(map(str, s))))


def print_size_k(k, sets):
  for s in sets:
    if (len(s) == k):
      print ("{%s}" % (','.join(map(str, s))))

k = 2
print("Printing sets of size %d" % k)
print_size_k(k, resultant_sets)
