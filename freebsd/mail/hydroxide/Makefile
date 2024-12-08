PORTNAME=           hydroxide
PORTVERSION=        0.2.29
DISTVERSIONPREFIX=  v
CATEGORIES=         mail

MAINTAINER=     0x1eef@protonmail.com
COMMENT=        A third-party, open-source ProtonMail bridge.
WWW=            https://github.com/emersion/hydroxide
LICENSE=        MIT
LICENSE_FILE=   ${WRKSRC}/LICENSE

USES=           go:modules
GO_PKGNAME=     hydroxide
GO_MODULE=      github.com/emersion/hydroxide
GO_TARGET=      ./cmd/hydroxide/
GO_ENV+=        GOMAXPROCS=1
PLIST_FILES=    bin/hydroxide
USE_RC_SUBR=    hydroxide

.include <bsd.port.mk>
