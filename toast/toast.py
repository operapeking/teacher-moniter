import requests
import platform
import os

if platform.system() == "Windows":
    import win10toast
    toaster = win10toast.ToastNotifier()
last = 2
system = platform.system()
def WindowsDanger():
    toaster.show_toast('Warning', 'Someone is coming')
def WindowsSafe():
    toaster.show_toast('Safe', 'No one is around')
def LinuxDanger():
    os.system('notify-send "Someone is coming"')
def LinuxSafe():
    os.system('notify-send "No one is around"')
def Check(Last, Now):
    if Last == 1 and Now == 0:
        if platform.system() == "Windows":
            WindowsSafe()
        if platform.system() == "Linux":
            LinuxSafe()
    if Last == 0 and Now == 1:
        if platform.system() == "Windows":
            WindowsDanger()
        if platform.system() == "Linux":
            LinuxDanger()
while True:
    try:
        data = requests.get("http://172.100.109.135:10086/query").json()
        Check(last, data["isComing"])
        last = data["isComing"]
    except:
        print("Exit")
        exit(1)