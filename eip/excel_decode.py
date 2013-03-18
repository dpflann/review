import string
#decode Excel column to number: A, B, ... AA, AB, ... BB ..

def decode_excel(S):
  base_10 = 0
  char_to_number_dict = { letter : number for number, letter in enumerate(string.uppercase, 1) }
  for i in range(0, len(S)):
    #base_10 += char_to_number_dict[S[i]] * pow(26, len(S) - 1 - i)
    base_10 = base_10 * 26 + char_to_number_dict[S[i]]
  return base_10

print "A -> ", decode_excel("A")
print "Z -> ", decode_excel("Z")
print "AA -> ", decode_excel("AA")
print "ZZ -> ", decode_excel("ZZ")
print "AAA -> ", decode_excel("AAA")
