## About

The lang/ruby33 FreeBSD port provides a standard installation of
Ruby 3.3 that replaces the
[official lang/ruby33 port](https://github.com/freebsd/freebsd-ports/tree/main/lang/ruby33).
This port installs Ruby's
[default and bundled gems](https://www.stdgems.org),
and includes their executables.

## Executables

The port installs the following executables:

* bin/bundle33
* bin/bundler33
* bin/erb33
* bin/gem33
* bin/irb33
* bin/racc33
* bin/rake33
* bin/rbs33
* bin/rdbg33
* bin/rdoc33
* bin/ri33
* bin/ruby33
* bin/syntax_suggest33
* bin/typeprof33

## Notes

* [hardenedbsd](https://hardenedbsd.org)

    * Set `hardening.harden_rtld` to 0. Otherwise the build will fail.

            # sysctl hardening.harden_rtld=0

## See also

* [devel/ruby-shims](https://github.com/0x1eef/ports/tree/main/freebsd/devel/ruby-shims)

## Changelog

* 2024/04/23: add v3.3.1
* 2023/12/25: add v3.3.0
