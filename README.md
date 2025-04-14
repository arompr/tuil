# 🌙 tuil - A Work in progress -
A Terminal UI for Brightness & Night Light Control to manually adjust screen brightness and night light right from your terminal.

I run a very lightweight Hyprland setup, and while it's incredibly efficient, I found myself missing a **user-friendly** and quick way to tweak brightness and night light without leaving my terminal.

Tuil fills that gap. While it's tailored for Hyprland via hyprsunset for nightlight control, it was designed to be easily adaptable—simply implement a new adapter and pass it to the nightlight use cases instead of the hyprsunset adapter.

![image](https://github.com/user-attachments/assets/5bcd5691-a947-47fe-86e0-cfa7ffe73c9e)

## 🛠️ TODO

- [ ] Cleanup `tui.go` for better structure and readability
- [ ] Improve "cleanness" of store caching
- [ ] Only initialize `hyprsunsetAdapter` if **Hyprland is installed and running**  
- [ ] Improve error logging  
