package Config

import (
	"log"
	"bufio"
	"os"
	"strings"
	"regexp"
)
var WorkingDirectory string


type ConfigElement struct {
    Value string
    Keys map[string] ConfigElement
}

type  Config  struct {
	fileName string
	internalMap map[string](map[string][]string)
}


func New(fileName string) (*Config, error) {
	conf := &Config{fileName, nil}
	err := conf.generate(fileName)

	return conf, err
}

func (this *Config)GetSection(section string) (map[string][]string, bool) {
	sectionMap, ok := this.internalMap[section]
	return sectionMap, ok
}

func (this *Config)GetFirstVariable(section string, variable string) (string, bool) {
	if sectionMap, isOK := this.GetSection(section); isOK {
		if variableSlice, ok := sectionMap[variable]; ok {
			return variableSlice[0], ok
		}
	}

	return "", false;
}
func (this *Config)GetVariable(section string, variable string) ([]string, bool) {
	if sectionMap, isOK := this.GetSection(section); isOK {
		variableSlice, ok := sectionMap[variable];

		return variableSlice, ok
	}
	return nil, false;
}
func (this *Config)generate(fileName string) error {
	this.internalMap = make(map[string](map[string][]string))

	confRegex := regexp.MustCompile("^(?P<section>[^\\.]+)\\.(?P<var>[^=]+)=(?P<value>.*)")
	configFile := fileName
	log.Print("Loading : ", configFile)
	file, err := os.Open(configFile)
	if err != nil {
		log.Print("Error reading file, aborting...", configFile)
		return err 
	}

	reader := bufio.NewReader(file)

	line, _, err :=  reader.ReadLine()
	for err == nil {
		confLine := string(line)

		if len(confLine) > 2 {
			if matches := confRegex.FindStringSubmatch(confLine); matches != nil {
				section := strings.ToLower(matches[1])
				variable := strings.ToLower(matches[2])
				value := strings.Trim(matches[3], "\" ")
				mapSection, ok :=  this.internalMap[section]
				if !ok {
					mapSection = make(map[string][]string)
					this.internalMap[section] = mapSection
				}

				variableSection, ok := mapSection[variable]
				if !ok {
					variableSection = make([]string, 0, 10)
				}
				mapSection[variable] = append(variableSection, value)
			}
		}

		line, _, err = reader.ReadLine()
	}
	return nil
}

func GenerateConfig2(fileName string) (error){
    var config ConfigElement

	configFile := fileName
	log.Print("Loading : ", configFile)
	file, err := os.Open(configFile)
	if err != nil {
		log.Print("Error reading file, aborting...", configFile)
		return err 
	}

	reader := bufio.NewReader(file)

	line, _, err :=  reader.ReadLine()
	for err == nil {
		confLine := string(line)
        keyvalue := strings.Split(confLine, "=")
        key := keyvalue[0]
        value := keyvalue[1]

        totalSections := len(key) - 1
        countSections := 0

        var currentMap = config 
        var sections = strings.Split(key, ".")

        for _, section := range sections {
            if countSections < totalSections {
                if _, ok := currentMap.Keys[section]; !ok {
                    currentMap.Keys[section] = ConfigElement { "", make(map[string]ConfigElement)}
                }

                currentMap = currentMap.Keys[section]
            } else {
                currentMap.Keys[section] = ConfigElement{ value, nil}
            }

            countSections++
        }

		line, _, err = reader.ReadLine()
	}
	return nil
}

