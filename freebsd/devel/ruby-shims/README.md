## About

The devel/ruby-shims port brings the OpenBSD port of the same name
to FreeBSD. The port adds support for the often used `.ruby-version`
file, and provides shortcuts for calling Ruby executables (eg 'ruby32'
can be 'ruby', 'irb32' can be 'irb', and so on). The shortcuts reflect
the Ruby version that's set by a `.ruby-version` file, or otherwise -
the most recent Ruby version that's installed.

## Expectations

The port is expected to be used with the
[/lang/ruby32-std](https://github.com/0x1eef/tree/main/FreeBSD/lang/ruby32-std),
[/lang/ruby31-std](https://github.com/0x1eef/tree/main/FreeBSD/lang/ruby31-std),
etc ports.

## Changes

The following changes have been made to the original OpenBSD port:

* Remove `set +o sh` from `rubyshim.sh`.
* Modify `rubyshim.sh` to be compatible with `/bin/sh` on FreeBSD.
* Change system-wide path from `/etc/ruby-version` to `/usr/local/etc/ruby-version`.
* Replace backticks with `$()` syntax.
