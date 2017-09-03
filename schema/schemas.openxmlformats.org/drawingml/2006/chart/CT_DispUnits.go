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

type CT_DispUnits struct {
	Choice       *CT_DispUnitsChoice
	DispUnitsLbl *CT_DispUnitsLbl
	ExtLst       *CT_ExtensionList
}

func NewCT_DispUnits() *CT_DispUnits {
	ret := &CT_DispUnits{}
	return ret
}

func (m *CT_DispUnits) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.DispUnitsLbl != nil {
		sedispUnitsLbl := xml.StartElement{Name: xml.Name{Local: "dispUnitsLbl"}}
		e.EncodeElement(m.DispUnitsLbl, sedispUnitsLbl)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DispUnits) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DispUnits:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "custUnit":
				m.Choice = NewCT_DispUnitsChoice()
				if err := d.DecodeElement(&m.Choice.CustUnit, &el); err != nil {
					return err
				}
			case "builtInUnit":
				m.Choice = NewCT_DispUnitsChoice()
				if err := d.DecodeElement(&m.Choice.BuiltInUnit, &el); err != nil {
					return err
				}
			case "dispUnitsLbl":
				m.DispUnitsLbl = NewCT_DispUnitsLbl()
				if err := d.DecodeElement(m.DispUnitsLbl, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_DispUnits %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DispUnits
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DispUnits and its children
func (m *CT_DispUnits) Validate() error {
	return m.ValidateWithPath("CT_DispUnits")
}

// ValidateWithPath validates the CT_DispUnits and its children, prefixing error messages with path
func (m *CT_DispUnits) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.DispUnitsLbl != nil {
		if err := m.DispUnitsLbl.ValidateWithPath(path + "/DispUnitsLbl"); err != nil {
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
