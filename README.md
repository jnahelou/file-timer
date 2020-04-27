# file-timer
*file-timer* is a simple program to start a timer and save the remaining time in a file.

By example, you can use it as countdown for OBS using text from file source.

## How to use
### Getting Help
```
$ file-timer -h
Usage of file-timer:
  -duration string
        Countdown duration (default "00h15m00s")
  -out-file string
        Path to OBS watchfile (default "/tmp/countdown.raw")

```
### Examples
Start a 1h timer :
```
$ file-timer -duration 1h
```

Change out file :
```
$ file-timer -out-file ~/dummy-file
```

## Contribute
Feel free to open issues or pull-requests
