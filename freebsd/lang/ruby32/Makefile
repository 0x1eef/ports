PORTNAME= ruby32
DISTVERSION= 3.2.4
CATEGORIES= lang ruby
MASTER_SITES= https://cache.ruby-lang.org/pub/ruby/3.2/
DISTNAME=ruby-${DISTVERSION}

MAINTAINER= 0x1eef@protonmail.com
COMMENT=Dynamic, open source programming language
WWW= https://www.ruby-lang.org/en/

LICENSE= BSD2CLAUSE RUBY
LICENSE_COMB= dual
LICENSE_FILE_BSD2CLAUSE= ${WRKSRC}/BSDL
LICENSE_FILE_RUBY= ${WRKSRC}/COPYING

BUILD_DEPENDS=	libffi>=0:devel/libffi
LIB_DEPENDS=	libyaml.so:textproc/libyaml
RUN_DEPENDS=	libffi>=0:devel/libffi

RUBYVER=32
RUBYSONAME="ruby${RUBYVER}"
USES=	ssl tar:xz
GNU_CONFIGURE=  yes
CONFIGURE_ARGS= --with-soname=${RUBYSONAME} \
	--program-suffix=${RUBYVER} \
	--enable-pthread \
	--enable-shared \
	--disable-yjit \
	--enable-rjit \
	--prefix="${PREFIX}"
CFLAGS+=-I/usr/include -I${PREFIX}/include
USE_LDCONFIG=yes

.include <bsd.port.mk>
