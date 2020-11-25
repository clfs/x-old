# x

Go experiments.

## github.com/clfs/x/cmd/stamp

Apply timestamps to piped text.

```text
Usage:
    ... | stamp [-f FORMAT] [-b]

Options:
    -b          Place the timestamp in brackets.
    -f FORMAT   Choose a timestamp format (optional; default "medium").
    -h          Show this help message.

Formats:
    small       1:00pm
    medium      13:00:00.000
    large       2020-11-25T13:00:00.000000000-08:00
    unix        1606329242

Example:
    $ tail -f important.log | stamp -f small -b
    [3:26am] An event.
    [3:35am] Another event, sometime later.
```
