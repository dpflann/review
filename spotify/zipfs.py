### Solution to Spotify Puzzle: Problem ID: zipfsong
### Source: https://www.spotify.com/us/jobs/tech/zipfsong/
### Author = DANIEL FLANNERY

import sys
import io

def calculate_zi(i):
  return (1.0 / i)

def calculate_qi(fi, zi):
  return (fi / zi)

def process_input(album_file):
  n_and_m = album_file.readline()
  n_and_m = n_and_m.split(' ')
  n_and_m = [ int(x) for x in n_and_m ]
  tracks = album_file.readlines()
  processed_tracks = {}

  for i in range(0, len(tracks)):
    track = tracks[i].strip("\n")
    f_and_title = track.split(' ')
    fi = int(f_and_title[0])
    zi = calculate_zi(i + 1)
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


