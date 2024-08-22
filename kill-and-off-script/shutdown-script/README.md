# System Window Manager

A Go project to close all open windows and shut down the system across different operating systems. The project supports Windows, macOS, and Linux.

## Overview

This project provides a CLI tool written in Go that performs the following actions based on the operating system:

- **Windows**: Closes all open windows and shuts down the system.
- **macOS**: Closes all open applications and shuts down the system.
- **Linux**: Closes all open windows and shuts down the system.

## Installation

### Linux

To run the project on Linux, you need to install `wmctrl`:

```bash
sudo apt-get install wmctrl
```
###Windows
Ensure you have Go installed. You do not need any additional packages for Windows.

###macOS
Ensure you have Go installed. No additional packages are needed.

##Usage
1. Build the project:
```bash
go build -o system-manager
```
2.Run the executable:
```bash
./system-manager
```
The executable will automatically detect your operating system and execute the appropriate actions.

##Files
- *main.go*: Contains the main logic to detect the operating system and call the respective functions.
- *linux.go*: Implements window closing and shutdown functions for Linux.
- *macos.go*: Implements window closing and shutdown functions for macOS.
- *windows.go*: Implements window closing and shutdown functions for Windows.

##Contributing
Feel free to contribute by submitting issues or pull requests.

##License
This project is licensed under the MIT License. See the LICENSE file for details.
