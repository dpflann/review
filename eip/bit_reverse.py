def reverse_bits(integer):
  result = 0
  i = 0
  while (integer >> i):
    result = (result << 1) | ((integer >> i) & 0x1)
    i += 1
  return result 

print "0x8 --> %X\n" % reverse_bits(0x8)
print "0xF --> %X\n" % reverse_bits(0xF)
print "0x90 --> %X\n" % reverse_bits(0x90)

