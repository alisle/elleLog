package Config

//Set the Buffer lengths
var MAX_QUEUE_PACKETS = "max.queue.packets"
var MAX_QUEUE_MESSAGES = "max.queue.messages"
var MAX_QUEUE_EVENTS = "max.queue.events"

var MAX_CPUS = "max.cpus"
var MAX_SUMMARY_TIME = "max.summarytime"

var OUTPUT_SHOWSTDOUT = "output.showstdout"
var OUTPUT_ATTACH_FILE = "output.attachfile"
var OUTPUT_ATTACH_ELASTISEARCH = "output.attachelastisearch"
var OUTPUT_ATTACH_STATSERVER = "output.attachstatsserver"

var LISTENER_ATTACH_UDP = "listener.attachudp"
var LISTENER_ATTACH_UNIX_STREAM = "listener.attachunixstream"

var ELASTICSEARCH_BULK_ENABLE = "elasticsearch.bulk.enable"
var ELASTICSEARCH_BULK_MAX_ITEMS = "elasticsearch.bulk.max_items"
var ELASTICSEARCH_BULK_MAX_SECS = "elasticsearch.bulk.max_seconds"
var ELASTICSEARCH_MAX_CONNECTIONS = "elasticsearch.max_connections"

var RFC3164_THREADS = "rfc3164.threads"


var SERVER_TCP_ADDRESS = "bind.tcp.address"
