# Solver for the Unblock game on Android

## Usage

Note: if you're too lazy or don't have Go compiler, go to unblock.go code,
Ctrl-A + Ctrl-C to copy it all, then go to http://play.go.dev/ and
paste the code there.  Then Run it and you've got the result.

Edit the `unblock.go` and change the value of the initial state of the
desk.  Example:

```
	start := Desk{
		0x44, 0x46, 0x00,
		0x06, 0x06, 0x66,
		0x06, 0x44, 0x66,
		0x44, 0x60, 0x66,
		0x00, 0x64, 0x40,
		0x44, 0x40, 0x00,
	}
```

Each hex number contains 2 hex digits, each digit encodes the cell
state.  Thus there must be 6x6 field.

Encodings are:
*  0  the cell is empty
*  4,5 the cell is occupied by a horizontal plank.
   Use either 4 and 5 to distinguish between two planks in the same row,
   for example `0x04, 0x45, 0x50` means a space, then two planks, and
   a space again.
*  6,7 the cell is occupied by a vertical plank.

Then run the solver:

```
go run unblock.go
```

The typical output is:
```
2025/11/01 18:30:25 unblock.go:357: step #0 input desks 1
2025/11/01 18:30:25 unblock.go:389: step #0 has finished with 8 states (9 total)
2025/11/01 18:30:25 unblock.go:357: step #1 input desks 8
2025/11/01 18:30:25 unblock.go:389: step #1 has finished with 20 states (29 total)
2025/11/01 18:30:25 unblock.go:357: step #2 input desks 20
2025/11/01 18:30:25 unblock.go:389: step #2 has finished with 21 states (50 total)
2025/11/01 18:30:25 unblock.go:357: step #3 input desks 21
2025/11/01 18:30:25 unblock.go:389: step #3 has finished with 19 states (69 total)
2025/11/01 18:30:25 unblock.go:357: step #4 input desks 19
2025/11/01 18:30:25 unblock.go:389: step #4 has finished with 35 states (104 total)
2025/11/01 18:30:25 unblock.go:357: step #5 input desks 35
2025/11/01 18:30:25 unblock.go:389: step #5 has finished with 75 states (179 total)
2025/11/01 18:30:25 unblock.go:357: step #6 input desks 75
2025/11/01 18:30:25 unblock.go:389: step #6 has finished with 158 states (337 total)
2025/11/01 18:30:25 unblock.go:357: step #7 input desks 158
2025/11/01 18:30:25 unblock.go:389: step #7 has finished with 317 states (654 total)
2025/11/01 18:30:25 unblock.go:357: step #8 input desks 317
2025/11/01 18:30:25 unblock.go:389: step #8 has finished with 549 states (1203 total)
2025/11/01 18:30:25 unblock.go:357: step #9 input desks 549
2025/11/01 18:30:25 unblock.go:389: step #9 has finished with 767 states (1970 total)
2025/11/01 18:30:25 unblock.go:357: step #10 input desks 767
2025/11/01 18:30:25 unblock.go:389: step #10 has finished with 863 states (2833 total)
2025/11/01 18:30:25 unblock.go:357: step #11 input desks 863
2025/11/01 18:30:25 unblock.go:389: step #11 has finished with 630 states (3463 total)
2025/11/01 18:30:25 unblock.go:357: step #12 input desks 630
2025/11/01 18:30:25 unblock.go:389: step #12 has finished with 338 states (3801 total)
2025/11/01 18:30:25 unblock.go:357: step #13 input desks 338
2025/11/01 18:30:25 unblock.go:389: step #13 has finished with 163 states (3964 total)
2025/11/01 18:30:25 unblock.go:357: step #14 input desks 163
2025/11/01 18:30:25 unblock.go:389: step #14 has finished with 36 states (4000 total)
2025/11/01 18:30:25 unblock.go:357: step #15 input desks 36
2025/11/01 18:30:25 unblock.go:389: step #15 has finished with 13 states (4013 total)
2025/11/01 18:30:25 unblock.go:357: step #16 input desks 13
2025/11/01 18:30:25 unblock.go:389: step #16 has finished with 2 states (4015 total)
2025/11/01 18:30:25 unblock.go:357: step #17 input desks 2
2025/11/01 18:30:25 unblock.go:389: step #17 has finished with 3 states (4018 total)
2025/11/01 18:30:25 unblock.go:357: step #18 input desks 3
2025/11/01 18:30:25 unblock.go:389: step #18 has finished with 6 states (4024 total)
2025/11/01 18:30:25 unblock.go:357: step #19 input desks 6
2025/11/01 18:30:25 unblock.go:389: step #19 has finished with 9 states (4033 total)
2025/11/01 18:30:25 unblock.go:357: step #20 input desks 9
2025/11/01 18:30:25 unblock.go:389: step #20 has finished with 29 states (4062 total)
2025/11/01 18:30:25 unblock.go:357: step #21 input desks 29
2025/11/01 18:30:25 unblock.go:389: step #21 has finished with 89 states (4151 total)
2025/11/01 18:30:25 unblock.go:357: step #22 input desks 89
2025/11/01 18:30:25 unblock.go:389: step #22 has finished with 329 states (4480 total)
2025/11/01 18:30:25 unblock.go:357: step #23 input desks 329
2025/11/01 18:30:25 unblock.go:389: step #23 has finished with 823 states (5303 total)
2025/11/01 18:30:25 unblock.go:357: step #24 input desks 823
2025/11/01 18:30:25 unblock.go:382: step #24 has found the solution after 616 attempts (5920 total)
============================================
The solutions is found in 25 steps!!!
=0==0==0=|0| .  .
 . |0| . |0||0||0|
 . |0|=0==0=|0||0|
=0==0=|0| . |0||0|
 .  . |0|=0==0= .
=0==0==0= .  .  .

=0==0==0=|0||0| .
 . |0| . |0||0||0|
 . |0|=0==0=|0||0|
=0==0=|0| .  . |0|
 .  . |0|=0==0= .
=0==0==0= .  .  .

=0==0==0=|0||0||0|
 . |0| . |0||0||0|
 . |0|=0==0=|0||0|
=0==0=|0| .  .  .
 .  . |0|=0==0= .
=0==0==0= .  .  .

=0==0==0=|0||0||0|
 . |0| . |0||0||0|
 . |0|=0==0=|0||0|
=0==0=|0| .  .  .
 .  . |0| . =0==0=
=0==0==0= .  .  .

=0==0==0=|0||0||0|
 . |0| . |0||0||0|
 . |0|=0==0=|0||0|
=0==0=|0| .  .  .
 .  . |0| . =0==0=
 .  .  . =0==0==0=

=0==0==0=|0||0||0|
 . |0| . |0||0||0|
 . |0|=0==0=|0||0|
=0==0= .  .  .  .
 .  . |0| . =0==0=
 .  . |0|=0==0==0=

=0==0==0=|0||0||0|
 . |0| . |0||0||0|
 . |0|=0==0=|0||0|
 .  .  .  . =0==0=
 .  . |0| . =0==0=
 .  . |0|=0==0==0=

=0==0==0=|0||0||0|
 .  .  . |0||0||0|
 .  . =0==0=|0||0|
 . |0| .  . =0==0=
 . |0||0| . =0==0=
 .  . |0|=0==0==0=

=0==0==0=|0||0||0|
 .  .  . |0||0||0|
=0==0= .  . |0||0|
 . |0| .  . =0==0=
 . |0||0| . =0==0=
 .  . |0|=0==0==0=

=0==0==0=|0||0||0|
 .  . |0||0||0||0|
=0==0=|0| . |0||0|
 . |0| .  . =0==0=
 . |0| .  . =0==0=
 .  .  . =0==0==0=

=0==0==0=|0||0||0|
 .  . |0||0||0||0|
=0==0=|0| . |0||0|
 . |0| .  . =0==0=
 . |0| .  . =0==0=
=0==0==0= .  .  .

=0==0==0= . |0||0|
 .  . |0| . |0||0|
=0==0=|0| . |0||0|
 . |0| .  . =0==0=
 . |0| . |0|=0==0=
=0==0==0=|0| .  .

=0==0==0= . |0||0|
 .  . |0| . |0||0|
=0==0=|0| . |0||0|
 . |0|=0==0= .  .
 . |0| . |0|=0==0=
=0==0==0=|0| .  .

=0==0==0= .  . |0|
 .  . |0| . |0||0|
=0==0=|0| . |0||0|
 . |0|=0==0=|0| .
 . |0| . |0|=0==0=
=0==0==0=|0| .  .

=0==0==0= .  .  .
 .  . |0| . |0||0|
=0==0=|0| . |0||0|
 . |0|=0==0=|0||0|
 . |0| . |0|=0==0=
=0==0==0=|0| .  .

 .  .  . =0==0==0=
 .  . |0| . |0||0|
=0==0=|0| . |0||0|
 . |0|=0==0=|0||0|
 . |0| . |0|=0==0=
=0==0==0=|0| .  .

 .  . |0|=0==0==0=
 .  . |0| . |0||0|
=0==0= .  . |0||0|
 . |0|=0==0=|0||0|
 . |0| . |0|=0==0=
=0==0==0=|0| .  .

 .  . |0|=0==0==0=
 .  . |0| . |0||0|
 .  . =0==0=|0||0|
 . |0|=0==0=|0||0|
 . |0| . |0|=0==0=
=0==0==0=|0| .  .

 . |0||0|=0==0==0=
 . |0||0| . |0||0|
 .  . =0==0=|0||0|
 .  . =0==0=|0||0|
 .  .  . |0|=0==0=
=0==0==0=|0| .  .

 . |0||0|=0==0==0=
 . |0||0| . |0||0|
=0==0= .  . |0||0|
 .  . =0==0=|0||0|
 .  .  . |0|=0==0=
=0==0==0=|0| .  .

 . |0||0|=0==0==0=
 . |0||0| . |0||0|
=0==0= .  . |0||0|
=0==0= .  . |0||0|
 .  .  . |0|=0==0=
=0==0==0=|0| .  .

 . |0||0|=0==0==0=
 . |0||0||0||0||0|
=0==0= . |0||0||0|
=0==0= .  . |0||0|
 .  .  .  . =0==0=
=0==0==0= .  .  .

 . |0||0|=0==0==0=
 . |0||0||0||0||0|
=0==0= . |0||0||0|
=0==0= .  . |0||0|
=0==0= .  .  .  .
=0==0==0= .  .  .

 . |0||0|=0==0==0=
 . |0||0| . |0||0|
=0==0= .  . |0||0|
=0==0= . |0||0||0|
=0==0= . |0| .  .
=0==0==0= .  .  .

 . |0||0|=0==0==0=
 . |0||0| .  . |0|
=0==0= .  .  . |0|
=0==0= . |0||0||0|
=0==0= . |0||0| .
=0==0==0= . |0| .

 . |0||0|=0==0==0=
 . |0||0| .  .  .
=0==0= .  .  .  .
=0==0= . |0||0||0|
=0==0= . |0||0||0|
=0==0==0= . |0||0|
```
