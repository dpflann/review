# refactored, was shift 1 then anding with integer, then shifting back to and with one to determine the value
# used local variables, but not necessary

def swap_bits(integer, i, j):
  if ((integer >> i) & 0x1) != ((integer >> j) & 0x1):
    integer = integer ^ (0x1 << i | 0x1 << j)

  return integer

print "0x9, swap at 3 and 0: ", swap_bits(0x9, 3, 0)
print "0xA, swap at 3 and 0: ", swap_bits(0xA, 3, 0)

