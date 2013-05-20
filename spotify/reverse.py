### Solution to Spotify Puzzle: Problem ID: reversebinary
### Source: https://www.spotify.com/us/jobs/tech/reversed-binary/
### Author = DANIEL FLANNERY

import sys

def reverse(n):
  x = 0
  while (n):
    x = (x << 1) | (n & 1)
    n = n >> 1
  return x

### Assuming much about what is being read from stdin
n = int(sys.stdin.readline())
if (n >= 1 and n <= 1000000000):
  print (reverse(n))



