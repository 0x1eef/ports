## About

The devel/ruby-shims port brings the OpenBSD port of the same name
to FreeBSD.

## Features 

* Adds support for a project-local `.ruby-version` file. 
* Adds support for a system-wide `/usr/local/etc/ruby-version` file.
* Adds shortcuts to the active Ruby's executables (eg `ruby` instead of `ruby32`, etc).

## Expectations

It is expected that this port will be used with one or more `lang/rubyXX-std` ports 
(eg [/lang/ruby32-std](https://github.com/0x1eef/tree/main/FreeBSD/lang/ruby32-std)).

## Changes

The following changes have been made to the original OpenBSD port:

* Remove `set +o sh` from `rubyshim.sh`.
* Modify `rubyshim.sh` to be compatible with `/bin/sh` on FreeBSD.
* Change system-wide path from `/etc/ruby-version` to `/usr/local/etc/ruby-version`.
* Replace backticks with `$()`.
