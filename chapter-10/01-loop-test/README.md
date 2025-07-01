By default, this repository use go 1.24.3


## Test with current version

Run:
```sh
go build
```

Building and running the program gives the following output:
```
0xc000118040
0xc000118048
0xc000118060
0xc000118068
0xc000118070
```

## Test with earlier version

To test with previous version (e.g 1.18)

Update the go.mod to `1.18`

Run:
```
GOTOOLCHAIN=go1.18 go build
```

Now re run the built file, you would see:
```
0xc0000160e8
0xc0000160e8
0xc0000160e8
0xc0000160e8
0xc0000160e8
```

The program prints out the same memory address five times. 
(The memory addresses might differ from what’s shown here, but all will be the same value.)