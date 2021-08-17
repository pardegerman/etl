package singer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

// Entry in a Catalog
type Entry struct {
	TapStreamID string      `json:"tap_stream_id,omitempty"`
	TableName   string      `json:"table_name,omitempty"`
	Schema      *Schema     `json:"schema,omitempty"`
	Stream      string      `json:"stream,omitempty"`
	Metadata    []*Metadata `json:"metadata"`

	/*
		KeyProperties     []string `json:"key_properties,omitempty"`
		Table             string   `json:"table,omitempty"`
		StreamAlias       string   `json:"stream_alias,omitempty"`
	*/
}

// Catalog contains streams
type Catalog struct {
	Streams []*Entry `json:"streams,omitempty"`
}

// Reads a singer catalog from the provided reader
func ReadCatalog(r io.Reader) (c *Catalog, err error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	catalog := new(Catalog)
	err = json.Unmarshal(bytes, catalog)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}

// GetStream iterates through slice and returns Stream struct matching stream name
func (c *Catalog) GetStream(streamID string) (*Entry, error) {
	for _, s := range c.Streams {
		if s.TapStreamID == streamID {
			return s, nil
		}
	}

	errorMessage := fmt.Sprintf("Stream %s not found", streamID)
	return &Entry{}, errors.New(errorMessage)
}

// Dump outputs all the streams in the Catalog to JSON
func (c *Catalog) Dump() (string, error) {
	bs, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
