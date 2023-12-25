PORTNAME= ruby32-std
RUBYSTD_VER=32
RUBYSTD_SONAME="ruby${RUBYSTD_VER}"
DISTVERSION= 3.2.2
DISTNAME=ruby-${DISTVERSION}
MASTER_SITES= https://cache.ruby-lang.org/pub/ruby/3.2/
CATEGORIES= lang ruby

LICENSE= BSD2CLAUSE RUBY
LICENSE_COMB= dual
LICENSE_FILE_BSD2CLAUSE= ${WRKSRC}/BSDL
LICENSE_FILE_RUBY= ${WRKSRC}/COPYING

GNU_CONFIGURE=  yes
CONFIGURE_ARGS= --with-soname=${RUBYSTD_SONAME} \
	--program-suffix=${RUBYSTD_VER} \
	--enable-pthread \
	--enable-shared \
	--disable-yjit \
	--enable-rjit \
	--prefix="${PREFIX}"
CFLAGS+=-I/usr/include -I${PREFIX}/include
USE_LDCONFIG=yes

USES=           ssl tar:xz
BUILD_DEPENDS=  libffi>=0:devel/libffi
LIB_DEPENDS=    libyaml.so:textproc/libyaml
RUN_DEPENDS=    libffi>=0:devel/libffi

MAINTAINER= 0x1eef@protonmail.com
COMMENT= A dynamic, open source programming language with a focus on simplicity and productivity.
WWW= https://www.ruby-lang.org/en/

.include <bsd.port.mk>
