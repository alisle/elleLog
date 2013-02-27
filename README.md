elleLog, Thoughts on SIEM
=======================

elleLog is an experimental POC SIEM designed to explore interesting aspects of SIEM design. elleLog is designed to be multi-threaded, scalable. 

It features a innovative plugin design, ElasticSearch back-end and built in Syslog server, as well as OSSIM Sensor support.

[Event Taxonomy](docs/event.taxonomy.md) discusses how the event taxonomy works.

[Plugin Functions](docs/plugin.functions.md) breaks down the functions which can be used within plugins.

[Tags](docs/plugin.tags.md) discusses the current standard set of tags which can be used within elleLog.

[OSSIM](docs/ossim.md) shows how to setup OSSIMs agent to send events to elleLog.
