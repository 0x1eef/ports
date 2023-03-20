## devel/ruby-shims

The devel/ruby-shims port brings the OpenBSD port of the same name
to FreeBSD. <br>
It is assumed that the port will be used with the
`lang/rubyXX-std` port(s)  (eg
[/lang/ruby32-std](https://github.com/0x1eef/ports/tree/main/freebsd/lang/ruby32-std)
).


## Features

* Adds support for a project-local `.ruby-version` file.
* Adds support for a system-wide `/usr/local/etc/ruby-version` file.
* Adds shortcuts to the active Ruby's executables (eg `ruby` instead of `ruby32`, etc).

## Changes

The following changes have been made to the original OpenBSD port:

* Remove `set +o sh` from `rubyshim.sh`.
* Modify `rubyshim.sh` to be compatible with `/bin/sh` on FreeBSD.
* Change system-wide path from `/etc/ruby-version` to `/usr/local/etc/ruby-version`.
* Replace backticks with `$()`.
