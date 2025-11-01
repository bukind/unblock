# Solver for the Unblock game on Android

## Usage

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
