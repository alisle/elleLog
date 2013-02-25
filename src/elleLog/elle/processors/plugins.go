package Processors 

// Imports
import (
    "os"
    "log"
    "errors"
    "elleLog/elle/config"
    "elleLog/elle/messages"
    "path/filepath"
    "strings"
    "regexp"
    "strconv"
    "encoding/base64"
)
// External Globals
var Plugins = make(map[string]*Plugin)
var EventsReceived = 0 
var LifetimeEventsReceived = 0

// Internal Globals
var messages chan *Messages.Message
var events chan Event

// Types
type Plugin struct {
    Name string
    Version string
    LineBegin string
    LineEnd string
    PairSep string
    MaxKey int
    Mapping map[string][]string
    Functions []functionInfo
    PositionMap map[int][]functionInfo
    KeyMap map[string][]functionInfo
    LitFunctions []functionInfo
}

type Event map[string]string

type fType int
const (
    PositionFunc fType  = iota + 1
    MapFunc 
    MapUntilFunc 
    RegexpFunc 
    SplitFunc 
    d64Func
    LitFunc

)
type functionDecl func (functionInfo, string) string

type functionInfo struct {
    FuncType fType
    Function functionDecl
    Arguments []string
    Tag string
}

var funcRegex = regexp.MustCompile("(?P<functionName>.*?)\\((?P<vars>.*)\\)")

func LoadAllPlugins(pluginDir string) {
    log.Print("Searching: ", pluginDir)
    filepath.Walk(pluginDir, func(path string, f os.FileInfo, err error) error {
        if !f.IsDir() && strings.HasSuffix(path, ".cfg") {
            if plugin, err := New(path); err != nil {
                log.Print("Error loading plugin: ", path)
            } else {
                log.Print("Loading ", plugin.Name, " Plugin")
                Plugins[plugin.Name] = plugin
            }
        }
        return nil
    })


}

func AttachMsgChannel(msgs chan *Messages.Message) {
    messages = msgs
}

func AttachEventsChannel(ents chan Event) {
    events = ents
}


func StartProcessing() {
    for x := 0; x < 20; x++ {
        go CheckMessage()
    }
}
func CheckMessage()  {
/* Go through all the plugins against the message and figure out the percentage of hits */
    for message := range messages {
        var bestEvent = Event {}

        for _, plugin := range Plugins {
            event := processMessage(message, plugin)
            
            if len(event) - len(plugin.LitFunctions) > len(bestEvent) {
                bestEvent = event
            }
            event["Plugin"] = plugin.Name
            event["raw_syslog"] = message.Content
            event["raw_from"] = message.IP
            event["raw_timestamp"] = message.TimeStamp
            event["raw_hostname"] = message.Hostname
        }

        if len(bestEvent) > 0  {
            events <- bestEvent
            EventsReceived++
            LifetimeEventsReceived++
        }
    }
}

func positionFunction(info functionInfo, context string) string {
    return context
}

func litFunction(info functionInfo, context string) string {
    log.Print("litFunction called, not good!")

    return ""
}

func processPositionFunctions(line string, plugin* Plugin, event Event) {
    if len(plugin.PositionMap) > 0 {
        var fields = strings.Fields(line)

        for key, functions := range plugin.PositionMap {
            if len(fields) > key {
                for _, value := range functions {
                    event[value.Tag] = value.Function(value, fields[key])
                }
            }
        }
    }
}

func regexFunction(info functionInfo, context string) string {
    if len(info.Arguments) != 2 {
        log.Printf("regexFunction: invalid number of arguments: needs 2, got %d", len(info.Arguments))
        return ""
    }

    regexKey := info.Arguments[0] 
    log.Print("Regexing " + regexKey)

    regexValue := info.Arguments[1]

    /* check if regexp actually compiles */
    compRegex, err := regexp.Compile(regexValue) 

    if err != nil {
        log.Printf("regexFunction: error compiling regexp: %s", err)
        return ""
    }

    result := compRegex.Find([]byte(context))
    if result != nil {
        log.Print("Match found: " + string(result))
        return string(result)
    } 

    return context
}
func d64Function(info functionInfo, context string) string {
    encoder := base64.StdEncoding
    maxLen := encoder.DecodedLen(len(context))

    buf := make([]byte, maxLen)

    if _, err := encoder.Decode(buf, []byte(context)); err != nil {
        log.Print("Failed to decode context")
        return context
    } 

    return string(buf)
}
func mapFunction(info functionInfo, context string) string {

    return context
}

func splitFunction(info functionInfo, context string) string {
    if len(info.Arguments) != 3 {
        log.Print("Error with Split Function")
        return ""
    }

    
    position, ok := strconv.ParseInt(info.Arguments[2], 0, 32)
    if ok != nil {
        log.Print("Split Function Invalid ", info.Arguments)
        return ""
    }

    fields := strings.Split(context, info.Arguments[1])

    if len(fields) < int(position) {
        return ""
    }

    return fields[position]
}

func processLitFunctions(plugin *Plugin, event Event) {
    for _, function := range plugin.LitFunctions {
        event[function.Tag] = function.Arguments[0]
    }
}

func processMessage(message *Messages.Message, plugin *Plugin) (Event) {
    var lines = [] string {}

    if plugin.LineEnd != "" {
        lines = strings.Split(message.Content, plugin.LineEnd)
    } else {
        lines = append(lines, message.Content)
    }

    var event = make(Event)
    
    
    /* Populate any literal Functions */
    processLitFunctions(plugin, event)

    for _, line := range lines {

        /* Process any Position Functions first */
        processPositionFunctions(line, plugin, event)


        /* We now have split on the pair, so for "ossec: Alert Level: 3"
        We now have: Cell[0] = ossec, Cell[1] = Alert Level, Cell[2] = 3
        */

        pivots := strings.Split(line, plugin.PairSep)

        /* Now we need to figure out how much of the cell makes up a key.
        For a cell containing "-  Source Network Address", the potential keys are:
        0 - "-  Source Network Address"
        1 - "Source Network Address"
        2 - "Network Address"
        3 - "Address"
        
        So we iterate through the cell until we hit one which we have a mapping for, we do this in reverse order.
        */

        // If we find a mapping the first cell in the next pivot will be the value we want.
        var takeField = false 
        var forMapping = ""

        for _, cell := range pivots {
            var fields = strings.Fields(cell)
            var fieldLen = len(fields)
            var currentKey = ""
            var currentKeyLength = 0
            
            if takeField && fieldLen > 0 {
                    word := fields[0]
                    if strings.HasPrefix(fields[0], "\"") {
                        for x := 1; x < len(fields); x++ {
                            word = word + " " + fields[x]
                            if strings.HasSuffix(fields[x], "\"") {
                                break
                            }
                        }
                    }
                    word = strings.Trim(word, "\" ")


                    for _, info := range plugin.KeyMap[forMapping] {
                        event[info.Tag] = info.Function(info, word)
                    }
                takeField = false
                forMapping = ""

            }

            x := fieldLen -1 
            for  x >= 0 {
                currentKeyLength += 1
                if currentKeyLength > plugin.MaxKey {
                    break;
                }
                /* THIS IS SLOW, change it to a buffer operation  or instead of the fields, using indexes and slices maybe? */
                if currentKey != "" {
                    currentKey = fields[x] + " " + currentKey
                } else {
                    currentKey = fields[x]
                }

                if _, ok := plugin.KeyMap[currentKey]; ok {
                    forMapping = currentKey
                    takeField = true
                }

                x--
            }
        }
    }
    return event
}

func createFunction(funcString string) (functionInfo, error) {
    f := functionInfo{}

    if matches := funcRegex.FindStringSubmatch(funcString); matches != nil {
        funcName := strings.ToLower(matches[1])
        variables := strings.Split(matches[2], ",")
        
        f.Arguments = make([]string, 0, 3)
        for _, variable := range variables {
            f.Arguments = append(f.Arguments, strings.Trim(variable, "\" "))
        }

        switch funcName {
            case "map": 
                f.FuncType = MapFunc
                f.Function = mapFunction
            case "split": 
                f.FuncType = SplitFunc
                f.Function = splitFunction
            case "regex": 
                f.FuncType = RegexpFunc
                f.Function = regexFunction
            case "pos":
                f.FuncType = PositionFunc
                f.Function = positionFunction
            case "d64":
                f.FuncType = d64Func
                f.Function = d64Function
            case "lit":
                f.FuncType = LitFunc
                f.Function = litFunction

            default:
                return functionInfo{}, errors.New("Invalid Function Type")
        }
        
        return f, nil
    }
     
    return functionInfo{}, errors.New("Invalid Function")
}
func New(fileName string) (*Plugin, error)  {
    plugin := Plugin{}

    config,err  := Config.New(fileName)
    if err != nil {
        log.Print("Failed to load Plugins: ", fileName);
        return  nil, err 
    }

    plugin.Name = config.GetString("plugin.name", "")
    plugin.Version = config.GetString("plugin.version", "0.1")
    plugin.LineBegin = config.GetString("seperators.line_begin", "")
    plugin.LineEnd = config.GetString("seperators.line_end", "")
    plugin.PairSep = config.GetString("seperators.pair", "")
    plugin.Mapping =  map[string][]string{}
    plugin.PositionMap = map[int][]functionInfo {}
    plugin.KeyMap = map[string][]functionInfo {}
    plugin.LitFunctions = make([]functionInfo, 0, 50)

    tags := config.GetMap("tags")
    if tags != nil {
        for key, value := range tags {
            for _, valueLit := range value {
                // Grab the function used
                if ret, err := createFunction(valueLit); err != nil {
                    log.Print("Unable to load Line: ",  key, "=", value, " ", err)
                } else {
                    ret.Tag = key

                    // If the function is PositionFunc, then we put it into a seperate list.
                    if ret.FuncType == PositionFunc {
                        if position, ok := strconv.ParseInt(ret.Arguments[0], 0, 64); ok == nil {
                            if _, ok := plugin.PositionMap[int(position)]; !ok {
                                plugin.PositionMap[int(position)] = make([]functionInfo, 0, 10)
                            }
                            plugin.PositionMap[int(position)] = append(plugin.PositionMap[int(position)], ret)

                        } else {
                            log.Print("Unable to load line: ", key, " ", value)
                        }
                    } else if ret.FuncType == LitFunc {
                        plugin.LitFunctions = append(plugin.LitFunctions, ret )
                    } else {
                        // Otherwise we take the first argument as the key needed for the value.
                        spaces := strings.Count(ret.Arguments[0], " ") + 1
                        if spaces > plugin.MaxKey {
                            plugin.MaxKey = spaces
                        }

                       if _, ok := plugin.KeyMap[ret.Arguments[0]]; !ok {
                            plugin.KeyMap[ret.Arguments[0]] = make([]functionInfo, 0, 10)
                        }

                        plugin.KeyMap[ret.Arguments[0]] = append(plugin.KeyMap[ret.Arguments[0]], ret)
                    }
                }
            }
        }
    }

    return &plugin, nil
}





