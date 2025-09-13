# 🌙 tuil

A Terminal UI for Brightness & Night Light Control to manually adjust screen brightness and night light right from terminal.

I run a very lightweight Hyprland setup, and while it's incredibly efficient, I found myself missing a quick and **user-friendly** way to tweak brightness and night light without leaving terminal.

tuil fills that gap for me.

I also made a simple ctl that starts Hyprsunset with the last saved value.
If you want to use it at startup as I do, build it then add to your Hyprland config file:

```console
exec-once tuilctl toggle last
```

![image](https://github.com/user-attachments/assets/5bcd5691-a947-47fe-86e0-cfa7ffe73c9e)

## Requirements

- Go
- **Hyprland** >= 0.48  
- **hyprsunset** >= 0.2  
- **brightnessctl** (used for screen brightness control)

## Build instructions

Build tuil executable

```console
go build -o tuil
```

Build tuilctl executable

```console
go build -o tuilctl ./ctl
```

Add to path. For example:

```console
sudo mv tuil /usr/local/bin/
sudo mv tuilctl /usr/local/bin/
```

## tuilctl Usage

- `tuilctl toggle night`  
  Applies the last saved nightlight temperature.
- `tuilctl toggle light`  
  Applies the light temperature to 6000K (turns off nightlight).
- `tuilctl toggle last`
  Applies the last saved temperature (night or light).

## 🛠️ TODO

- [ ] Cleanup `tui.go` for better structure and readability
- [ ] Improve error logging
- [ ] Apply rollback on process kill or close without save
- [ ] Improve min/max value handling
- [ ] Add to the AUR - I use Arch, btw
