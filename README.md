# SE Player

A remote-controllable web-based SE player for OBS.



## How to

1. Add some files to `contents/se/` directory. (prefered: wav)
	* mp3 files have a short blank (~= 0.1 sec) at the beginning of the audio.
	* You can remove blank and export to wav by Audacity.
1. Launch `se-player.exe` (linux: `se-player.elf`, OS X: `se-player.macho`)
1. Open http://127.0.0.1:3000/control.html in your browser.

### Simple

4. Click `Play at HERE`
1. Click filename button

### Advance (Control audio via OBS / Poor latency)

4. Launch OBS
1. Add browser source
1. Set URL to `http://127.0.0.1:3000/player.html`
1. Set Width=~160 Height=~90 (for check status)
1. Set `Control audio via OBS`
1. Apply
1. Click filename button at Control UI

