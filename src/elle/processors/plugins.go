package Processors 
import (
	"os"
	"log"
	"elle/config"
    "elle/rfc3164"
	"path/filepath"
	"strings"
)


type Plugin struct {
	Name string
    Version string
	LineBegin string
	LineEnd string
	PairSep string
	MaxKey int
	Mapping map[string][]string
}

var Plugins = make(map[string]*Plugin)


type Event map[string]string

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

func CheckMessage(message *RFC3164.Message) {
/* Go through all the plugins against the message and figure out the percentage of hits */
    var bestEvent = Event {} 

    for _, plugin := range Plugins {
        event := processMessage(message, plugin)

        if len(event) > len(bestEvent) {
            bestEvent = event
        }
        event["Plugin"] = plugin.Name
    }
    dumpEvent(bestEvent)
}

func dumpEvent(event Event) {

    log.Print("New Event:")
    for key,value := range event {
        log.Print("\t", key, " : ", value)
    }

    log.Print("\n")
}

func processMessage(message *RFC3164.Message, plugin *Plugin) (Event) {

    var lines = [] string {}

    if plugin.LineEnd != "" {
        lines = strings.Split(message.Content, plugin.LineEnd)
    } else {
        lines = append(lines, message.Content)
    }

    var event = make(Event)

    for _, line := range lines {
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
                for _, key := range plugin.Mapping[forMapping] {
                    event[key] = fields[0]
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

                if _, ok := plugin.Mapping[currentKey]; ok {
                    forMapping = currentKey
                    takeField = true

                    break;
                }

                x--
            }
        }
    }
    return event
}
func New(fileName string) (*Plugin, error)  {
	plugin := Plugin{}

	config,err  := Config.New(fileName)
	if err != nil {
		log.Print("Failed to load plugin: ", fileName);
		return  nil, err 
	}

	plugin.Name, _ = config.GetFirstVariable("plugin", "name")
	plugin.Version, _ = config.GetFirstVariable("plugin", "version")
	plugin.LineBegin, _ = config.GetFirstVariable("seperators", "line_begin")
	plugin.LineEnd, _ = config.GetFirstVariable("seperators", "line_end")
	plugin.PairSep, _ = config.GetFirstVariable("seperators", "pair")
	plugin.Mapping =  map[string][]string{}

	mapping, _ := config.GetSection("mapping")

	for key, value := range  mapping {
		slice, ok := plugin.Mapping[value[0]]
		if !ok {
			slice =  make([]string, 0, 10)
		}

		spaces := strings.Count(value[0], " ") + 1
		if spaces > plugin.MaxKey {
			plugin.MaxKey = spaces
		}

		plugin.Mapping[value[0]] = append(slice, key)
	}

    log.Print(plugin.Mapping)
	return &plugin, nil
}





