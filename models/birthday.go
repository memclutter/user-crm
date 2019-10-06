package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type Birthday time.Time

func (b Birthday) Value() (driver.Value, error) {
	return time.Time(b), nil
}

func (b *Birthday) Scan(src interface{}) error {
	t, ok := src.(time.Time)
	if !ok {
		return errors.New("invalid database value for birthday type")
	}
	*b = Birthday(t)
	return nil
}

func (b *Birthday) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	*b = Birthday(t)
	return nil
}

func (b Birthday) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(b).Format("2006-01-02"))), nil
}
