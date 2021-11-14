# laptop-control
---
Just another means of controlling my laptop's hardware in a simple way.

### Tips:
In order to use this tool to control your screen brightness, you're going to want to install the file at ./udev/rules.d/brightness-permissions.rules into the directory `/etc/udev/rules.d/` and run the commands `udevadm control --reload-rules && udevadm trigger`. This is all contingent on your user being a member of the `video` group.

### Aspects expected to be able to manage:
* volume
* screen brightness
* keyboard brightness
