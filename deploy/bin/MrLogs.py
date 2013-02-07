#!/usr/bin/env python
# MrLogs takes syslog logs and cuts them up and resends them to a specified
# syslog server.
#
# Author: Alex Lisle alisle@alienvault.com

__version__ = '0.01'

import getopt
import sys
import re
import netsyslog
import random
import syslog
import time
from datetime import datetime

_MAX_LINES = 1000
_SERVERS = ["127.0.0.1"]
_USE_DATE = False
_EPS = 10
_ONLYONCE = False

messages = []

def ShowHelp():
    print "Welcome to MrLogs " + __version__
    print "Usage:"
    print "\t" + sys.argv[0] + " [OPTION]... [LOGS]..."
    print "\t-n (--number-of-lines)\tSpecify the number of lines to take from each log, default is 1000"
    print "\t-v (--version)\t\tPrint Version"
    print "\t-s (--server)\t\tSpecify the syslog server to use, you can give a comma seperated "
    print "\t-p (--port)\t\tSpecify a port number for your servers"
    print "\t\t\t\tlist to send to multiple servers,  default is localhost"
    print "\t-d (--keep-date)\tKeep the Dates within the logs when sending them, default is to ignore the date"
    print "\t-e (--eps)\tState the maximum EPS, default is 10"
    print "\t-o (--only-once)\tOnly send for one second"
    print "\t-f (--facility)\tSyslog Facility to use"
    print "\t-p (--priority)\tSyslog priority to use"
    sys.exit()

def Version():
    print "MrLog Version:" + __version__

def ProcessArgs(args):
    global _MAX_LINES, _SERVERS,_USE_DATE, _EPS, _ONLYONCE, _FACILITY, _PRIORITY

    for opt, arg in args:
        if opt in ('-n', '--number-of-lines'):
            _MAX_LINES =  int(arg)
        elif opt in ('-h', '--help'):
            ShowHelp()
        elif opt in ('-v', '--version'):
            Version()
        elif opt in ('-s', '--server'):
            _SERVERS = [ server.strip() for server in arg.split(",") ]
        elif opt in ('-d', '--keep-date'):
            _USE_DATE = True
        elif opt in ('-e', '--eps'):
            _EPS = int(arg)
        elif opt in ('-o', '--only-once'):
            _ONLYONCE = True
        elif opt in ('-f', '--facility'):
            _FACILITY = arg
        elif opt in ('-p', '--priority'):
            _PRIORITY = arg


def ProcessFiles(files):
    global messages, _USE_DATE
    syslogMask = re.compile(r'^(?P<date>\w{3}\s{1,2}\d{1,2}\s\d{2}:\d{2}:\d{2})\s(?P<host>\S+)\s(?P<msg>.*)')

    for file in files:
        print "Loading lines from " + file

        currentFile = open(file)
        line = currentFile.readline()

        for x in range(1, _MAX_LINES):
            if not line:
                break

            matches  = syslogMask.search(line)
            if matches is not None:
                message = {}
                if _USE_DATE:
                    message["Date"] =  matches.group(1)
                else:
                    message["Date"] =  datetime.now().strftime('%b %d %H:%M:%S')

                message["Host"] = matches.group(2)
                message["Msg"] = matches.group(3)

                messages.append(message)

            line = currentFile.readline()

def StartLogging():

    global messages, _SERVERS, _EPS, _ONLYONCE

    smoothEPS = _EPS
    maxTime = 1

    # We need to smooth out the EPS, if we chuck all the logs as fast as we
    # can, we end up filling up the buffers and packets get dropped. So we
    # limit the output
    if _EPS > 100:
        smoothEPS = _EPS / 100
        maxTime = 0.01

    logger = netsyslog.Logger()

    for server in _SERVERS:
        host, port = server.split(":")

        if host:
            print "Adding host: " + host
            logger.add_host(host)
            if port:
                print "Adding port: " + port
                logger.PORT = int(port)
        else:
            logger.add_host(server)

    packetsSent = 0
    eps_time_start = time.time()
    while 1:
        smooth_time_start = time.time()

        messages_sent = 0
        for x in range(0, smoothEPS):
            messages_sent +=1
            message = messages[random.randrange(0, len(messages))]
            pri = netsyslog.PriPart(syslog.LOG_USER, syslog.LOG_INFO)
            header = netsyslog.HeaderPart(message["Date"], message["Host"])
            msg = netsyslog.MsgPart(tag="", content=message["Msg"])
            packet = netsyslog.Packet(pri, header, msg)
            logger.send_packet(packet)

        time_taken = time.time() - smooth_time_start
        if time_taken < maxTime:
            time.sleep(maxTime - time_taken)

        time_taken = time.time() - eps_time_start
        if time_taken > 1:
            if _EPS > 100:
                print "Current EPS=" + str(messages_sent * 100)
            else:
                print "Current EPS=" + str(messages_sent)

            eps_time_start = time.time()

        packetsSent += messages_sent
        if _ONLYONCE and packetsSent > _EPS:
            break


try:
    options, remainder = getopt.getopt(sys.argv[1:], 'n:hvs:de:of:p:', ['only-once', 'number-of-lines=',
                                                            'server=', 'help', 'version', 'keep-date', '--eps='])

    ProcessArgs(options)
    ProcessFiles(remainder)

    if len(messages) < 1:
        ShowHelp()

    StartLogging()

except getopt.GetoptError:
    ShowHelp()


