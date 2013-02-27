OSSSIM Support
===============


elleLog offers support for OSSIM Sensors. In order to enable this, please edit 
the  etc/ellelog.cfg, and add the line:

```
Listener.Attach.AVLogger=":4001"
```

After that on your OSSIM instance edit the /etc/ossim/agent/config.cfg and in the section:
```
[output-server]
enable=True
ip=192.168.0.10
port=40001
send_events=True
```

change the ip to the the IP address of your elleLog server.

That's it, it will now report to the elleLog Server.

Reporting to both OSSIM Server and elleLog
===========================================

This requires a change to the OSSIM agent, I will post the changes needed shortly.



