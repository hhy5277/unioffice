// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_SlideTransition struct {
	// Transition Speed
	SpdAttr ST_TransitionSpeed
	// Advance on Click
	AdvClickAttr *bool
	// Advance after time
	AdvTmAttr *uint32
	Choice    *CT_SlideTransitionChoice
	// Sound Action
	SndAc  *CT_TransitionSoundAction
	ExtLst *CT_ExtensionListModify
}

func NewCT_SlideTransition() *CT_SlideTransition {
	ret := &CT_SlideTransition{}
	return ret
}

func (m *CT_SlideTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SpdAttr != ST_TransitionSpeedUnset {
		attr, err := m.SpdAttr.MarshalXMLAttr(xml.Name{Local: "spd"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AdvClickAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "advClick"},
			Value: fmt.Sprintf("%v", *m.AdvClickAttr)})
	}
	if m.AdvTmAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "advTm"},
			Value: fmt.Sprintf("%v", *m.AdvTmAttr)})
	}
	e.EncodeToken(start)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.SndAc != nil {
		sesndAc := xml.StartElement{Name: xml.Name{Local: "p:sndAc"}}
		e.EncodeElement(m.SndAc, sesndAc)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spd" {
			m.SpdAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "advClick" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdvClickAttr = &parsed
		}
		if attr.Name.Local == "advTm" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.AdvTmAttr = &pt
		}
	}
lCT_SlideTransition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "blinds":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Blinds, &el); err != nil {
					return err
				}
			case "checker":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Checker, &el); err != nil {
					return err
				}
			case "circle":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Circle, &el); err != nil {
					return err
				}
			case "dissolve":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Dissolve, &el); err != nil {
					return err
				}
			case "comb":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Comb, &el); err != nil {
					return err
				}
			case "cover":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Cover, &el); err != nil {
					return err
				}
			case "cut":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Cut, &el); err != nil {
					return err
				}
			case "diamond":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Diamond, &el); err != nil {
					return err
				}
			case "fade":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Fade, &el); err != nil {
					return err
				}
			case "newsflash":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Newsflash, &el); err != nil {
					return err
				}
			case "plus":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Plus, &el); err != nil {
					return err
				}
			case "pull":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Pull, &el); err != nil {
					return err
				}
			case "push":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Push, &el); err != nil {
					return err
				}
			case "random":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Random, &el); err != nil {
					return err
				}
			case "randomBar":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.RandomBar, &el); err != nil {
					return err
				}
			case "split":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Split, &el); err != nil {
					return err
				}
			case "strips":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Strips, &el); err != nil {
					return err
				}
			case "wedge":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Wedge, &el); err != nil {
					return err
				}
			case "wheel":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Wheel, &el); err != nil {
					return err
				}
			case "wipe":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Wipe, &el); err != nil {
					return err
				}
			case "zoom":
				m.Choice = NewCT_SlideTransitionChoice()
				if err := d.DecodeElement(&m.Choice.Zoom, &el); err != nil {
					return err
				}
			case "sndAc":
				m.SndAc = NewCT_TransitionSoundAction()
				if err := d.DecodeElement(m.SndAc, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_SlideTransition %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideTransition
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideTransition and its children
func (m *CT_SlideTransition) Validate() error {
	return m.ValidateWithPath("CT_SlideTransition")
}

// ValidateWithPath validates the CT_SlideTransition and its children, prefixing error messages with path
func (m *CT_SlideTransition) ValidateWithPath(path string) error {
	if err := m.SpdAttr.ValidateWithPath(path + "/SpdAttr"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.SndAc != nil {
		if err := m.SndAc.ValidateWithPath(path + "/SndAc"); err != nil {
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
