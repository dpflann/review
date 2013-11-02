import math

# convert string-represented number s in base b1 into a string in base b2
def base_converter(b1, s, b2):
  base_10 = 0
  char_to_number_dict = { "0" : 0, "1" : 1, "2" : 2, "3" : 3, "4" : 4, "5" : 5, "6" : 6, "7" : 7, "8" : 8, "9" : 9, "A" : 10, "B" : 11, "C" : 12, "D" :13 , "E" : 14, "F" : 15 }
  number_to_char_dict = {char_to_number_dict[k]: k for k in char_to_number_dict}
  for i in range(0, len(s)):
    base_10 += char_to_number_dict[s[i]] * pow(b1, len(s) - 1 - i)

  remainders = []
  D = base_10
  while (D / b2):
    r = D % b2
    D = D / b2
    remainders.append(r)
  remainders.append(D % b2)
  t = ''
  for i in range(0, len(remainders)):
    t = number_to_char_dict[remainders[i]] + t
  print t


print """base_converter(16, "AB", 4)"""
base_converter(16, "AB", 4)
print """base_converter(2, "AB", 16)"""
base_converter(2, "10101011", 16)
print """base_converter(10, "32", 2)"""
base_converter(10, "32", 2)

