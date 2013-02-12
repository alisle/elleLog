Event Taxonomy
==============

Overview
--------
This taxonomy is designed for normalisation of the log events received by the SIEM. This allows 
searching for information and events based on content rather than device type. 

The taxonomy has been designed within a tree structure, with a naming convention which hasn’t 
been finalised. The structure allows for easy adding of event types without affecting the IDs
of the rest of the tree. This makes changing the taxonomy simple with effects been localised.
 
Tree
----
Each level of the tree is a event within it’s own right. So for example an event of type 
Attack.Network is as valid as Attack.Network.Back Orifice, please note when the naming convention 
has been completed this will look like ATT.2 and ATT.2.1 respectfully. 

This allows for as much detail to be captured within the taxonomy as possible, while not 
complicating correlation and searching of events.

Tags
----
Each node of the tree will have a set of tags which the SIEM will expect to be filled for the 
event. These tags will be inherited from the branches higher in the tree. While it is possible 
to define an event with no tags associated with it, these definitions are given as a strong 
set of guidelines for plugin development. Also this allows for the SIEM to guess on the 
taxonomy of an event which is received without a taxonomy defined within it.


Correlation Rules
-----------------
Correlation rules will be applied by taxonomy.  It will be possible to write a rule for ATT.2.1,
however it will also be possible to write a rule for ATT.2 and this rule will be applied to all
taxonomy events whos parent is ATT.2 for instance ATT.2.1, ATT.2.2 etc. 

This allows for a minimal set of rules to pick patterns out of the events while not 
sacrificing the amount of data being captured for each alarm.

Attacks (ATT)
==============
```
  ATT.1 Backdoor
     ATT.1.1 Runtime Detected
     ATT.1.2 Connection 
       ATT.1.2.1 Outbound
       ATT.1.2.2 Inbound
       ATT.1.2.3 Init
       ATT.1.2.4 Confirmation
       ATT.1.2.5 Response
       ATT.1.2.6 Attempt
    ATT.1.3 Get System Info
    ATT.1.4 Download
    ATT.1.5 Upload
    ATT.1.6 Retrieve Process List
    ATT.1.7 Process
       ATT.1.7.1 Start
       ATT.1.7.2 Stop
    ATT.1.8 KeyLogger
       ATT.1.8.1 Start
       ATT.1.8.2 Stop
  ATT.2 Black List
    ATT.2.1 DNS
    ATT.2.2 WEB
       ATT.2.2.1 URI
       ATT.2.2.2 User Agent
  ATT.3 Botnet
    ATT.3.1 Connection
       ATT.3.1.1 Attempt
       ATT.3.1.2 Outbound
       ATT.3.1.3 Inbound
       ATT.3.1.4 Server Attempt
       ATT.3.1.5 Client Attempt
    ATT.3.2 Credit Card submission
    ATT.3.3 Configuration attempt
    ATT.3.4 Download
    ATT.3.5 Upload
  ATT.4 DOS
  ATT.5 Exploit
    ATT.5.1 Local
    ATT.5.2 Remote
    ATT.5.3 Shell Code Detected
  ATT.6 Flood
    ATT.6.1 FIN Flood
    ATT.6.2 RST Flood
    ATT.6.3 SYN Flood
    ATT.6.4 Added Host to FIN Flood Blacklist
    ATT.6.5 Added Host to RST Flood Blacklist
    ATT.6.6 Added Host to SYN Flood Blacklist
    ATT.6.7 Removed Host from SYN Flood Blacklist
    ATT.6.8 Removed Host from RST Blacklist
    ATT.6.9 Removed Host from SYNC Blacklist
  ATT.7 Spoof
    ATT.7.1 DNS
    ATT.7.2 ARP
    ATT.7.3 IP
  ATT.8 are
    ATT.8.1 Toolbar
    ATT.8.2 Keylogger
    ATT.8.3 Pord Stealer
  ATT.9 Phishing
    ATT.9.1 SPAM
  ATT.10 Port Scan
    ATT.10.1 FIN Scan
    ATT.10.2 SYN Scan
```

State(STA)
===========
```
  STA.1 Anti-Virus
    STA.1.1 Host doesn’t have Anti-Virus installed
    STA.1.2 Host Anti-Virus out of date
    STA.1.3 Host Anti-Virus subscription has expired
  STA.2 Logging
    STA.2.1 Log Cleared
    STA.2.2 Log Sent
    STA.2.3 Error
       STA.2.3.1 Send Error
       STA.2.3.2 Log Full
    STA.2.4 Logging Set to INFO
    STA.2.5 Logging Set to DEBUG
  STA.3 Policy
    STA.3.1 URL List
       STA.3.1.1 Loaded
       STA.3.1.2 Not Loaded
       STA.3.1.3 Expired
       STA.3.1.4 Error
           STA.3.1.4.1 Filter Settings Incorrect
           STA.3.1.4.2 DNS Settings Incorrect
           STA.3.1.4.3 Device not Registered
           STA.3.1.4.4 Subscription Expired
           STA.3.1.4.5 Retry
           STA.3.1.4.6 Write Failure
    STA.3.2 Policy is up
    STA.3.3 Policy is down
    STA.3.4 Policy is unn
    STA.3.5 Policy has been Added
    STA.3.6 Policy has been Removed
    STA.3.7 Policy has been Modified
    STA.3.8 Access Rules
       STA.3.8.1 Added
       STA.3.8.2 Deleted
       STA.3.8.3 Modified
       STA.3.8.4 Restored to default


  STA.4 Devices
    STA.4.1 Time
       STA.4.1.1 Updated
           STA.4.1.1.1 Manually
           STA.4.1.1.2 Automatically
       STA.4.1.2 Zone
    STA.4.2 Temperature
       STA.4.2.1 Normal
       STA.4.2.2 Hot
       STA.4.2.3 Critical
       STA.4.2.4 Error
           STA.4.2.4.1 Failed to read temperature
    STA.4.3 Interface
       STA.4.3.1 Link
           STA.4.3.1.1 Up
           STA.4.3.1.2 Down
       STA.4.3.2 Ethernet
           STA.4.3.2.1 Port up
           STA.4.3.2.2 Port down
    STA.4.4 IP
       STA.4.4.1 Assigned
       STA.4.4.2 Released
       STA.4.4.3 Changed
           STA.4.4.3.1 ork interface changed for IP
       STA.4.4.4 Error
           STA.4.4.4.1 Conflict
           STA.4.4.4.2 IP/MAC Changing too often
           STA.4.4.4.3 Bad address length
    STA.4.5 Operating System
    STA.4.6 Availability
       STA.4.6.1 Host is Up
       STA.4.6.2 Host is Down
       STA.4.6.3 Host is Active
       STA.4.6.4 Host is Active
       STA.4.6.5 Host has been Deactivated
       STA.4.6.6 Host is Rebooting
           STA.4.6.6.1 Scheduled
           STA.4.6.6.2 Unexpected
           STA.4.6.6.3 By Administrator
           STA.4.6.6.4 By Service
           STA.4.6.6.5 Diagnostic
       STA.4.6.7 Host is Shutting Down
           STA.4.6.7.1 Scheduled
           STA.4.6.7.2 Unexpected
           STA.4.6.7.3 By Administrator
           STA.4.6.7.4 By Service
           STA.4.6.7.5 Diagnostic
    STA.4.7 Services
       STA.4.7.1 Enabled
       STA.4.7.2 Disabled
    STA.4.8 Error
       STA.4.8.1 License Expired


  STA.5 High Availability
    STA.5.1 Transitioned
       STA.5.1.1 Primary has transitioned to Active
       STA.5.1.2 Backup has transitioned to Active
       STA.5.1.3 Primary has transitioned to Idle
       STA.5.1.4 Backup has transitioned to Idle
    STA.5.2 Preempt
       STA.5.2.1 Backup being preempted by Primary
       STA.5.2.2 Primary preempting Backup
       STA.5.2.3 Backup detected Active, going Idle
    STA.5.3 Heartbeat
       STA.5.3.1 Synchronized
           STA.5.3.1.1 Backup has been successfully synchronized
           STA.5.3.1.2 Primary has been successfully synchronized
       STA.5.3.2 Missed
           STA.5.3.2.1 Primary missed backup heartbeat
           STA.5.3.2.2 Backup missed Primary heartbeat
       STA.5.3.3 Error
           STA.5.3.3.1 Primary received error from Backup
           STA.5.3.3.2 Backup received error from Primary
           STA.5.3.3.3 Primary received HB from wrong source
           STA.5.3.3.4 Backup received HB from wrong source
           STA.5.3.3.5 Unable to process HB packet
           STA.5.3.3.6 Unable to synchronize
           STA.5.3.3.7 Received HB from incompatible source
    STA.5.4 Discovery
       STA.5.4.1 Primary
           STA.5.4.1.1 Active 
           STA.5.4.1.2 Idle
       STA.5.4.2 Backup
           STA.5.4.2.1 Active
           STA.5.4.2.2 Idle
       STA.5.4.3 Error
           STA.5.4.3.1 Unable to find Primary
           STA.5.4.3.2 Unable to find Secondary
    STA.5.5 Message Received
       STA.5.5.1 Primary received reboot signal from Backup
       STA.5.5.2 Backup received reboot signal from Primary


  STA.6 Network
    STA.6.1 ARP
       STA.6.1.1 Request 
           STA.6.1.1.1 Received
           STA.6.1.1.2 Sent
       STA.6.1.2 Response
           STA.6.1.2.1 Received
           STA.6.1.2.2 Sent
       STA.6.1.3 Error
           STA.6.1.3.1 Timeout
    STA.6.2 BOOTP
       STA.6.2.1 Client
           STA.6.2.1.1 Response received
       STA.6.2.2 Server
           STA.6.2.2.1 Request received
    STA.6.3 DHCP
       STA.6.3.1 Client
           STA.6.3.1.1 Lease
              STA.6.3.1.1.1 Request lease
              STA.6.3.1.1.2 Received new IP Lease
              STA.6.3.1.1.3 Lease has expired
              STA.6.3.1.1.4 Lease dropped
              STA.6.3.1.1.5 Declining lease
           STA.6.3.1.2 Enabled
           STA.6.3.1.3 Received
              STA.6.3.1.3.1 ACK 
              STA.6.3.1.3.2 NACK
              STA.6.3.1.3.3 DISCOVER
              STA.6.3.1.3.4 OFFER
              STA.6.3.1.3.5 RELEASE
           STA.6.3.1.4 Error
              STA.6.3.1.4.1 Did not receive ACK
              STA.6.3.1.4.2 No DHCP server found
              STA.6.3.1.4.3 Packet malformed
              STA.6.3.1.4.4 Multiple DHCP Servers
       STA.6.3.2 Server
           STA.6.3.2.1 Scopes changed
           STA.6.3.2.2 Sanity check passed
           STA.6.3.2.3 Send OFFER
           STA.6.3.2.4 Error
              STA.6.3.2.4.1 Sanity check failed
              STA.6.3.2.4.2 IP conflict
    STA.6.4 DNS
       STA.6.4.1 DDNS
           STA.6.4.1.1 Association Added
           STA.6.4.1.2 Association Removed
           STA.6.4.1.3 Association Enabled
           STA.6.4.1.4 Association Disabled
           STA.6.4.1.5 Association Updated
       STA.6.4.2 Error
           STA.6.4.2.1 No valid DNS server
           STA.6.4.2.2 No response from DNS server
       STA.6.4.3 Zone Transfer
           STA.6.4.3.1 Aed
           STA.6.4.3.2 Denied
           STA.6.4.3.3 Requested
       STA.6.4.4 Version Request
       STA.6.4.5 Authors Request
       STA.6.4.6 Inverse Query
    STA.6.5 Finger
       STA.6.5.1 Search
       STA.6.5.2 Version
       STA.6.5.3 Query
    STA.6.6 ICMP
       STA.6.6.1 Address Mask
           STA.6.6.1.1 Reply
           STA.6.6.1.2 Request
       STA.6.6.2 Alternate Host Address
       STA.6.6.3 Destination Unreachable
       STA.6.6.4 Echo
       STA.6.6.5 Information
           STA.6.6.5.1 Request
           STA.6.6.5.2 Reply
       STA.6.6.6 PING
       STA.6.6.7 Redirect
       STA.6.6.8 Timestamp
       STA.6.6.9 Traceroute
    STA.6.7 IGMPv2
       STA.6.7.1 Host
           STA.6.7.1.1 Joined Group
           STA.6.7.1.2 Left Group
       STA.6.7.2 Membership report received
       STA.6.7.3 Query 
           STA.6.7.3.1 General
           STA.6.7.3.2 Membership
       STA.6.7.4 Packet
           STA.6.7.4.1 Dropped
              STA.6.7.4.1.1 Wrong checksum received
              STA.6.7.4.1.2 Decoding error
              STA.6.7.4.1.3 Not Handled
       STA.6.7.5 Timeout
    STA.6.8 IGMPv3
       STA.6.8.1 Host
           STA.6.8.1.1 Joined Group
           STA.6.8.1.2 Left Group
       STA.6.8.2 Membership report received
       STA.6.8.3 Query 
           STA.6.8.3.1 General
           STA.6.8.3.2 Membership
       STA.6.8.4 Packet Dropped
           STA.6.8.4.1 Wrong checksum received
           STA.6.8.4.2 Decoding error
           STA.6.8.4.3 Not Handled
       STA.6.8.5 Timeout
    STA.6.9 IMAP
       STA.6.9.1 List
       STA.6.9.2 Delete
       STA.6.9.3 Create
       STA.6.9.4 Examine
       STA.6.9.5 Fetch
       STA.6.9.6 Rename
       STA.6.9.7 Subscribe
       STA.6.9.8 Status
       STA.6.9.9 Unsubscribe
       STA.6.9.10 SSLv2
           STA.6.9.10.1 Client Hello
           STA.6.9.10.2 Server Hello
           STA.6.9.10.3 Error
              STA.6.9.10.3.1 Invalid Client Hello
              STA.6.9.10.3.2 Invalid Server Hello
       STA.6.9.11 SSLv3
           STA.6.9.11.1 Client Hello
           STA.6.9.11.2 Server Hello
           STA.6.9.11.3 Error
              STA.6.9.11.3.1 Invalid Version
              STA.6.9.11.3.2 Invalid Timestampd
    STA.6.10 IP
       STA.6.10.1 Error
           STA.6.10.1.1 Conflict with another ethernet address
           STA.6.10.1.2 Header checksum Failed
           STA.6.10.1.3 Malformed packet


    STA.6.11 Multicast Packet
       STA.6.11.1 Adding 
           STA.6.11.1.1 Interface
           STA.6.11.1.2 VPN
       STA.6.11.2 Removing
           STA.6.11.2.1 Interface
           STA.6.11.2.2 VPN
       STA.6.11.3 TCP
           STA.6.11.3.1 Accepted
           STA.6.11.3.2 Dropped
              STA.6.11.3.2.1 Wrong MAC address 
              STA.6.11.3.2.2 Invalid source IP
       STA.6.11.4 UDP
           STA.6.11.4.1 Accepted
           STA.6.11.4.2 Dropped
              STA.6.11.4.2.1 Wrong MAC address
              STA.6.11.4.2.2 Invalid source IP
       STA.6.11.5 Error
           STA.6.11.5.1 Max address limit reached
    STA.6.12 NAT
       STA.6.12.1 Enabled
       STA.6.12.2 Disabled
    STA.6.13 SMTP
       STA.6.13.1 Error
           STA.6.13.1.1 Authentication
           STA.6.13.1.2 Connection limit reached
    STA.6.14 Tunnel
       STA.6.14.1 L2TP
           STA.6.14.1.1 Connect
              STA.6.14.1.1.1 User
              STA.6.14.1.1.2 Sesson
           STA.6.14.1.2 Disconnect
              STA.6.14.1.2.1 Timeout
              STA.6.14.1.2.2 User
              STA.6.14.1.2.3 From client
              STA.6.14.1.2.4 Session
           STA.6.14.1.3 Error
              STA.6.14.1.3.1 Adding to IP pool fai*led.
              STA.6.14.1.3.2 Not ready
           STA.6.14.1.4 LCP
              STA.6.14.1.4.1 Up
              STA.6.14.1.4.2 Down


       STA.6.14.2 PPPoE
           STA.6.14.2.1 Authentication
              STA.6.14.2.1.1 Successful
                 STA.6.14.2.1.1.1 CHAP
              STA.6.14.2.1.1.2 PAP
              STA.6.14.2.1.2 Failed
              STA.6.14.2.1.2.1 CHAP
              STA.6.14.2.1.2.2 PAP
           STA.6.14.2.2 Disconnect
              STA.6.14.2.2.1 Timeout
              STA.6.14.2.2.2 From user
           STA.6.14.2.3 LCP
              STA.6.14.2.3.1 Up
              STA.6.14.2.3.2 Down
           STA.6.14.2.4 ork Connected
           STA.6.14.2.5 ork Disconnected
           STA.6.14.2.6 Link
              STA.6.14.2.6.1 Up
              STA.6.14.2.6.2 Down
              STA.6.14.2.6.3 Finished
           STA.6.14.2.7 Negotiation started
           STA.6.14.2.8 Error
              STA.6.14.2.8.1 Decode
              STA.6.14.2.8.2 Not ready
       STA.6.14.3 PPTP
           STA.6.14.3.1 Disconnect
              STA.6.14.3.1.1 Traffic timeout
       STA.6.14.4 IPSec
           STA.6.14.4.1 IKE
              STA.6.14.4.1.1 Accept 
              STA.6.14.4.1.1.1 Proposal
              STA.6.14.4.1.1.2 Peer lifetime
              STA.6.14.4.1.2 Error
              STA.6.14.4.1.2.1 Proposed ID mismatch
              STA.6.14.4.1.2.2 No match for remote ork address
              STA.6.14.4.1.2.3 Authentication method doesn’t match
              STA.6.14.4.1.2.4 DH group doesn’t match
              STA.6.14.4.1.2.5 Encryption algorithm doesn’t match
              STA.6.14.4.1.2.6 Encryption algorithm key length doesn’t match
              STA.6.14.4.1.2.7 Hash algorithm doesn’t match
              STA.6.14.4.1.2.8 XAUTH required but no user name given
              STA.6.14.4.1.2.9 XAUTH required but no user pord given
              STA.6.14.4.1.2.10 Proposal doesn’t match
              STA.6.14.4.1.2.11 Protocol mismatch
              STA.6.14.4.1.2.12 ID mismatch
              STA.6.14.4.1.2.13 IP compression algorithm doesn’t match
              STA.6.14.4.1.2.14 Timeout
              STA.6.14.4.1.2.15 AH
                 STA.6.14.4.1.2.15.1 Authentication algorithm doesn’t match
                 STA.6.14.4.1.2.15.2 Authentication key length doesn’t match
                 STA.6.14.4.1.2.15.3 Authentication key rounds doesn’t match
                 STA.6.14.4.1.2.15.4 Perfect ard secrecy mismatch
              STA.6.14.4.1.2.16 ESP
                 STA.6.14.4.1.2.16.1 Authentication algorithm doesn’t match
                 STA.6.14.4.1.2.16.2 Authentication key length doesn’t match
                 STA.6.14.4.1.2.16.3 Authentication key rounds doesn’t match
                 STA.6.14.4.1.2.16.4 Encryption algorithm doesn’t match
                 STA.6.14.4.1.2.16.5 Encryption key length doesn’t match
                 STA.6.14.4.1.2.16.6 Encryption key rounds doesn’t match
                 STA.6.14.4.1.2.16.7 Mode mismatch
                 STA.6.14.4.1.2.16.8 Perfect ard secrecy mismatch
              STA.6.14.4.1.3 SA
              STA.6.14.4.1.3.1 Add
              STA.6.14.4.1.3.2 Delete
              STA.6.14.4.1.3.3 Expired
           STA.6.14.4.2 IKEv2
              STA.6.14.4.2.1 Authentication 
              STA.6.14.4.2.1.1 Successful
              STA.6.14.4.2.1.2 Failed
              STA.6.14.4.2.2 Negotiation complete
              STA.6.14.4.2.3 Error
              STA.6.14.4.2.3.1 ID mismatch
              STA.6.14.4.2.3.2 Decrypt failed
              STA.6.14.4.2.3.3 Proposal doesn’t match
              STA.6.14.4.2.3.4 Invalid SPI size
              STA.6.14.4.2.3.5 Invalid state
              STA.6.14.4.2.3.6 Attribute not found
              STA.6.14.4.2.3.7 Payload processing error
              STA.6.14.4.2.3.8 Payload validation failed
              STA.6.14.4.2.3.9 Illegal SPI
              STA.6.14.4.2.3.10 SA Invalid
              STA.6.14.4.2.3.11 Connection Failed
              STA.6.14.4.2.3.12 Decryption failed
              STA.6.14.4.2.3.13 Negotiations failed
                 STA.6.14.4.2.3.13.1 Extra payloads present
                 STA.6.14.4.2.3.13.2 Invalid input state
                 STA.6.14.4.2.3.13.3 Invalid output state
                 STA.6.14.4.2.3.13.4 Missing required payloads
              STA.6.14.4.2.4 SA
              STA.6.14.4.2.4.1 Add
              STA.6.14.4.2.4.2 Delete
              STA.6.14.4.2.4.3 Expired


    STA.6.15 Web
       STA.6.15.1 Server
           STA.6.15.1.1 Alert
              STA.6.15.1.1.1 Modsecurity
           STA.6.15.1.2 Started
           STA.6.15.1.3 Shun
              STA.6.15.1.3.1 Expected
              STA.6.15.1.3.2 Unexpected
           STA.6.15.1.4 Error
              STA.6.15.1.4.1 Invalid URI
              STA.6.15.1.4.1.1 Filename too long
              STA.6.15.1.4.2 Multiple invalid URI
              STA.6.15.1.4.3 Not enough resources
    STA.6.16 WLAN
       STA.6.16.1 Enabled
       STA.6.16.2 Disabled
           STA.6.16.2.1 Unexpected
           STA.6.16.2.2 Scheduled
       STA.6.16.3 Reboot
       STA.6.16.4 Recover
       STA.6.16.5 Error
           STA.6.16.5.1 Sequence number out of order
           STA.6.16.5.2 Max concurrent users reached
           STA.6.16.5.3 Rogue Access Point Found
           STA.6.16.5.4 Received message from unn AP
       STA.6.16.6 Access Point
           STA.6.16.6.1 Classification
              STA.6.16.6.1.1 Created Rule
              STA.6.16.6.1.2 Enabled Rule
              STA.6.16.6.1.3 Disabled Rule
              STA.6.16.6.1.4 Deleted Rule
           STA.6.16.6.2 Added
           STA.6.16.6.3 Removed
           STA.6.16.6.4 Rebooting
           STA.6.16.6.5 Rebooted
           STA.6.16.6.6 Removed AP Tree
           STA.6.16.6.7 Removed AP Table
           STA.6.16.6.8 Virtual Access Point Enabled
           STA.6.16.6.9 Virtual Access Point Disabled
           STA.6.16.6.10 Virtual Access Point Removed
           STA.6.16.6.11 Access Point Enabled
           STA.6.16.6.12 Access Point Disabled
           STA.6.16.6.13 Access Point Removed
           STA.6.16.6.14 New Channel
           STA.6.16.6.15 Added Wired MAC
           STA.6.16.6.16 Node
              STA.6.16.6.16.1 Added
              STA.6.16.6.16.2 Removed
              STA.6.16.6.16.3 Modified
           STA.6.16.6.17 Error
              STA.6.16.6.17.1 AP is invalid
              STA.6.16.6.17.2 Virtual AP is invalid
              STA.6.16.6.17.3 Unable to assign Virtual AO
              STA.6.16.6.17.4 AP using wrong key
              STA.6.16.6.17.5 Unable to generate AP Key
              STA.6.16.6.17.6 Duplicate name
              STA.6.16.6.17.7 Unable to assign bridge
              STA.6.16.6.17.8 Group doesn’t exist
              STA.6.16.6.17.9 Unsecure AP
              STA.6.16.6.17.10 Did not boot
              STA.6.16.6.17.11 Incomplete configuration
           STA.6.16.6.18 ESSID
              STA.6.16.6.18.1 Added
              STA.6.16.6.18.2 Removed
              STA.6.16.6.18.3 Modified
              STA.6.16.6.18.4 Error
              STA.6.16.6.18.4.1 Duplicate
           STA.6.16.6.19 BSSID
              STA.6.16.6.19.1 Added
              STA.6.16.6.19.2 Removed
              STA.6.16.6.19.3 Modified
       STA.6.16.7 Error
           STA.6.16.7.1 Unable to find AP
           STA.6.16.7.2 Unable to find STA


       STA.6.16.8 Packet
           STA.6.16.8.1 Sent
              STA.6.16.8.1.1 AP Message
              STA.6.16.8.1.2 STA Message
           STA.6.16.8.2 Accepted
           STA.6.16.8.3 Dropped
              STA.6.16.8.3.1 Unsecure SAP Message
              STA.6.16.8.3.2 Unsecure AP Message Code




       STA.6.16.9 Security
           STA.6.16.9.1 Enabled
           STA.6.16.9.2 Disabled
           STA.6.16.9.3 MAC Filter
              STA.6.16.9.3.1 Enabled
              STA.6.16.9.3.2 Disabled


    STA.6.17 WAN
       STA.6.17.1 DOS Protection
           STA.6.17.1.1 Started
           STA.6.17.1.2 Stopped
       STA.6.17.2 Link
           STA.6.17.2.1 Up
           STA.6.17.2.2 Down
       STA.6.17.3 Error
           STA.6.17.3.1 Not ready
           STA.6.17.3.2 Node limit too many IP Addresses in use
    STA.6.18 SSO
       STA.6.18.1 Client
           STA.6.18.1.1 Error
              STA.6.18.1.1.1 Server return error
              STA.6.18.1.1.2 Domain name is too long
              STA.6.18.1.1.3 Username is too long
              STA.6.18.1.1.4 Server resolution failed
              STA.6.18.1.1.5 Server timeout
              STA.6.18.1.1.6 Probe Failed
       STA.6.18.2 Server
           STA.6.18.2.1 Server is up
           STA.6.18.2.2 Server is down
           STA.6.18.2.3 Error
              STA.6.18.2.3.1 Configuration error
    STA.6.19 RADIUS
       STA.6.19.1 Client
           STA.6.19.1.1 Error
              STA.6.19.1.1.1 Communication failure
              STA.6.19.1.1.2 Server timeout
       STA.6.19.2 Server
           STA.6.19.2.1 Error
              STA.6.19.2.1.1 Communication Problem
              STA.6.19.2.1.2 Configuration Error
    STA.6.20 LDAP
       STA.6.20.1 Added new member
       STA.6.20.2 Client
           STA.6.20.2.1 Using non-admin account
           STA.6.20.2.2 Authentication
              STA.6.20.2.2.1 Aed
              STA.6.20.2.2.2 Denied
STA.6.20.2.2.2.1Server doesn’t a CHAP
           STA.6.20.2.3 Error
              STA.6.20.2.3.1 Bind failed
              STA.6.20.2.3.2 Server certificate has wrong host name
              STA.6.20.2.3.3 Communication failure
              STA.6.20.2.3.4 Directory mismatch
              STA.6.20.2.3.5 Schema mismatch
              STA.6.20.2.3.6 Server certificate not valid
              STA.6.20.2.3.7 Server name resolution failed
              STA.6.20.2.3.8 Server timeout
              STA.6.20.2.3.9 Server not using TLS
       STA.6.20.3 Server
           STA.6.20.3.1 Server is up
           STA.6.20.3.2 Server is down
           STA.6.20.3.3 Member
              STA.6.20.3.3.1 Added
              STA.6.20.3.3.2 Removed
  STA.7 Cryptography
    STA.7.1 Servers
       STA.7.1.1 PKI
           STA.7.1.1.1 Added CA
           STA.7.1.1.2 Removed CA
           STA.7.1.1.3 Error
              STA.7.1.1.3.1 Untrusted CA
              STA.7.1.1.3.2 Duplicate local certificate
              STA.7.1.1.3.3 Duplicate local certificate name
              STA.7.1.1.3.4 Import failed
              STA.7.1.1.3.5 Invalid Certificate Format
              STA.7.1.1.3.6 Unable to verify certificate
              STA.7.1.1.3.7 Unable to verify certificate chain
              STA.7.1.1.3.8 Public-private key mismatch


       STA.7.1.2 CRL
           STA.7.1.2.1 Requesting CRL
           STA.7.1.2.2 CRL Loaded
           STA.7.1.2.3 Error
              STA.7.1.2.3.1 Failed to get CRL
              STA.7.1.2.3.2 Unable to process CRL
              STA.7.1.2.3.3 Bad CRL Format
              STA.7.1.2.3.4 Cannot connect to CRL Server
              STA.7.1.2.3.5 Expired
              STA.7.1.2.3.6 Missing
              STA.7.1.2.3.7 Validation failure for Root Certificate
       STA.7.1.3 OCSP
           STA.7.1.3.1 Received Response
           STA.7.1.3.2 Sending Request
           STA.7.1.3.3 Resolved Domain Name
           STA.7.1.3.4 Error
              STA.7.1.3.4.1 Response error
              STA.7.1.3.4.2 Send failed
              STA.7.1.3.4.3 Failed to Resolve DNS
       STA.7.1.4 SSL
           STA.7.1.4.1 Accept
              STA.7.1.4.1.1 Website in whitelist#
              STA.7.1.4.1.2 HTTPS via SSL2
           STA.7.1.4.2 Dropped
              STA.7.1.4.2.1 Website in blacklist
              STA.7.1.4.2.2 Untrusted CA
    STA.7.2 Key
       STA.7.2.1 Error
           STA.7.2.1.1 Unable to load key
           STA.7.2.1.2 Failed to decrypt
           STA.7.2.1.3 Private key does not match certificate
           STA.7.2.1.4 Preshared key mismatch
    STA.7.3 Certificate
       STA.7.3.1 Chain
           STA.7.3.1.1 Error
              STA.7.3.1.1.1 Not Complete
              STA.7.3.1.1.2 No Root
              STA.7.3.1.1.3 Circular
       STA.7.3.2 Error
           STA.7.3.2.1 Failed to load
           STA.7.3.2.2 Expired
           STA.7.3.2.3 Not yet valid
           STA.7.3.2.4 Certificate with invalid date
           STA.7.3.2.5 Revoked
           STA.7.3.2.6 Not Found
           STA.7.3.2.7 Bad signature
           STA.7.3.2.8 Corrupt
    STA.7.4 Cipher
       STA.7.4.1 Error
           STA.7.4.1.1 Failed to set cipher
           STA.7.4.1.2 Weak cipher being used


  STA.8 Cellular
    STA.8.1 3G Device Detected
    STA.8.2 Error
       STA.8.2.1 No SIM Detected
       STA.8.2.2 Data usage limit reached
```

Connections (CON)
==================
```
  CON.1 Connection
    CON.1.1 Aed
       CON.1.1.1 DNS
       CON.1.1.2 Unn
       CON.1.1.3 FTP
       CON.1.1.4 SMTP
       CON.1.1.5 TCP
    CON.1.2 Denied
       CON.1.2.1 DNS
           CON.1.2.1.1 Blacklist
           CON.1.2.1.2 Policy
       CON.1.2.2 Unn
       CON.1.2.3 FTP
           CON.1.2.3.1 Data connection from non-default port
           CON.1.2.3.2 Blacklist
           CON.1.2.3.3 Policy
       CON.1.2.4 STMP
           CON.1.2.4.1 Blacklist
           CON.1.2.4.2 Policy
       CON.1.2.5 TCP
           CON.1.2.5.1 Abort Received
           CON.1.2.5.2 Denied, from LAN
           CON.1.2.5.3 Reject Received
           CON.1.2.5.4 FIN Packet
           CON.1.2.5.5 Handshake violation
           CON.1.2.5.6 SYN/FIN Packet
           CON.1.2.5.7 Blacklist
           CON.1.2.5.8 Policy
  CON.2 Packets
    CON.2.1 Aed
       CON.2.1.1 Broadcast
       CON.2.1.2 Unn
       CON.2.1.3 UDP
       CON.2.1.4 TCP
       CON.2.1.5 NNTP
       CON.2.1.6 ICMP
       CON.2.1.7 IP
       CON.2.1.8 PPTP
       CON.2.1.9 IPSec
    CON.2.2 Dropped
       CON.2.2.1 Broadcast
       CON.2.2.2 ICMP
           CON.2.2.2.1 From LAN
           CON.2.2.2.2 No match
           CON.2.2.2.3 Policy
           CON.2.2.2.4 Blacklist
       CON.2.2.3 NNTP
           CON.2.2.3.1 Blacklist
           CON.2.2.3.2 Policy
       CON.2.2.4 IP
           CON.2.2.4.1 Connection Limit Reached
              CON.2.2.4.1.1 Source
              CON.2.2.4.1.2 Destination
           CON.2.2.4.2 Expired
           CON.2.2.4.3 Blacklist
           CON.2.2.4.4 Policy
       CON.2.2.5 Unn
           CON.2.2.5.1 Blacklist
           CON.2.2.5.2 Policy
       CON.2.2.6 UDP
           CON.2.2.6.1 Checksum Error
           CON.2.2.6.2 From LAN dropped
           CON.2.2.6.3 Blacklist
           CON.2.2.6.4 Policy
       CON.2.2.7 TCP
           CON.2.2.7.1 Duplicate 
           CON.2.2.7.2 Fragmented
           CON.2.2.7.3 Received on closing connection
           CON.2.2.7.4 Received on closed connection
           CON.2.2.7.5 Invalid ACK number
           CON.2.2.7.6 Invalid header length
           CON.2.2.7.7 Invalid MSS option length
           CON.2.2.7.8 Invalid option length
           CON.2.2.7.9 Invalid SACK option
           CON.2.2.7.10 Invalid SEQ number
           CON.2.2.7.11 Invalid source port
           CON.2.2.7.12 Invalid wi scale option length
           CON.2.2.7.13 Invalid wi scale option value
           CON.2.2.7.14 Non-permitted Option
           CON.2.2.7.15 Missing mandatory ACK flag
           CON.2.2.7.16 Missing mandatory SYN flag
           CON.2.2.7.17 Received with SYN flag on existing connection
           CON.2.2.7.18 Bad header
           CON.2.2.7.19 Invalid flag
           CON.2.2.7.20 SYN/FIN Packet
           CON.2.2.7.21 Checksum Error
           CON.2.2.7.22 Blacklist
           CON.2.2.7.23 Policy
       CON.2.2.8 PPTP
           CON.2.2.8.1 Blacklist
           CON.2.2.8.2 Policy
       CON.2.2.9 IPSec
           CON.2.2.9.1 Blacklist
           CON.2.2.9.2 Policy
           CON.2.2.9.3 Invalid Host
```
Accounts & Accounts Management (ACC)
=====================================
```
  ACC.1 Account
    ACC.1.1 Administrator
       ACC.1.1.1 Created
       ACC.1.1.2 Modified
       ACC.1.1.3 Removed
       ACC.1.1.4 Enabled
       ACC.1.1.5 Disabled
       ACC.1.1.6 Expired
    ACC.1.2 Normal User
       ACC.1.2.1 Created
       ACC.1.2.2 Modified
       ACC.1.2.3 Removed
       ACC.1.2.4 Enabled
       ACC.1.2.5 Disabled
       ACC.1.2.6 Expired
    ACC.1.3 Guest
       ACC.1.3.1 Created
       ACC.1.3.2 Modified
       ACC.1.3.3 Removed
       ACC.1.3.4 Enabled
       ACC.1.3.5 Disabled
       ACC.1.3.6 Expired


  ACC.2 Login
    ACC.2.1 Aed
       ACC.2.1.1 Administrator
       ACC.2.1.2 Normal User
       ACC.2.1.3 Guest
    ACC.2.2 Denied
       ACC.2.2.1 Administrator
           ACC.2.2.1.1 Bad Credentials
           ACC.2.2.1.2 Logins disabled 
           ACC.2.2.1.3 Already logged on
           ACC.2.2.1.4 Blocked
              ACC.2.2.1.4.1 Modsecurity
              ACC.2.2.1.4.2 Too many failed login attempts
              ACC.2.2.1.4.3 Policy
              ACC.2.2.1.4.4 From that location
              ACC.2.2.1.4.5 From that zone
              ACC.2.2.1.4.6 From that Interface
       ACC.2.2.2 Normal user
           ACC.2.2.2.1 Bad Credentials
           ACC.2.2.2.2 Disabled
           ACC.2.2.2.3 Already logged on
           ACC.2.2.2.4 Blocked
       ACC.2.2.3 Guest
           ACC.2.2.3.1 Bad Credentials
           ACC.2.2.3.2 Disabled
           ACC.2.2.3.3 Already logged on
           ACC.2.2.3.4 Blocked
              ACC.2.2.3.4.1 Modsecurity
              ACC.2.2.3.4.2 Too many failed login attempts
              ACC.2.2.3.4.3 Policy
              ACC.2.2.3.4.4 Pord expired
              ACC.2.2.3.4.5 From that location
              ACC.2.2.3.4.6 From that zone
              ACC.2.2.3.5.7 From that Interface
       ACC.2.2.4 Unn user


  ACC.3 Logout
    ACC.3.1 Administrator
       ACC.3.1.1 Inactivity
       ACC.3.1.2 Locked out
       ACC.3.1.3 Max session time
    ACC.3.2 Normal user
       ACC.3.2.1 Inactivity
       ACC.3.2.2 Locked out
       ACC.3.2.3 Max session time


  ACC.4 Access 
    ACC.4.1 Aed
       ACC.4.1.1 File
       ACC.4.1.2 Directory
       ACC.4.1.3 Index
       ACC.4.1.4 Cookie
       ACC.4.1.5 Connection
       ACC.4.1.6 Website
       ACC.4.1.7 Newsgroup
       ACC.4.1.8 ActiveX
       ACC.4.1.9 Java
    ACC.4.2 Denied
       ACC.4.2.1 File
       ACC.4.2.2 Directory
       ACC.4.2.3 Index
       ACC.4.2.4 Cookie
       ACC.4.2.5 Connection
       ACC.4.2.6 Website
       ACC.4.2.7 Newsgroup
       ACC.4.2.8 ActiveX
       ACC.4.2.9 Java
    ACC.4.3 Non-Existent 
       ACC.4.3.1 File
       ACC.4.3.2 Directory
       ACC.4.3.3 Index
       ACC.4.3.4 Cookie
       ACC.4.3.5 Connection
       ACC.4.3.6 Website
       ACC.4.3.7 Newsgroup
       ACC.4.3.8 ActiveX
       ACC.4.3.9 Java


  ACC.5 File and Directory
    ACC.5.1 File
       ACC.5.1.1 Created
       ACC.5.1.2 Modifed
       ACC.5.1.3 Deleted
       ACC.5.1.4 Permissions
           ACC.5.1.4.1 Owner
              ACC.5.1.4.1.1 Exec Set
              ACC.5.1.4.1.2 Exec Unset
              ACC.5.1.4.1.3 Read Set
              ACC.5.1.4.1.4 Read Unset
              ACC.5.1.4.1.5 Write Set
              ACC.5.1.4.1.6 Write Unset
           ACC.5.1.4.2 Group
              ACC.5.1.4.2.1 Exec Set
              ACC.5.1.4.2.2 Exec Unset
              ACC.5.1.4.2.3 Read Set
              ACC.5.1.4.2.4 Read Unset
              ACC.5.1.4.2.5 Write Set
              ACC.5.1.4.2.6 Write Unset
           ACC.5.1.4.3 Other
              ACC.5.1.4.3.1 Exec Set
              ACC.5.1.4.3.2 Exec Unset
              ACC.5.1.4.3.3 Read Set
              ACC.5.1.4.3.4 Read Unset
              ACC.5.1.4.3.5 Write Set
              ACC.5.1.4.3.6 Write Unset
    ACC.5.2 Directory
       ACC.5.2.1 Created
       ACC.5.2.2 Modifed
       ACC.5.2.3 Deleted
       ACC.5.2.4 Permissions
           ACC.5.2.4.1 Owner
              ACC.5.2.4.1.1 Exec Set
              ACC.5.2.4.1.2 Exec Unset
              ACC.5.2.4.1.3 Read Set
              ACC.5.2.4.1.4 Read Unset
              ACC.5.2.4.1.5 Write Set
              ACC.5.2.4.1.6 Write Unset
           ACC.5.2.4.2 Group
              ACC.5.2.4.2.1 Exec Set
              ACC.5.2.4.2.2 Exec Unset
              ACC.5.2.4.2.3 Read Set
              ACC.5.2.4.2.4 Read Unset
              ACC.5.2.4.2.5 Write Set
              ACC.5.2.4.2.6 Write Unset
           ACC.5.2.4.3 Other
              ACC.5.2.4.3.1 Exec Set
              ACC.5.2.4.3.2 Exec Unset
              ACC.5.2.4.3.3 Read Set
              ACC.5.2.4.3.4 Read Unset
              ACC.5.2.4.3.5 Write Set
              ACC.5.2.4.3.6 Write Unset
```
VOIP(VOI)
==========
```
  VOI.1 H.232
    VOI.1.1 Connect
    VOI.1.2 Setup
    VOI.1.3 Address
    VOI.1.4 End Session
    VOI.1.5 RAS
       VOI.1.5.1 Admission 
           VOI.1.5.1.1 Confirm
           VOI.1.5.1.2 Reject
       VOI.1.5.2 Bidth Reject
       VOI.1.5.3 Disengage
           VOI.1.5.3.1 Confirm
           VOI.1.5.3.2 Reject
       VOI.1.5.4 Gatekeeper Reject
       VOI.1.5.5 Location
           VOI.1.5.5.1 Confirm
           VOI.1.5.5.2 Reject
       VOI.1.5.6 Registration Reject
       VOI.1.5.7 Unregistration Reject
       VOI.1.5.8 Error
           VOI.1.5.8.1 Unn message
  VOI.2 SIP
    VOI.2.1 Request
    VOI.2.2 Response
  VOI.3 Call
    VOI.3.1 Connected
    VOI.3.2 Disconnected
  VOI.4 Endpoint 
    VOI.4.1 Added
    VOI.4.2 Removed
```
