== mail/hydroxide

hydroxide is a third-party, open source ProtonMail bridge. The bridge translates SMTP, and IMAP
requests to HTTP requests that are forwarded to the ProtonMail Web API. The bridge allows a
ProtonMail user with a paid subscription to use their email client of choice rather than being
restricted solely to ProtonMail's web client.

== /etc/rc.conf variables

* hydroxide_enable
  When set to "yes", hydroxide will start at boot time. Defaults to "no".

* hydroxide_user
  The user who will run hydroxide. Defaults to "_hydroxide".

* hydroxide_flags
  The command line arguments given to hydroxide. Defaults to "-disable-carddav".
