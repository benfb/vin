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

Namesake
--------
`vin` is named after [Vin Scully](https://en.wikipedia.org/wiki/Vin_Scully).
