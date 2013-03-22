def parity(integer):
  count = 0
  while (integer > 0):
    if integer & 1 == 1:
      count += 1
    integer = integer >> 1
  return False if count % 2 == 0 else True

def parity_2(integer):
  result = 0
  while (integer):
    result = result ^ (integer & 1)
    integer = integer >> 1
  return True if result else False

print "1 --> ", parity(1)
print "45 --> ", parity(45)
print "10 --> ", parity(10)

print "1 --> ", parity_2(1)
print "45 --> ", parity_2(45)
print "10 --> ", parity_2(10)

