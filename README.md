# DESCRIPTION

* the old go-bindata is no longer maintained, and has been forked.
* I would love to migrate to the new go-bindata, but it does not support the CLI interface.
* But I LIKE the CLI interface.

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