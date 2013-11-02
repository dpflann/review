def binary_to_octal(n):
  "return an octal string representation"
  result = ""
  while (n):
    result = str(n % 8) + result
    n = n / 8
  return result

print binary_to_octal.__doc__
print "binary_to_octal(128) --> ", binary_to_octal(128)
print "binary_to_octal(7) --> ", binary_to_octal(7)
print "binary_to_octal(0xFF) --> ", binary_to_octal(0xFF)

