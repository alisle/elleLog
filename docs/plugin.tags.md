Plugin Tags
============

Intro
-----------------
It's worth noting that currently elleLog doesn't enforce any limit on the tags which can be used
and any tag defined within a plugin will be used happily. This is a mixed blessing of course
and will be changed in the future. The current tag definitions (like mosti things in elleLog) 
have not been settled on and are subject to expansion.

If you have a suggestions please don't hesitate to drop an email

### Current Tags
```
applicataion_proto - Application Protocol (HTTP, POP, IMAP etc)

event_action - Action
event_priority
event_id 
event_message
event_type
event_taxonomy


device_address - The Device IP  which generated the action/event
device_hostname - The Device hostname which generated the action/event
device_mac - The Device MAC
device_inbound_interface - The Device interface which received the packet.
device_outbound_interface - The Device interface which send the packet.
device_serial_number

device_type
device_vendor
device_version

destination_address - The destination IP
destination_hostname - The destination hostname
destination_mac - The destination MAC
destination_NTDomain - The destination domain name.
destination_port - The destination port
destination_process - The name of the process (e.g. telnetd, sshd etc)
destination_userid - The destination user id
destination_username - The destination user name
destination_vpn - The destination VPN
destination_workstation - The destination workstation

logon_type - The logon type (e.g. Local, Networked)
authentication_type - The authentication type

file_name - The filename
file_size - The file size
file_create_time - The file creation time
file_hash - The file HASH
file_id - The files ID
file_modification_time - The file modification time
file_path - The file path
file_permission - The files permission

file_old_name - The filename
file_old_size - The file size
file_old_create_time - The file creation time
file_old_hash - The file HASH
file_old_id - The files ID
file_old_modification_time - The file modification time
file_old_path - The file path
file_old_permission - The files permission

firewall_rule - Firewall Rule Number

virus_name - Virus Name
virus_engine - AntiVirus Engine / Version

email_recipient - Email Recipient
email_sender - Email Sender
email_relay - Email Relay

wireless_ssid - Wireless SSID
wirless_channel - Wireless Channel
wireless_encryption - Encryption Type

bytes_in - The number of bytes received.
bytes_out - The number of bytes sent.

request_url - the URL requested
request_cookies - Cookies associated
request_method - the request Method (GET/POST)

source_address - The source IP
source_hostname - The source hostname
source_mac - The source MAC
source_NTDomain - The source domain name.
source_port - The source port.
source_process - The name of the process 
source_userid - The source user id
source_username - The source user name
source_vpn - The source VPN
source_workstation - The source workstation


customheader_0 - Custom field
customheader_1  - Custom field
customheader_2  - Custom field
customheader_3  - Custom field
customheader_4  - Custom field
customheader_5  - Custom field
customheader_6  - Custom field
customheader_7  - Custom field
customheader_8  - Custom field
customheader_9  - Custom field


customfield_0  - Custom field
customfield_1  - Custom field
customfield_2  - Custom field
customfield_3  - Custom field
customfield_4  - Custom field
customfield_5  - Custom field
customfield_6  - Custom field
customfield_7  - Custom field
customfield_8  - Custom field
customfield_9  - Custom field

date_string
date_epoch
```

These have been added for support for AV. Please don't use them outside of that context.
```
plugin_id 
plugin_sid
```
