PORTNAME=       ruby-shims
PORTVERSION=    0.1.0
CATEGORIES=     devel

MAINTAINER=     0x1eef@protonmail.com
COMMENT=        Ruby version redirection shim
WWW=            https://cvsweb.openbsd.org/cgi-bin/cvsweb/ports/devel/ruby-shims
LICENSE=        ISCL
BATCH=			1
DISTDIR=        ${FILESDIR}
DISTFILES=      rubyshim.sh
BINS=           bin/ruby \
                bin/irb \
                bin/rdoc \
                bin/ri \
                bin/rake \
                bin/gem \
                bin/bundle \
                bin/bundler \
                bin/rbs \
                bin/rdbg \
                bin/typeprof \
                bin/erb \
                bin/racc
PLIST_FILES=    libexec/rubyshim ${BINS}

do-extract:
do-fetch:
do-build:
	${MKDIR} ${STAGEDIR}${PREFIX}/libexec/ ${STAGEDIR}${PREFIX}/bin/
	${INSTALL_SCRIPT} ${DISTDIR}/rubyshim.sh ${STAGEDIR}${PREFIX}/libexec/rubyshim
.	for f in ${BINS}
		ln -sf ${PREFIX}/libexec/rubyshim ${STAGEDIR}${PREFIX}/${f}
.	endfor

.include <bsd.port.mk>
