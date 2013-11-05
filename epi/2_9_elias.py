# Create the Elias encoding for an array of positive integers [i, j, k ... ] --> STRING_ELIAS_ENCODING: array --> str
# Create an array of positive integers [ i, j, k ... ] from an Elias encoded string. str --> array

def elias_encode(numbers):
  s = ''
  for n in numbers:
    length = 0
    while (n):
      length += 1
      value = n & 0x1
      s = str(value) + s
      n = n >> 1
    s = reduce(lambda x, y: str(x) + str(y), (length - 1) * [0], '') + s
  return s


def test_elias_encode():
  assert(elias_encode([int('0b1', 2)]) == "1")
  assert(elias_encode([int('0b1', 2), int('0b101', 2), int('0b1111', 2)]) == "0001111001011")
  assert(elias_encode([int('0b11', 2)]) != "11")
  print("Tests Pass")

test_elias_encode()

def elias_decode(string):
  length = len(string)
  i = 0
  numbers = []
  zeroes = 0
  while (i < length):
    if (string[i] == '0'):
      zeroes += 1
      i += 1
    if (string[i] == '1'):
      num = string[i: i + zeroes + 1]
      numbers.append(int(num, 2))
      i = i + zeroes + 1
      zeroes = 0
  return numbers

def test_elias_decode():
  assert(elias_decode("1") == [1])
  assert(elias_decode("00111") == [7])
  assert(elias_decode("0011100101") == [7, 5])
  print("Tests Pass")

test_elias_decode()
