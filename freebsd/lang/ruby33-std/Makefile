PORTNAME= ruby-std
DISTVERSION= 3.3.0
RUBY_MAJOR=33
RUBY_SONAME="ruby${RUBY_MAJOR}"
CATEGORIES= lang ruby
DISTNAME=ruby-${DISTVERSION}

MASTER_SITES= https://cache.ruby-lang.org/pub/ruby/3.3/

MAINTAINER= 0x1eef@protonmail.com
COMMENT= A dynamic, open source programming language with a focus on simplicity and productivity.
WWW= https://www.ruby-lang.org/en/

LICENSE= BSD2CLAUSE RUBY
LICENSE_COMB= dual
LICENSE_FILE_BSD2CLAUSE= ${WRKSRC}/BSDL
LICENSE_FILE_RUBY= ${WRKSRC}/COPYING

GNU_CONFIGURE=  yes
CONFIGURE_ARGS= --with-soname=${RUBY_SONAME} \
	--program-suffix=${RUBY_MAJOR} \
	--enable-pthread \
	--enable-shared \
	--enable-rjit \
	--prefix="${PREFIX}"
CFLAGS+=-I/usr/include -I${PREFIX}/include
USE_LDCONFIG=yes

USES=           ssl tar:xz
BUILD_DEPENDS=  libffi>=0:devel/libffi
LIB_DEPENDS=    libyaml.so:textproc/libyaml
RUN_DEPENDS=    libffi>=0:devel/libffi

.include <bsd.port.mk>
