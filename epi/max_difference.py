# n 3-D coordinates (x, y, z)
# Ascent - drains battery
# Descent - charges battery
import sys

def energy_consumption(coordinates):
  zs = [ z for (x, y, z) in coordinates ]
  _min = sys.maxint
  capacity = 0
  for z in zs:
    capacity = max(capacity, z - _min)
    _min = min(_min, z)

  return capacity


def energy_consumption_2(coordinates):
  zs = [ z for (x, y, z) in coordinates ]
  _max = (-sys.maxint - 1)
  _min = sys.maxint
  difference = 0
  for z in zs:
    if z >= _max:
      _max = z
    if z <= _min:
      _min = z
    difference = _max - _min

  return difference

def test_energy_consumption():
  coords = [ (1, 2, 3), (4, 5, 6), (1, 4, 2), (5, 8, 9) ]
  assert(energy_consumption(coords) == 7)
  coords.append((1, 7, 25))
  assert(energy_consumption(coords) == 23)
  print("(1) Test Pass")

test_energy_consumption()

def test_energy_consumption_2():
  coords = [ (1, 2, 3), (4, 5, 6), (1, 4, 2), (5, 8, 9) ]
  assert(energy_consumption_2(coords) == 7)
  coords.append((1, 7, 25))
  assert(energy_consumption_2(coords) == 23)
  print("(2) Test Pass")

test_energy_consumption_2()
