package configuration

import (

	"reflect"
	"strings"
	"errors"

	"gopkg.in/yaml.v2"
	"github.com/google/logger"
)

const stringTag = "string"

type Parameter struct {
	name      string      `string:"name"`
	unique    bool        `string:"unique"`
	required  bool        `string:"required"`
	contType  contentType `string:"contentType"`
	longdesc  string      `string:"longdesc,omitempty"`
	shortdesc string      `string:"shrotdesc,omitempty"`
}

type parameterFieldNum int

/*
	This constant pool should be updated together
	with parameter fields
 */
const (
	pfn_name      parameterFieldNum = iota
	pfn_unique
	pfn_required
	pfn_contType
	pfn_longdesc
	pfn_shortdesc
)

func (p *Parameter) UnmarshalYAML(b []byte) error {
	param := make(map[string]interface{})
	err := yaml.Unmarshal(b, &param)
	if err != nil {
		logger.Error("Can't unmarshall parameters from yaml, err: %s", err.Error())
		return err
	}

	prmType := reflect.TypeOf(Parameter{})
	for i := 0; i < prmType.NumField(); i++ {
		str := prmType.Field(i).Tag.Get(stringTag)
		strs := strings.Split(str, ",")
		if len(strs) == 0 {
			logger.Warning("WARNING: there is no tag for parameter field")
			continue
		}

		getParam := func() (interface{}, error) {
			val, ok := param[strs[0]]
			// if there is no such value and no ommitempty
			if !ok && len(strs) == 1 {
				logger.Error("Can't get necessary field %s from YAML", strs[0])
				return nil, errors.New("failed to unmarshal yaml")
			}
			return val, nil
		}

		switch parameterFieldNum(i) {
		case pfn_name:
			val, err := getParam()
			if err != nil {
				return err
			}
			if val == nil {
				continue
			}
			p.name = val.(string)
		case pfn_unique:
			val, err := getParam()
			if err != nil {
				return err
			}
			if val == nil {
				continue
			}
			p.unique = val.(bool)
		case pfn_required:
			val, err := getParam()
			if err != nil {
				return err
			}
			if val == nil {
				continue
			}
			p.required = val.(bool)
		case pfn_contType:
			val, err := getParam()
			if err != nil {
				return err
			}
			if val == nil {
				continue
			}
			p.contType = val.(contentType)
		case pfn_longdesc:
			val, err := getParam()
			if err != nil {
				return err
			}
			if val == nil {
				continue
			}
			p.longdesc = val.(string)
		case pfn_shortdesc:
			val, err := getParam()
			if err != nil {
				return err
			}
			if val == nil {
				continue
			}
			p.shortdesc = val.(string)
		default:
			logger.Warning("WARNING: %s field haven't been prepared to deserialisation from YAML", strs[0])
		}
	}

	return nil
}
