# Given a set of four points, determine if they form an xy-aligned rectangle

class Point:
  x = None
  y = None
  
  def __init__(self, x, y):
    self.x = x
    self.y = y


def is_xy_aligned(points):
  xs = set([])
  ys = set([])

  for p in points:
    xs.add(p.x)
    ys.add(p.y)

  return len(xs) == 2 and len(ys) == 2

def test_is_xy_aligned():
  p_1 = Point(0, 0)
  p_2 = Point(0, 4)
  p_3 = Point(4, 0)
  p_4 = Point(4, 4)
  p_5 = Point(100, 100)
  points_1 = [ p_1, p_2, p_3, p_4 ]
  points_2 = [ p_1, p_2, p_3, p_5 ]
  assert(is_xy_aligned(points_1) == True)
  assert(is_xy_aligned(points_2) == False)
  print("Test Pass")

test_is_xy_aligned();
