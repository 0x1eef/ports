PORTNAME= ruby
CATEGORIES= lang ruby
DISTVERSION= 3.2.1
MASTER_SITES= https://cache.ruby-lang.org/pub/ruby/3.2/

MAINTAINER= 0x1eef@protonmail.com
COMMENT= A dynamic, open source programming language with a focus on simplicity and productivity.
WWW= https://www.ruby-lang.org/en/

LICENSE= BSD2CLAUSE RUBY
LICENSE_COMB= dual
LICENSE_FILE_BSD2CLAUSE= ${WRKSRC}/BSDL
LICENSE_FILE_RUBY= ${WRKSRC}/COPYING

GNU_CONFIGURE=  yes
CONFIGURE_ARGS= --enable-pthread --enable-shared --prefix="${PREFIX}" --program-suffix=32
USES=           ssl tar:xz
BUILD_DEPENDS=  libffi>=0:devel/libffi
LIB_DEPENDS=    libyaml.so:textproc/libyaml
RUN_DEPENDS=    libffi>=0:devel/libffi
USE_LDCONFIG=	yes

##
# Options
OPTIONS_DEFINE= LIBEDIT READLINE YJIT
OPTIONS_DEFAULT= LIBEDIT
OPTIONS_SINGLE= EDIT
OPTIONS_SINGLE_EDIT= LIBEDIT READLINE
LIBEDIT_DESC= Use libedit
READLINE_DESC= Use libreadline
YJIT_DESC= Enable YJIT

##
# libedit
LIBEDIT_USES= libedit
LIBEDIT_CONFIGURE_ON= --enable-libedit --with-libedit-dir=${LOCALBASE}

##
# Readline
READLINE_USES= readline
READLINE_CONFIGURE_ON= --disable-libedit --with-readline-dir=${LOCALBASE}

##
# YJIT
YJIT_BUILD_DEPENDS= ${RUST_DEFAULT}>=1.58.0:lang/${RUST_DEFAULT}
YJIT_CONFIGURE_ENABLE= yjit

.include <bsd.port.mk>
