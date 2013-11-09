# determine if two xy-aligned rectangles intersect
#   ----------
#   -    A   -
#   -    ************
#   -----*    B     *
#        ************
#
#

class Rectangle:
  height = 0
  width = 0
  x = 0
  y = 0
  def __init__(self, x, y, height, width):
    self.height = height
    self.width = width
    self.x = x
    self.y = y


def intersection(rectangle_1, rectangle_2):
    x_intersect = None
    intersect_width = -1
    if ((rectangle_1.x <= rectangle_2.x) and (rectangle_2.x <= (rectangle_1.x + rectangle_1.width))):
      x_intersect = rectangle_2.x
      intersect_width = (rectangle_1.x + rectangle_1.width - rectangle_2.x)
    elif ((rectangle_2.x <= rectangle_1.x) and (rectangle_1.x <= (rectangle_2.x + rectangle_2.width))):
      x_intersect = rectangle_1.x
      intersect_width = (rectangle_2.x + rectangle_2.width - rectangle_1.x)

    y_intersect = None
    intersect_height = -1
    if ((rectangle_1.y <= rectangle_2.y) and (rectangle_2.y <= (rectangle_1.y + rectangle_1.height))):
      y_intersect = rectangle_2.y
      intersect_height = (rectangle_1.y + rectangle_1.height - rectangle_2.y)
    elif ((rectangle_2.y <= rectangle_1.y) and (rectangle_1.y <= (rectangle_2.y + rectangle_2.height))):
      y_intersect = rectangle_1.y
      intersect_height = (rectangle_2.y + rectangle_1.height - rectangle_1.y)

    return (x_intersect, y_intersect, intersect_width, intersect_height)


def test_intersection():
  r1 = Rectangle(0, 0, 4, 4)
  r2 = Rectangle(1, 1, 4, 4)
  r3 = Rectangle(100, 100, 1, 1)
  assert(intersection(r1, r2) == (1, 1, 3, 3))
  assert(intersection(r1, r3) == (None, None, -1, -1))
  r4 = Rectangle(3, 0, 3, 3)
  r5 = Rectangle(0, 1, 5, 5)
  assert(intersection(r4, r5) == (3, 1, 2, 2))
  assert(intersection(r4, r3) == (None, None, -1, -1))
  print("Tests Pass")

test_intersection()
