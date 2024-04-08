PORTNAME=           portzap
PORTVERSION=        0.7.3
DISTVERSIONPREFIX=  v
CATEGORIES=         hardenedbsd

MAINTAINER=     0x1eef@protonmail.com
COMMENT=        portzap manages a local copy of the hardenedbsd ports tree.
WWW=            https://git.hardenedbsd.org/0x1eef/portzap
LICENSE=        BSD0CLAUSE
USE_GITLAB=     yes
GL_SITE=	https://git.hardenedbsd.org
GL_ACCOUNT=     0x1eef
GL_PROJECT=	portzap
GL_COMMIT=	4e647bfbe4998627615b839810a32cd95320a8f8

RUN_DEPENDS=	doas:security/doas
BATCH=          yes
NO_ARCH=        yes
NO_BUILD=       yes
MAN=	man/man8/portzap.8
PLIST_FILES= bin/portzap \
	libexec/portzap/git-changed-files \
	libexec/portzap/git-removed-files \
	libexec/portzap/git-rev \
	libexec/portzap/portzap-clone \
	libexec/portzap/portzap-install \
	libexec/portzap/portzap-pull \
	libexec/portzap/portzap-setup \
	man/man8/portzap.8.gz \
	%%DATADIR%%/doas.conf

do-install:
	${MAKE} -C ${WRKSRC} PREFIX=${STAGEDIR}${PREFIX} install

.include <bsd.port.mk>
