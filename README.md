update-checks
=============

Scrapes websites to check for updates of various programs.

I'm lazy, so I built this to let me know when packages I maintain on aur.archlinux.org have new releases (before someone marks it out of date on the AUR).
Basically just scrapes the release page for each program, and sends me an email when the latest version on that page is newer than the version in the AUR.
Slap this sucker in cron, and its package maintenance for dummies.

Things are shitty and hard-coded, so you prolly don't want to use this without adding flags for email address and such.
