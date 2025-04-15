# 🌙 tuil - A Work in progress -
A Terminal UI for Brightness & Night Light Control to manually adjust screen brightness and night light right from terminal.

I run a very lightweight Hyprland setup, and while it's incredibly efficient, I found myself missing a quick and **user-friendly** way to tweak brightness and night light without leaving terminal.

tuil fills that gap. While it uses hyprsunset for a tailored Hyprland nightlight control experience, it was designed to be easily adaptable - just implement a new adapter and pass it to the nightlight use cases instead of the hyprsunset adapter.

![image](https://github.com/user-attachments/assets/5bcd5691-a947-47fe-86e0-cfa7ffe73c9e)

## Requirements
- **Hyprland** >= 0.48  
- **hyprsunset** >= 0.2  
- **brightnessctl** (used for screen brightness control)


## 🛠️ TODO
- [ ] Cleanup `tui.go` for better structure and readability
- [ ] Only initialize `hyprsunsetAdapter` and nightlight slider if **Hyprland is installed and running**
- [ ] Improve error logging
- [ ] Add build instructions
- [ ] Improve min/max value handling 
- [ ] Add to the AUR - I use Arch, btw
