# gnp

Generate New Password - quick & dirty way to spit out a new password

## Installing

1. Git clone this repo
2. `go build -o gnp`
3. Do whatever you want with the output. Alias it, move it to your `~/bin`, etc.

## Useage

```
Usage of gnp:
  -l int
    	Password lenght (default 24)
  -n int
    	Minimum numbers (default 1)
  -s int
    	Minimum amount of special characters (default 1)
  -seed int
    	Additional custom seed (default 1337)
  -u int
    	Minimum uppercased characters (default 1)
```
