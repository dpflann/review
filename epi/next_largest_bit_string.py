## Algorithm for creating the next largets bit string


def next_largest_bit_string(n):
  n_str = [ s for s in n ][2:]
  i = len(n_str) - 1
  while n_str[i] == "1":
    n_str[i] = "0"
    i -= 1
  n_str[i] = "1"
  bin_array = ['0','b'] + n_str
  return ''.join(bin_array)

def n_l_b_s(n):
  i = 1
  while n & i >= 1:
    n &= ~i
    i <<= 1
  n |= i
  return n


a = bin(123)
print(a)
print(next_largest_bit_string(a))
print(123)
print(n_l_b_s(123))

a = bin(13)
print(a)
print(next_largest_bit_string(a))
print(13)
print(n_l_b_s(13))
