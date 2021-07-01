# ensure

Check piped data for the correct hash.

This tool doesn't provide strong security. If you really need to
prevent attacks on data integrity, you should rely on signatures
rather than known hashes.

## Installation

```text
go install github.com/clfs/x/ensure@latest
```

## Usage

```text
$ ensure -help
usage:
        ensure -help
        ensure -list
        ... | ensure -alg crc32 -sum 2747fc56 | ...
```

## Inspiration

This is Rust's current [`rustup`](https://rustup.rs/) install command:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

The `curl` flags are:

```text
-s, --silent        Silent mode
-S, --show-error    Show error even when -s is used
-f, --fail          Fail silently (no output at all) on HTTP errors
```

If you check `man curl` though, you'll see that the `-f` flag doesn't
always work. HTTP 401 (Unauthorized) and HTTP 407 (Proxy Authentication
Required) status codes can still cause `curl` to print to standard ouput:

```text
-f, --fail
       (HTTP) Fail silently (no output at all) on server  errors.  This
       is  mostly done to better enable scripts etc to better deal with
       failed attempts. In normal cases when an HTTP  server  fails  to
       deliver  a  document,  it  returns  an  HTML document stating so
       (which often also describes why and more). This flag  will  pre-
       vent curl from outputting that and return error 22.

       This  method is not fail-safe and there are occasions where non-
       successful response codes will  slip  through,  especially  when
       authentication is involved (response codes 401 and 407).
```

I haven't validated this assumption, but maybe an network-level attacker
could serve malicious content via a HTTP 407 response, which would get
executed by `sh`. My hunch is that I'm wrong, but `ensure` would at least
prevent that outright:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs \
  | ensure -alg crc32 -sum a6026086 \
  | sh
```
