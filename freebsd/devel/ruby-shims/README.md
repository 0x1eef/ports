## About

The devel/ruby-shims port brings the OpenBSD port of the same name
to FreeBSD.

The port adds support for the `.ruby-version` file,
and a system-wide version file at `/usr/local/etc/ruby-version`.
With the port installed, the active Ruby version (and its executables)
can be called without a suffix (eg as `ruby` instead of `ruby32`, etc.)

## Expectations

This port is expected to be used with the
[/lang/ruby32-std](https://github.com/0x1eef/tree/main/FreeBSD/lang/ruby32-std),
port and other similar `rubyXX-std` ports.

## Changes

The following changes have been made to the original OpenBSD port:

* Remove `set +o sh` from `rubyshim.sh`.
* Modify `rubyshim.sh` to be compatible with `/bin/sh` on FreeBSD.
* Change system-wide path from `/etc/ruby-version` to `/usr/local/etc/ruby-version`.
* Replace backticks with `$()` syntax.
