# Simple CLI Tool For Automate startup apps

## Simple Explanation

This project has three distinct ways to manage and launch installed applications:<br>


---

### `main.go`: Direct Launcher

**`main.go`** kicks off a **fixed set of predefined applications** automatically. It's straightforward and has no command-line options—just run it, and your apps will start.

---

### `main2.go`: Flexible Sets

**`main2.go`** gives you choices. You can select from **different predefined groups of applications** using command-line arguments. For instance, you could launch `set1` (e.g., `app1`, `app2`, ...) or `set2` (e.g., `appN`, `appM`, ...) depending on your needs.<br>
At the moment is necessary to change these sets through code, I'll work on new solutions later on.

---

### `main3.go`: Custom App Bundles

**`main3.go`** is your tool for personalization. It lets you **create your own custom set of applications**. After you define your bundle, it saves the configuration to a text file and then builds a dedicated executable file, **`executer.go`**. This means you end up with a standalone executable tailored to launch *your* specific apps.

<img src="Go_Logo_Blue.svg.png" height="50px">
