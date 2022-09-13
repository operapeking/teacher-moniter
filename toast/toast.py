"""
Note:
Use Command "nohup" to run in the background like:
nohup python toast.py & >toast.log
use Command "jobs" to check whether it is running successfully
"""
import requests
import platform
import os
import time
if platform.system() != "Linux":
    print("Program can only run in linux system.")
    exit(0)
last = 2
url = "http://172.100.109.135:10086/query"
def ChangeLog(data):
    print("{} changed status to {} in {}".format(data["ip"], data["isComing"], data["time"]))
def Check(Last, Now):
    if Last == 1 and Now == 0:
        os.system('notify-send "No one is around"')
        return True
    if Last == 0 and Now == 1:
        os.system('notify-send "Someone is coming"')
        return True
    return False
while True:
    time.sleep(0.1)
    data = requests.get(url).json()
    if Check(last, data["isComing"]) == True:
        ChangeLog(data)
    last = data["isComing"]