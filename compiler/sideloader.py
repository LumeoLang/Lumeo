import os

global file_name
global file_icon

def ask():
    global file_name
    global file_icon
    file_name = input("File name: ")
    file_icon = input("File icon: ")

def issue(message):
    print(message)
    ask()


if not file_name.endswith(".exe"):
    file_name += ".exe"

if not file_name.endswith(".ico"):
    issue("Icon file must be .ico")