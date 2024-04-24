## About

The lang/ruby32 FreeBSD port provides a standard installation of
Ruby 3.2 that replaces the
[official lang/ruby32 port](https://github.com/freebsd/freebsd-ports/tree/main/lang/ruby32).
This port installs Ruby's
[default and bundled gems](https://www.stdgems.org),
and includes their executables.

## Executables

The port installs the following executables:

* bin/ruby32
* bin/gem32
* bin/irb32
* bin/bundle32
* bin/bundler32
* bin/erb32
* bin/racc32
* bin/rake32
* bin/rbs32
* bin/rdbg32
* bin/rdoc32
* bin/ri32
* bin/typeprof32

## Notes

* [hardenedbsd](https://hardenedbsd.org)

    * Set `hardening.harden_rtld` to 0. Otherwise the build will fail.

            # sysctl hardening.harden_rtld=0

## See also

* [devel/ruby-shims](https://github.com/0x1eef/ports/tree/main/freebsd/devel/ruby-shims)

## Changelog

* 2024/14/23: Add v3.2.4
* 2023/12/25: Rewrite Makefile.
* 2023/12/13: Add v3.2.2
