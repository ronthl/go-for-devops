# Joining a file path
This branch is to demonstrate how to join a file path using `filepath.Join()` method.

The main benefit of it is that it uses the native pathing rules. On a Unix-like system, this might be:
```
/home/ronthl/go-for-dev/config/config.json
```
While on Windoes, this might be:
```
C:\Documents\ronth\go-for-dev\config\config.json
```

Checkout the `main.go` file.

**Concept**: A file path's final directory or file is called the **base**.
The path your binary is running in is called the **working directory**.
