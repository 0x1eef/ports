## mail/hydroxide

[hydroxide](https://github.com/emersion/hydroxide) is a third-party,
open source ProtonMail bridge.

## /etc/rc.conf variables

* `hydroxide_enable`
  When set to `yes`, hydroxide will start at boot. Defaults to `no`.
* `hydroxide_user`
  The user that runs the hydroxide daemon. Defaults to `_hydroxide`.
* `hydroxide_flags`
  The command line arguments given to hydroxide. Defaults to `-disable-carddav`.
