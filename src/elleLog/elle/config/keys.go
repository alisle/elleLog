package Config

//Set the Buffer lengths
var MAX_QUEUE_PACKETS = "max.queue.packets"
var MAX_QUEUE_MESSAGES = "max.queue.messages"
var MAX_QUEUE_EVENTS = "max.queue.events"

var MAX_CPUS = "max.cpus"
var MAX_SUMMARY_TIME = "max.summarytime"

var OUTPUT_SHOWSTDOUT = "output.showstdout"
var OUTPUT_ATTACH_FILE = "output.attach.file"
var OUTPUT_ATTACH_ELASTISEARCH = "output.attach.elasticsearch"
var OUTPUT_ATTACH_STATSERVER = "output.attach.statsserver"

var LISTENER_ATTACH_AV_LOGGER = "listener.attach.avlogger"
var LISTENER_ATTACH_UDP = "listener.attach.udp"
var LISTENER_ATTACH_UNIX_STREAM = "listener.attach.unixstream"

var ELASTICSEARCH_BULK_ENABLE = "elasticsearch.bulk.enable"
var ELASTICSEARCH_BULK_MAX_ITEMS = "elasticsearch.bulk.max_items"
var ELASTICSEARCH_BULK_MAX_SECS = "elasticsearch.bulk.max_seconds"
var ELASTICSEARCH_MAX_CONNECTIONS = "elasticsearch.max_connections"

var MESSAGE_THREADS = "message.threads"

var SERVER_TCP_ADDRESS = "server.listener.tcp_address"
var SERVER_HTTP_ADDRESS = "server.listener.http_address"

var SERVER_TEMPLATE_DIR = "server.http_address.template_directory"
var SERVER_HTML_DIR = "server.http_address.html_directory"

var DEFAULT_CONFIG_FILE = "etc/ellelog.cfg"
var DEFAULT_PLUGIN_DIR = "etc/plugins/"
