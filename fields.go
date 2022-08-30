package tallyGo

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Option struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Field struct {
	Key     string    `json:"key"`
	Label   string    `json:"label"`
	Type    FieldType `json:"type"`
	Value   []byte    `json:"value"`
	Options []Option  `json:"options"`
}

type FieldType string

const (
	TallyFieldType_HIDDEN          FieldType = "HIDDEN_FIELDS"
	TallyFieldType_CALCULATED      FieldType = "CALCULATED_FIELDS"
	TallyFieldType_TEXT            FieldType = "INPUT_TEXT"
	TallyFieldType_NUMBER          FieldType = "INPUT_NUMBER"
	TallyFieldType_PHONE           FieldType = "INPUT_PHONE_NUMBER"
	TallyFieldType_EMAIL           FieldType = "INPUT_EMAIL"
	TallyFieldType_LINK            FieldType = "INPUT_LINK"
	TallyFieldType_DATE            FieldType = "INPUT_DATE"
	TallyFieldType_TIME            FieldType = "INPUT_TIME"
	TallyFieldType_TEXTAREA        FieldType = "INPUT_TEXTAREA"
	TallyFieldType_MULTIPLE_CHOICE FieldType = "INPUT_MULTIPLE_CHOICE"
	TallyFieldType_CHECKBOXES      FieldType = "INPUT_CHECKBOXES"
	TallyFieldType_DROPDOWN        FieldType = "INPUT_DROPDOWN"
	TallyFieldType_FILE_UPLOAD     FieldType = "INPUT_FILE_UPLOAD"
	TallyFieldType_PAYMENT         FieldType = "INPUT_PAYMENT"
	TallyFieldType_RATING          FieldType = "INPUT_RATING"
	TallyFieldType_LINEAR_SCALE    FieldType = "INPUT_LINEAR_SCALE"
)

type HiddenFieldValue string

func (v *HiddenFieldValue) UnmarshalJSON(data []byte) error {
	*v = HiddenFieldValue(data)
	return nil
}

type CalculatedStringField string

func (v *CalculatedStringField) UnmarshalJSON(data []byte) error {
	*v = CalculatedStringField(data)
	return nil
}

type CalculatedNumberField float64
type TextField string
type NumberField float64
type PhoneField string
type EmailField string
type LinkField string
type DateField time.Time
type TimeField time.Time
type TextAreaField string
type MultipleChoiceField string
type CheckboxField string
type DropdownField string

type FileUploadFieldValue struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	MimeType string `json:"mimeType"`
	Size     int    `json:"size"`
}

func (v *FileUploadFieldValue) UnmarshallJSON(data []byte) error {
	// TODO
	return nil
}

type PaymentFieldValue struct {
	Price    float64
	Currency string
	Name     string
	Email    string
	Link     string
}

func (v *PaymentFieldValue) UnmarshallJSON(data []byte) error {
	// TODO
	return nil
}

//var TypeConstToField = map[FieldType]FieldValue{
//	TallyFieldType_HIDDEN: new(HiddenFieldValue),
//	TallyFieldType_CALCULATED:
//}

func (f Field) GetValue() (interface{}, error) {
	switch f.Type {

	// we know some fields are strings
	case TallyFieldType_HIDDEN, TallyFieldType_TEXT, TallyFieldType_PHONE, TallyFieldType_EMAIL, TallyFieldType_LINK, TallyFieldType_TEXTAREA:
		// todo: add the rest of the field types here
		s := new(string)
		err := json.Unmarshal(f.Value, s)
		return s, err
	case TallyFieldType_NUMBER:
		// todo: add rest of types
		n := new(float64)
		err := json.Unmarshal(f.Value, n)
		return n, err
	case TallyFieldType_MULTIPLE_CHOICE:
		s := new(string)
		err := json.Unmarshal(f.Value, s)
		if err != nil {
			return nil, err
		}
		for _, o := range f.Options {
			if o.ID == *s {
				return o.Text, nil
			}
		}
		return nil, errors.New(fmt.Sprintf("no option matching %v", s))
	case TallyFieldType_CALCULATED:
		// calculated fields can be string or number
		if string(f.Value[0]) == "'" {
			s := new(string)
			err := json.Unmarshal(f.Value, s)
			return s, err
		} else {
			n := new(float64)
			err := json.Unmarshal(f.Value, n)
			return n, err
		}

	}

	return nil, errors.New(fmt.Sprintf("field type %v not implemented", f.Type))
}

func (f Field) GetStringValue() *string {
	v, err := f.GetValue()
	if err != nil {
		return nil
	}

	if s, ok := v.(string); ok {
		return &s
	} else {
		return nil
	}
}
func (f Field) GetFloatValue() *float64 {
	v, err := f.GetValue()
	if err != nil {
		return nil
	}

	if s, ok := v.(float64); ok {
		return &s
	} else {
		return nil
	}
}
