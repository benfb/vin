# vin
![Build Status](https://img.shields.io/travis/benfb/vin.svg?style=flat-square)
[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/benfb/vin)

play ball!

Installing or Updating
----------------------
It's pretty easy once you have [Go installed](https://golang.org/dl/):

```
go get -u github.com/benfb/vin
go install github.com/benfb/vin
```

Client/Server Model
-------------------
Vin has the ability to notify you via text message when a blacked-out game becomes available to watch on MLB.tv. To do this, it needs to periodically check with the MLB servers to see whether a game is finished. Most vin commands can be run locally, but to get notified, you have to run a server somewhere with `vin server` and then send a request to be notified from a computer with `vin watch`. You'll get a text 90 minutes after the game finishes (when you can watch it).

Examples
--------
### Results
```
$ vin results

Dodgers (69 - 55) @ Reds (53 - 71)
+---------+------+------+------+
|  TEAM   | RUNS | HITS | ERRS |
+---------+------+------+------+
| Dodgers |   18 |   21 |    1 |
| Reds    |    9 |   14 |    0 |
+---------+------+------+------+
Inning: Final ✓


Nationals (73 - 50) @ Orioles (67 - 56)
+-----------+------+------+------+
|   TEAM    | RUNS | HITS | ERRS |
+-----------+------+------+------+
| Nationals |    0 |    0 |    0 |
| Orioles   |    0 |    0 |    0 |
+-----------+------+------+------+
Inning: 1 ▾

Probable pitchers:
+-------------+-----+------+
|   PITCHER   | WON | LOST |
+-------------+-----+------+
| A.J. Cole   |   0 |    0 |
| Dylan Bundy |   6 |    4 |
+-------------+-----+------+
```

### Standings
```
$ vin standings nlc

National League Central
+---+---------------------+------+-----+------+------+-----+
| # |        TEAM         | PCT  | WON | LOST | BACK | STR |
+---+---------------------+------+-----+------+------+-----+
| 1 | Chicago Cubs        | .634 |  78 |   45 |    - | L1  |
| 2 | St. Louis Cardinals | .537 |  66 |   57 |   12 | W1  |
| 3 | Pittsburgh Pirates  | .512 |  62 |   59 |   15 | L3  |
| 4 | Milwaukee Brewers   | .431 |  53 |   70 |   25 | W1  |
| 5 | Cincinnati Reds     | .427 |  53 |   71 | 25.5 | L2  |
+---+---------------------+------+-----+------+------+-----+

$ vin standings -a

+----+-----------------------+------+-----+------+-----+
| #  |         TEAM          | PCT  | WON | LOST | STR |
+----+-----------------------+------+-----+------+-----+
|  1 | Chicago Cubs          | .634 |  78 |   45 | L1  |
|  2 | Washington Nationals  | .593 |  73 |   50 | L1  |
|  3 | Texas Rangers         | .584 |  73 |   52 | L2  |
|  4 | Cleveland Indians     | .582 |  71 |   51 | W1  |
|  5 | Toronto Blue Jays     | .565 |  70 |   54 | L1  |
...
| 28 | Arizona Diamondbacks  | .411 |  51 |   73 | L1  |
| 29 | Minnesota Twins       | .395 |  49 |   75 | L4  |
| 30 | Atlanta Braves        | .363 |  45 |   79 | W1  |
+----+-----------------------+------+-----+------+-----+
```

###
Namesake
--------
`vin` is named after [Vin Scully](https://en.wikipedia.org/wiki/Vin_Scully).
