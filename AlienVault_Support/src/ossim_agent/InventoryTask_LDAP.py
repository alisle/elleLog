# -*- coding: utf-8 -*-
# Copyright (C) 2012 Alienvault
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
import ldap, ldif
import sys
import re
from cStringIO import StringIO
from InventoryTask import InventoryTask
from Event import HostInfoEvent
from Logger import Logger
#
# GLOBAL VARIABLES
#
logger = Logger.logger


class LDAP_TASK(InventoryTask):
    '''
    '''
    def __init__(self, task_name, task_params, task_period, task_reliability, task_enable, task_type,task_type_name):
        '''
        Constructor
        '''
        self._running = False
        self._validTask = True
        #ldaphost:192.168.12.200;ldapport:389;ldapuser:admin;ldappass:temporal;ldapdomain:alienvault.com;ldapbasedn:"ou=kktuaDevel,dc=testcfg,dc=qa,dc=alienvault,dc=com"
        self._pattern = re.compile("ldaphost:(?P<ldaphost>[^;]+);ldapport:(?P<ldapport>[^;]+);ldapuser:(?P<ldapuser>[^;]+);ldappass:(?P<ldappass>[^;]+);ldapdomain:(?P<ldapdomain>[^;]+);ldapbasedn:\"(?P<basedn>[^;]+)\"")
        values = self._pattern.match(task_params)
        self._ldapHost = ''
        self._ldapPort = ''
        self._ldapUser = ''
        self._ldapPass = ''
        self._ldapDomain = ''
        self._ldapBasedn = ''
        
        if values:
            groupdict = values.groupdict()
            self._ldapHost = groupdict['ldaphost']
            self._ldapPort = groupdict['ldapport']
            self._ldapUser = groupdict['ldapuser']
            self._ldapPass = groupdict['ldappass']
            self._ldapDomain = groupdict['ldapdomain']
            self._ldapBasedn = groupdict['basedn']
        else:
            logger.warning("Invalid ldap task")
            self._validTask = False
        self._ldapURL = 'ldap://%s:%s' % (self._ldapHost, self._ldapPort)
        self._ldapInstance = None
        InventoryTask.__init__(self, task_name, task_params, task_period, task_reliability, task_enable, task_type,task_type_name)


    def doJob(self):
        logger.info("Starting LDAP")
        try:
            self._ldapInstance = ldap.initialize(self._ldapURL)
            self._ldapInstance.simple_bind_s()
        except ldap.LDAPError, e:
            logger.error("Error creating LDAP instance: %s -  %s" % (self._ldapURL , str(e)))
            logger.info("Ending collector...")
            return
        logger.info("Connected to LDAP Server")
        try:
            data = self._ldapInstance.search_s(self._ldapBasedn, ldap.SCOPE_SUBTREE)
            organizationunit = ''
            for dn, entry in data:
                event = HostInfoEvent()
                if entry.has_key('ou'):
                    organizationunit = ','.join(["%s" % s for s in entry['ou']])
                    event['organization'] = organizationunit 
                if entry.has_key('cn'):
                    tmp = ','.join(["%s" % s for s in entry['cn']])
                    event['username'] = tmp
                if entry.has_key('mail'):
                    tmp = ','.join(["%s" % s for s in entry['mail']])
                    event['mail'] = tmp
                self.send_message(event)
            self._ldapInstance.unbind_s()
        except Exception, e:
            logger.error("Error running ldap query: %s" % str(e))
        self._running = False
        logger.info("LDAP collector ending..")
    def get_running(self):
        self._running
