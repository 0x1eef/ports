PORTNAME=           portzap
PORTVERSION=        0.5.2
DISTVERSIONPREFIX=  v
CATEGORIES=         hardenedbsd

MAINTAINER=     0x1eef@protonmail.com
COMMENT=        portzap is a utility for staying up to date with the HardenedBSD ports tree.
WWW=            https://git.hardenedbsd.org/0x1eef/portzap
LICENSE=        BSD0CLAUSE
USE_GITHUB=     yes
GH_ACCOUNT=     0x1eef

BATCH=          yes
NO_ARCH=        yes
NO_BUILD=       yes
PLIST_FILES=	libexec/portzap/install-file \
		libexec/portzap/install-directory \
		libexec/portzap/commands/clone \
		libexec/portzap/commands/pull \
		libexec/portzap/commands/install \
		libexec/portzap/functions/requirements.sh \
		bin/portzap
do-install:
	${MKDIR} ${STAGEDIR}${PREFIX}/libexec/portzap
	${MKDIR} ${STAGEDIR}${PREFIX}/libexec/portzap/commands
	${MKDIR} ${STAGEDIR}${PREFIX}/libexec/portzap/functions
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/functions/requirements.sh ${STAGEDIR}${PREFIX}/libexec/portzap/functions/
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/commands/clone ${STAGEDIR}${PREFIX}/libexec/portzap/commands/
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/commands/pull ${STAGEDIR}${PREFIX}/libexec/portzap/commands/
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/commands/install ${STAGEDIR}${PREFIX}/libexec/portzap/commands/
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/install-file ${STAGEDIR}${PREFIX}/libexec/portzap/
	${INSTALL_SCRIPT} ${WRKSRC}/libexec/portzap/install-directory ${STAGEDIR}${PREFIX}/libexec/portzap/
	${INSTALL_SCRIPT} ${WRKSRC}/bin/portzap ${STAGEDIR}${PREFIX}/bin/

.include <bsd.port.mk>
