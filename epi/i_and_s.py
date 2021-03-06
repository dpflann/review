def itoa(n):
  s = ""
  sign = False

  if (n < 0):
    sign = True
    n = n * -1
  if (n == 0):
    return "0"

  while (n):
    v = n % 10
    s = str(v) + s
    n = n / 10
  if (sign):
    s = '-' + s
  return s



def test_itoa():
  assert(itoa(567) == '567')
  print(itoa(567))
  print(itoa(-567))
  assert(itoa(-567) == '-567')
  assert(itoa(0) == "0")
  print(itoa(0))

test_itoa()

def atoi(a):
  n = 0
  negative = False
  i = 0

  if (a[0] == '-'):
    negative = True
    i = 1

  while (i < len(a)):
    n = n * 10 + int(a[i])
    i += 1

  if (negative):
    n = n * -1

  return n

def test_atoi():
  assert(atoi("567") == 567)
  print(atoi("567"))
  assert(atoi("-567") == -567)
  print(atoi("-567"))
  assert(atoi("0") == 0)
  print(atoi("0"))

test_atoi()
