# Spiral index implementation

A simple programming exercise for the evening

## Problem statement

All numbers on an infinite grid that extends in all four directions can be
identified with a single number in the following manner:

| 36 | 35 | 34 | 33 | 32 | 31 | 30 |
| 37 | 16 | 15 | 14 | 13 | 12 | 29 |
| 38 | 17 | 4 | 3 | 2 | 11 | 28 |
| 39 | 18 | 5 | 0 | 1 | 10 | 27 |
| 40 | 19 | 6 | 7 | 8 | 9 | 26 |
| 41 | 20 | 21 | 22 | 23 | 24 | 25 |
| 42 | 43 | 44 | 44 | 45 | 46 | 47 |

Where 0 corresponds with teh coordinates `(0 , 0)`. 5 corresponds with the
coordinates `(-1, 0)` and 29 with `(3, 2)`.

Print out the spiral index of any pair of coordinates that are input by the
user.


## Usage

```
$ spiral_index -h
Usage of spiral_index:
  -x int
    	X coordinate (default 3)
  -y int
    	Y coordinate (default 2)
```


Example

```
$ spiral_index
Using coordinates x=3, y=2
29
```

## Complexity

Given coordinates (i, j), and the max(i, j), i.e the max value of a given pair
of coordinates referred to as n:

Time complexity: O(n)
Space complexity: O(1)
