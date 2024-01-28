package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"
)

var defaultBlackList = []string{"password", "authorization", "token", "name", "lastname", "email", "card"}

func NewFormatter(fields ...string) logrus.Formatter {
	// defining formatter
	customFormatter := &logrus.TextFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:   "data",
			logrus.FieldKeyLevel: "log_type",
			logrus.FieldKeyTime:  "date",
		},
	}

	// initializing blacklist map
	blacklist := make(map[string]bool)
	for _, bl := range defaultBlackList {
		blacklist[bl] = true
	}

	for _, field := range fields {
		blacklist[strings.TrimSpace(strings.ToLower(field))] = true
	}

	return &jsonMaskSensitiveDataFormatter{
		TextFormatter: customFormatter,
		blacklist:     blacklist,
	}
}

type Mask interface {
	Format(e *logrus.Entry) ([]byte, error)
}

type jsonMaskSensitiveDataFormatter struct {
	*logrus.TextFormatter
	blacklist map[string]bool
}

func (f *jsonMaskSensitiveDataFormatter) Format(e *logrus.Entry) ([]byte, error) {
	result := make(map[string]any)

	// time is date
	result["time"] = e.Time.Format(time.RFC3339)
	// logger level
	result["log_level"] = e.Level.String()
	result["data"] = e.Data

	// set message
	msg, err := getMessage(e.Message)
	if err != nil {
		return nil, err
	}
	result["message"] = msg

	err = f.maskFields(result)
	if err != nil {
		return nil, fmt.Errorf("failed to mask blacklisted fields, %w", err)
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json with masked values, %w", err)
	}
	bytes = append(bytes, '\n')
	return bytes, nil
}

func (f *jsonMaskSensitiveDataFormatter) maskFields(fields map[string]any) error {
	for k, v := range fields {
		if v == nil {
			continue
		}

		vv := reflect.ValueOf(v)
		if vv.Kind() == reflect.Map {
			innerMap, err := toMap(vv.Interface())
			if err != nil {
				return err
			}
			if err = f.maskFields(innerMap); err != nil {
				return err
			}
			fields[k] = innerMap
		} else {
			value := vv.Interface()
			if f.isInBlackList(k) {
				value = strings.Repeat("*", utf8.RuneCountInString(fmt.Sprintf("%v", value)))
			}
			fields[k] = value
		}
	}

	return nil
}

func (f *jsonMaskSensitiveDataFormatter) isInBlackList(key string) bool {
	_, ok := f.blacklist[strings.TrimSpace(strings.ToLower(key))]
	return ok
}

func toMap(i any) (map[string]any, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	aux := make(map[string]any)
	err = json.Unmarshal(data, &aux)
	if err != nil {
		return nil, err
	}

	return aux, nil
}

func getMessage(msg string) (any, error) {
	var result any = msg
	if isJSON(msg) {
		err := json.Unmarshal([]byte(msg), &result)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal the message, %w", err)
		}
	}

	return result, nil
}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
