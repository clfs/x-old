# ensure
Check hashes with Unix pipes.

## Disclaimer
This tool provides **no** additional security - it only prevents accidental
data modification. If you need to prevent *malicious* data modification, you
should use a public-key signature system.

## Installation
```
go get github.com/clfs/ensure
```

## Usage
```
$ ensure
Usage:
	$ ... | ensure md5 1a79a4d60de6718e8e5b326e338ae533 | ...
	Check the hash, then pass standard input to standard output.
Options:
	-list	List all supported algorithms.
	-help	Print this help message.
	-quiet	Suppress error messages.
```

## Inspiration
I thought of this while installing Rust. As of now, the recommended `rustup`
[installation] method looks like this:
```
curl https://sh.rustup.rs -sSf | sh
```

The `curl` flags are:
```
-s, --silent        Silent mode (don't output anything)
-S, --show-error    Show error. With -s, make curl show errors when they occur
-f, --fail          Fail silently (no output at all) on HTTP errors (H)
```

If you check `man curl` though, you'll see that the `-f` flag isn't exactly
fail-safe. At minimum, the 401 (Unauthorized) and 407 (Proxy Authentication
Required) HTML response codes still cause `curl` to print to standard ouput.
```
-f, --fail
       (HTTP)  Fail  silently (no output at all) on server errors. This
       is mostly done to better enable scripts etc to better deal  with
       failed  attempts.  In  normal cases when an HTTP server fails to
       deliver a document, it  returns  an  HTML  document  stating  so
       (which  often  also describes why and more). This flag will pre‐
       vent curl from outputting that and return error 22.

       This method is not fail-safe and there are occasions where  non-
       successful  response  codes  will  slip through, especially when
       authentication is involved (response codes 401 and 407).
```

Instead, you can strengthen the `-f` flag by using `ensure` in tandem. Here's
what that might look like (with the hash truncated for clarity):
```
curl https://sh.rustup.rs -sSf | ensure sha256 9bbf4987[...] | sh
```

This isn't a great solution, to be fair - I'd definitely prefer the `rustup`
team use signing keys for their installation script. The [RVM] installer is a
useful example, even if it uses outdated GPG tooling.

## License
MIT; check the LICENSE file.

[installation]: https://github.com/rust-lang/rustup.rs/#other-installation-methods
[RVM]: https://rvm.io/rvm/install
