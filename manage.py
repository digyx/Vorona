import os
import subprocess
import sys
import time

import pymongo
import requests


def initialize():
    subprocess.run("docker-compose build --no-cache --pull", shell=True)
    subprocess.run("docker-compose up -d", shell=True)

    check_vue()


def check_vue():
    print("Installing Vue packages...")

    while True:
        time.sleep(1)
        try:
            req = requests.get("http://localhost:5000")

            if req.status_code == 200:
                break
        except:
            pass


def git_update():
    subprocess.run("git pull", shell=True)


def reload(service=None):
    if service == None:
        subprocess.run("docker-compose down", shell=True)
        subprocess.run("docker-compose build", shell=True)
        subprocess.run("docker-compose up -d", shell=True)
        check_vue()
        return

    subprocess.run("docker-compose stop {}".format(service), shell=True)
    subprocess.run("docker-compose build {}".format(service), shell=True)
    subprocess.run("docker-compose up -d {}".format(service), shell=True)
    
    if service == "frotnend" or service == "artist-dashboard":
        check_vue()


if __name__ == "__main__":
    if len(sys.argv) == 1:
        print("Usage: python manage.py <command>")
        print("Use the command \"list\" to see available commands")
        print()
        exit()

    elif sys.argv[1] == "update":
        git_update()

    elif sys.argv[1] == "start":
        subprocess.run("docker-compose build", shell=True)
        subprocess.run("docker-compose up -d", shell=True)
        check_vue()

    elif sys.argv[1] == "stop":
        subprocess.run("docker-compose down", shell=True)

    elif sys.argv[1] == "reload":
        try:
            reload(sys.argv[2])
        except IndexError:
            reload()
    
    elif sys.argv[1] == "clean-start":
        subprocess.run("docker-compose build --no-cache --pull", shell=True)
        subprocess.run("docker-compose up -d", shell=True)
        check_vue()
    
    elif sys.argv[1] == "server":
        subprocess.run("docker-compose down -f docker-compose-server.yml api mongo", shell=True)
        subprocess.run("git pull", shell=True)
        subprocess.run("docker-compose build -f docker-compose-server.yml api mongo", shell=True)
        subprocess.run("docker-compose up -d -f docker-compose-server.yml api mongo", shell=True)
    
    elif sys.argv[1] == "list":
        print("python manage.py <command>")
        print()
        print("Available commands:")
        print("\tupdate       Run git pull on repo")
        print("\tstart        Start the development environment")
        print("\tstop         Stop the development environment")
        print("\treload       Restart the dev environment")
        print("\tclean-start  Start dev environment without docker cache")
        print("\tserver       Start a production instance")
        print()

    else:
        print("Invalid command")
        exit()
    
    print("Done.")
