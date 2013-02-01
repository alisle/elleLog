package ws

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "flag"
    "elleLog/elle/config"
    "elleLog/elle/processors"
    "text/template"
    "log"
    "encoding/json"
)

type Socket struct {
    ws *websocket.Conn
    queue chan Processors.Event
}
var ActiveSockets = make([]Socket, 0, 100)
var template_dir string
var static_dir string
var home* template.Template
var addr *string

func Initialize() {
    template_dir = Config.GlobalConfig.GetString(Config.SERVER_TEMPLATE_DIR, "lib/www/template")
    static_dir = Config.GlobalConfig.GetString(Config.SERVER_HTML_DIR, "lib/www/static")
    addr = flag.String("addr", Config.GlobalConfig.GetString(Config.SERVER_HTTP_ADDRESS, ":8080"), "http service address")
    home = template.Must(template.ParseFiles(template_dir + "/index.html"))
}

func homeHandler (c http.ResponseWriter, req *http.Request) {
    path := req.URL.Path
    log.Print("Path: " + path)

    if path == "/" {
        home.Execute(c, req.Host)
    } else {
        http.ServeFile(c, req, static_dir + path)
    }
}

func Process(event Processors.Event) {
    for _, connection := range ActiveSockets {
        connection.queue <- event
    }
}

func sendPacket(wsock *websocket.Conn, event Processors.Event) error {
    if jsonPacket, err := json.Marshal(event); err == nil {
        if err = websocket.Message.Send(wsock, string(jsonPacket)); err != nil {
            return err
        }
    } else {
        return err
    }
    
    return nil
}
func wsHandler(wsock *websocket.Conn) {
    log.Print("Added WS: ", wsock.RemoteAddr())
    sock := Socket{ wsock, make(chan Processors.Event) }
    ActiveSockets = append(ActiveSockets, sock)
    for event :=  range sock.queue {
        if err := sendPacket(wsock, event); err != nil {
            log.Print("Error: ", err)
            return 
        }
    }
}

func Launch() {
    go func() {
        log.Print("HTTP Server Listening on ",  (*addr))

        http.HandleFunc("/", homeHandler)
        http.Handle("/ws", websocket.Handler(wsHandler))

        if err:= http.ListenAndServe(*addr,nil); err != nil {
            log.Print("Unable to Create Websockets")
            return
        }
    }()
}

