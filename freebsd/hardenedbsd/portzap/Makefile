PORTNAME=           portzap
PORTVERSION=        0.4.0
DISTVERSIONPREFIX=  v
CATEGORIES=         hardenedbsd

MAINTAINER=     0x1eef@protonmail.com
COMMENT=        A utility for managing a local copy of the HardenedBSD ports collection
WWW=            https://git.hardenedbsd.org/0x1eef/portzap
LICENSE=        BSD0CLAUSE
USE_GITHUB=     yes
GH_ACCOUNT=     0x1eef

BATCH=          yes
NO_ARCH=        yes
NO_BUILD=       yes
PLIST_FILES=	libexec/portzap/install-file \
		libexec/portzap/install-directory \
		libexec/portzap/functions/fs.sh \
		libexec/portzap/functions/git.sh \
		libexec/portzap/functions/perms.sh \
		bin/portzap
do-install:
	${MKDIR} ${STAGEDIR}${PREFIX}/libexec/portzap
	${MKDIR} ${STAGEDIR}${PREFIX}/libexec/portzap/functions
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/functions/fs.sh ${STAGEDIR}${PREFIX}/libexec/portzap/functions/fs.sh
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/functions/git.sh ${STAGEDIR}${PREFIX}/libexec/portzap/functions/git.sh
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/functions/perms.sh ${STAGEDIR}${PREFIX}/libexec/portzap/functions/perms.sh
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/install-file ${STAGEDIR}${PREFIX}/libexec/portzap/install-file
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/install-directory ${STAGEDIR}${PREFIX}/libexec/portzap/install-directory
	${INSTALL_SCRIPT} ${WRKSRC}/bin/portzap ${STAGEDIR}${PREFIX}/bin/

.include <bsd.port.mk>
