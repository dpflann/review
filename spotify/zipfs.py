### Solution to Spotify Puzzle: Problem ID: zipfsong
### Source: https://www.spotify.com/us/jobs/tech/zipfsong/
### Author = DANIEL FLANNERY

import sys
import io

def calculate_zi(i, base):
  return ((1.0 * base) / i)

def calculate_qi(fi, zi):
  return (fi / zi)

def process_input(album_file):
  n_and_m = album_file.readline().split(' ')
  n_and_m = [ int(x) for x in n_and_m ]
  processed_tracks = {}
  tracks = album_file.readlines()

  # First Track is the basis
  track = tracks[0].strip("\n")
  f_and_title = track.split(' ')
  fi = int(f_and_title[0])
  zi = fi
  zipf_base = fi
  qi = calculate_qi(fi, zi)

  #process remaning tracks
  for i in range(1, n_and_m[0]):
    track = tracks[i].strip("\n")
    f_and_title = track.split(' ')

    fi = int(f_and_title[0])
    zi = calculate_zi(i + 1, zipf_base)
    qi = calculate_qi(fi, zi)

    if qi in processed_tracks:
      processed_tracks[qi].append( f_and_title[1] )
    else:
      processed_tracks[qi] = [ f_and_title[1] ]

  ordered_keys = sorted(processed_tracks, reverse=True)
  for key in ordered_keys[ 0 : n_and_m[1] ]:
    for title in processed_tracks[key]:
      print title


process_input(sys.stdin)


