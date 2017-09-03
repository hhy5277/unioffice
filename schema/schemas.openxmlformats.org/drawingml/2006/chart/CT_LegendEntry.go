// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_LegendEntry struct {
	Idx    *CT_UnsignedInt
	Choice *CT_LegendEntryChoice
	ExtLst *CT_ExtensionList
}

func NewCT_LegendEntry() *CT_LegendEntry {
	ret := &CT_LegendEntry{}
	ret.Idx = NewCT_UnsignedInt()
	return ret
}

func (m *CT_LegendEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seidx := xml.StartElement{Name: xml.Name{Local: "idx"}}
	e.EncodeElement(m.Idx, seidx)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LegendEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
lCT_LegendEntry:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "idx":
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case "delete":
				m.Choice = NewCT_LegendEntryChoice()
				if err := d.DecodeElement(&m.Choice.Delete, &el); err != nil {
					return err
				}
			case "txPr":
				m.Choice = NewCT_LegendEntryChoice()
				if err := d.DecodeElement(&m.Choice.TxPr, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_LegendEntry %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LegendEntry
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LegendEntry and its children
func (m *CT_LegendEntry) Validate() error {
	return m.ValidateWithPath("CT_LegendEntry")
}

// ValidateWithPath validates the CT_LegendEntry and its children, prefixing error messages with path
func (m *CT_LegendEntry) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
