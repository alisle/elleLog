package Processors 
import (
	"os"
	"log"
	"elle/config"
	"path/filepath"
	"strings"
)


type Plugin struct {
	Name string
	Tag string
	LineBegin string
	LineEnd string
	PairSep string
	MaxKey int
	Mapping map[string][]string
}

var Plugins = make(map[string]*Plugin)

func LoadAllPlugins(pluginDir string) {
	log.Print("Searching: ", pluginDir)
	filepath.Walk(pluginDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(path, ".cfg") {
			if plugin, err := New(path); err != nil {
				log.Print("Error loading plugin: ", path)
			} else {
				Plugins[plugin.Tag] = plugin
			}
		}
		return nil
	})


}


func New(fileName string) (*Plugin, error)  {
	plugin := Plugin{}

	config,err  := Config.New(fileName)
	if err != nil {
		log.Print("Failed to load plugin: ", fileName);
		return  nil, err 
	}

	plugin.Name, _ = config.GetFirstVariable("plugin", "name")
	plugin.Tag, _ = config.GetFirstVariable("plugin", "tag")
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

	return &plugin, nil
}




