## devel/ruby-shims

This port brings the OpenBSD port of the same name
to FreeBSD. The port is intended to be combined with
the ruby ports within this repository (eg
[/lang/ruby33](https://github.com/0x1eef/ports/tree/main/freebsd/lang/ruby33)
).

## Features

* Adds support for a project-local `.ruby-version` file.
* Adds support for a system-wide `/usr/local/etc/ruby-version` file.
* Adds shortcuts to the active Ruby's executables (eg `ruby` instead of `ruby32`, etc).

## Changelog

* Remove `set +o sh` from `rubyshim.sh`.
* Modify `rubyshim.sh` to be compatible with `/bin/sh` on FreeBSD.
* Change system-wide path from `/etc/ruby-version` to `/usr/local/etc/ruby-version`.
* Replace backticks with `$()`.
