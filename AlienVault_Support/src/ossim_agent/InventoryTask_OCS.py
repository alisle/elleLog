# -*- coding: utf-8 -*-
# Copyright (C) 2012 Cristobal Rosa Galisteo
#
# This file is part of AVInventory.
#
# AVInventory is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 2 of the License, or
# (at your option) any later version.
#
# VLMa is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with VLMa. If not, see <http://www.gnu.org/licenses/>.
import time
import re
import socket
from xml.dom import minidom, Node

from InventoryTask import InventoryTask
from Event import HostInfoEvent
from Logger import Logger

#
# GLOBAL VARIABLES
#
logger = Logger.logger

class OCS_TASK(InventoryTask):
    def __init__(self, task_name, task_params, task_period, task_reliability, task_enable, task_type,task_type_name, fmkip, fmkport):
        '''
        Constructor
        '''
        self._running = False
        self._fmkSocket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self._framework_ip = fmkip
        self._framework_port = fmkport
        #self._fmkSocket.connect((self._framework_ip, int(self._framework_port)))
        InventoryTask.__init__(self, task_name, task_params, task_period, task_reliability, task_enable, task_type,task_type_name)
        
    def getOCSInventory(self):
        data = []
        try:
            logger.info("Connecting to fmkd %s:%s" % (self._framework_ip, int(self._framework_port)))
            self._fmkSocket.connect((self._framework_ip, int(self._framework_port)))
            self._fmkSocket.send("control action=\"getOCSInventory\"\n")
            continue_reading = True
            data =''
            while continue_reading:
                char=self._fmkSocket.recv(1)
                data+=char
                if char == '\n':
                    continue_reading = False
            logger.info("Connected..")
            self._fmkSocket.close()
        except:
            logger.warning("Error retrieving ocs inventory data..")
        return data
    def getText(self, nodelist):
        rc = ""
        for node in nodelist:
            if node.nodeType == node.TEXT_NODE:
                rc = rc + node.data
        return rc
    def doJob(self):
        self._running = True
        logger.info("OCS Process")
        data = self.getOCSInventory()
        if not data or data=='':
            return
        xml = re.findall("ocsinventory=\"([^\"]+)\"", data)
        if len(xml)>0:
            try:
                dom = minidom.parseString(xml[0])
            except Exception, e:
                logger.warning("OCS: invalid data:%s,%s" % (xml[0], str(e)))
            else:
                for host in  dom.getElementsByTagName('host'):
                    hostData = HostInfoEvent()
                    tmp = host.getElementsByTagName('ip')
                    if tmp and len(tmp) == 1:
                        tmp = tmp[0]
                        hostData['ip'] = self.getText(tmp.childNodes)
                    tmp = host.getElementsByTagName('hostname')
                    if tmp and len(tmp) == 1:
                        tmp = tmp[0]
                        hostData['hostname'] = self.getText(tmp.childNodes)
                    tmp = host.getElementsByTagName('mac')
                    if tmp and len(tmp) == 1:
                        tmp = tmp[0]
                        hostData['mac'] = self.getText(tmp.childNodes)
                    tmp = host.getElementsByTagName('os')
                    if tmp and len(tmp) == 1:
                        tmp = tmp[0]
                        hostData['os'] = self.getText(tmp.childNodes)
                    tmp = host.getElementsByTagName('video')
                    if tmp and len(tmp) == 1:
                        tmp = tmp[0]
                        hostData['video'] = self.getText(tmp.childNodes)
                    tmp = host.getElementsByTagName('memory')
                    if tmp and len(tmp) == 1:
                        tmp = tmp[0]
                        hostData['memory'] = self.getText(tmp.childNodes)
                    tmp = host.getElementsByTagName('video')
                    if tmp and len(tmp) == 1:
                        tmp = tmp[0]
                        hostData['video'] = self.getText(tmp.childNodes)
                    #tmp = host.getElementsByTagName('domain')
                    #if tmp and len(tmp) == 1:
                    #    tmp = tmp[0]
                    #    hostData['domain'] = self.getText(tmp.childNodes)
                    self.send_message(hostData)
        self._running = False
        logger.info("End ocs inventory job")
    def get_running(self):
        return self._running
