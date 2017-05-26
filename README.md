# DESCRIPTION

* the old [go-bindata](https://github.com/jteeuwen/go-bindata) is no longer maintained.
* It and has been forked [several](https://github.com/lestrrat/go-bindata) [times](https://github.com/tmthrgd/go-bindata).
* I would love to migrate to the new [go-bindata](https://github.com/tmthrgd/go-bindata), but it does not support the CLI interface.
* But I LIKE the CLI interface. :(

So here it is.

* Pack static file assets into a byte array.
* byte array contents are compressed

TODO

* Store file metadata

# Usage (CLI)

```
$ ls templates | packasset
```

```
$ packasset -input file.txt
```