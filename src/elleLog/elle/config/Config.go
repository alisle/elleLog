package Config

import (
	"log"
	"bufio"
	"os"
	"strings"
    "strconv"
)
var WorkingDirectory string

var GlobalConfig *Config

type ConfigElement struct {
    Value *[]string
    Keys map[string] ConfigElement
}

type  Config  struct {
	fileName string
	internalMap map[string](map[string][]string)
    Values ConfigElement
}


func New(fileName string) (*Config, error) {
    conf := &Config{fileName, nil, ConfigElement { nil, make(map[string]ConfigElement)}}
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

func (this *Config)GetFloat(key string, defaultValue float64) (float64) {
    value := this.GetString(key, "")

    if value == "" {
        return defaultValue 
    }

    valueFloat, err := strconv.ParseFloat(value, 64)
    if err != nil {
        log.Print("Unable to convert string: ", value)
    }

    return valueFloat
}
func (this *Config)GetInt(key string, defaultValue int) (int) {
    value := this.GetString(key, "")

    if value == "" {
        return defaultValue
    }

    valueInt, err := strconv.ParseInt(value, 0, 64)
    if err != nil {
        log.Print("Unable to convert string: ", value)
        return defaultValue
    }

    return int(valueInt)
}
func (this *Config)GetBool(key string, defaultValue bool) (bool) {
    value  := this.GetString(key, "")

    if value == "" {
        return defaultValue
    } 

    return strings.Contains(strings.ToLower(value), "true") || strings.Contains(value, "1")
}

func (this *Config)GetAllStrings(key string) ([]string) {
    var sections = strings.Split(key, ".")

    currentElement := this.Values
    for _, section := range sections {
        if _, ok := currentElement.Keys[section]; !ok {
            return nil
        }
        currentElement = currentElement.Keys[section]
    }
    return (*currentElement.Value)
}

func (this *Config)GetString(key string, defaultValue string) (string) {
    var allStrings = this.GetAllStrings(key)

    if allStrings == nil {
        return defaultValue
    } 
    return allStrings[0]
}
func (this* Config)GetMap(key string) (map[string][]string) {
    var sections = strings.Split(key, ".")
    currentElement := this.Values
    for _, section := range sections {
        if _, ok := currentElement.Keys[section]; !ok {
            return nil
        }
        currentElement = currentElement.Keys[section]
    }
    
    var returnMap = make(map[string][]string)

    for mapKey, mapValues := range currentElement.Keys {
        returnMap[mapKey] = make([]string, len(*mapValues.Value), cap(*mapValues.Value))
        copy(returnMap[mapKey], (*mapValues.Value))
    }


    return returnMap
}

func (this *Config)generate(fileName string) (error){

	configFile := fileName
	log.Print("Loading Config: ", configFile)
	file, err := os.Open(configFile)
	if err != nil {
		log.Print("Error reading file, aborting...", configFile)
		return err 
	}

	reader := bufio.NewReader(file)

	line, _, err :=  reader.ReadLine()
	for err == nil {
		confLine := strings.TrimSpace(string(line))
        keyvalue := strings.Index(confLine, "=")
        if strings.HasPrefix(confLine, "#") || keyvalue < 0  {
            line, _, err = reader.ReadLine()
            continue
        }
        key := confLine[0:keyvalue]
        value := strings.Trim(strings.TrimSpace(confLine[keyvalue + 1:]), "\"")

        var currentMap = this.Values
        var sections = strings.Split(key, ".")

        for _, section := range sections {
            section = strings.TrimSpace(section)
            if _, ok := currentMap.Keys[section]; !ok {
                var newValue = make([]string, 0, 10)
                currentMap.Keys[section] = ConfigElement { &newValue, make(map[string]ConfigElement)}
            }

            currentMap = currentMap.Keys[section]
        }
        
        *currentMap.Value = append((*currentMap.Value), value) 
		line, _, err = reader.ReadLine()
	}

	return nil
}

