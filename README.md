# x

Go experiments.

## github.com/clfs/x/cmd/stamp

```text
Usage:
    ... | stamp [-f FORMAT] [-b]

Options:
    -b          Place the timestamp in brackets.
    -f FORMAT   Choose a timestamp format (optional; default "medium").
    -h          Show this help message.

Formats:
    small       10:33am
    medium      10:33:36.042
    large       2020-11-25T10:33:55.740453-08:00
    unix        1606329242

Example:
    $ tail -f important.log | stamp -f small -b
    [10:26am] An event.
    [10:35am] Another event, sometime later.
```
