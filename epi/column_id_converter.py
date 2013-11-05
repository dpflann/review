# convert the alpha characters in an Excel spreadsheet to their numerical value.
# e.g. 'A' = 1st column, 'AA' = 27th column

def column_id(string):
  n = 0
  for c in string:
    n = (n * 26) + (ord(c) - ord('A') + 1)
  return n


def test():
  assert(column_id('A') == 1)
  print('A = %d', column_id('A'))
  assert(column_id('AA') == 27)
  print('AA = %d', column_id('AA'))
  print("Pass")


test()
