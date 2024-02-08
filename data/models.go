package data

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Attribute represents an attribute
type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// UserTrait represents a user trait
type UserTrait struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// EventData represents event data
type EventData struct {
	EventName       string               `json:"event"`
	EventType       string               `json:"event_type"`
	AppID           string               `json:"app_id"`
	UserID          string               `json:"user_id"`
	MessageID       string               `json:"message_id"`
	PageTitle       string               `json:"page_title"`
	PageURL         string               `json:"page_url"`
	BrowserLanguage string               `json:"browser_language"`
	ScreenSize      string               `json:"screen_size"`
	Attributes      map[string]Attribute `json:"attributes,omitempty"`
	UserTraits      map[string]UserTrait `json:"traits,omitempty"`
}

var mutex sync.Mutex

// ProcessRequest processes a request and returns EventData
func ProcessRequest(dynamicMap map[string]interface{}) EventData {
	mutex.Lock()
	defer mutex.Unlock()

	eventData := EventData{
		EventName:       getString(dynamicMap, "ev"),
		EventType:       getString(dynamicMap, "et"),
		AppID:           getString(dynamicMap, "id"),
		UserID:          getString(dynamicMap, "uid"),
		MessageID:       getString(dynamicMap, "mid"),
		PageTitle:       getString(dynamicMap, "t"),
		PageURL:         getString(dynamicMap, "p"),
		BrowserLanguage: getString(dynamicMap, "l"),
		ScreenSize:      getString(dynamicMap, "cs"),
		Attributes:      make(map[string]Attribute),
		UserTraits:      make(map[string]UserTrait),
	}

	populateDynamicFields(dynamicMap, "atrk", "atrv", "atrt", func(index string) {
		key := dynamicMap["atrk"+index].(string)
		value := dynamicMap["atrv"+index].(string)
		attrType := dynamicMap["atrt"+index].(string)
		eventData.Attributes[key] = Attribute{Value: value, Type: attrType}
	})

	populateDynamicFields(dynamicMap, "uatrk", "uatrv", "uatrt", func(index string) {
		key := dynamicMap["uatrk"+index].(string)
		value := dynamicMap["uatrv"+index].(string)
		traitType := dynamicMap["uatrt"+index].(string)
		eventData.UserTraits[key] = UserTrait{Value: value, Type: traitType}
	})

	return eventData
}

// getString returns a string value from a map
func getString(data map[string]interface{}, key string) string {
	if value, ok := data[key].(string); ok {
		return value
	}
	return ""
}

// populateDynamicFields populates dynamic fields based on prefixes
func populateDynamicFields(data map[string]interface{}, keyPrefix, valuePrefix, typePrefix string, callback func(index string)) {
	keys := getDynamicKeys(data, keyPrefix)
	sort.Strings(keys)
	for _, index := range keys {
		callback(index)
	}
}

// getDynamicKeys returns dynamic keys based on a prefix
func getDynamicKeys(data map[string]interface{}, prefix string) []string {
	var keys []string
	for key := range data {
		if strings.HasPrefix(key, prefix) {
			index := key[len(prefix):]
			if _, err := strconv.Atoi(index); err == nil {
				keys = append(keys, index)
			}
		}
	}
	return keys
}
