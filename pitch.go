package gamedayapi

import "encoding/xml"

type Pitch struct {
	XMLName        xml.Name `xml:"pitch"`
	ID             string   `xml:"id,attr"`
	Type           string   `xml:"type,attr"`
	Des            string   `xml:"des,attr"`
	DesEs          string   `xml:"des_es,attr"`
	TFS            string   `xml:"tfs,attr"`
	TFSZulu        string   `xml:"tfs_zulu,attr"`
	X              string   `xml:"x,attr"`
	Y              string   `xml:"y,attr"`
	SvID           string   `xml:"sv_id,attr"`
	StartSpeed     string   `xml:"start_speed,attr"`
	EndSpeed       string   `xml:"end_speed,attr"`
	SzTop          string   `xml:"sz_top,attr"`
	SzBottom       string   `xml:"sz_bot,attr"`
	PFXX           string   `xml:"pfx_x,attr"`
	PFXZ           string   `xml:"pfx_z,attr"`
	PX             string   `xml:"px,attr"`
	PZ             string   `xml:"pz,attr"`
	X0             string   `xml:"x0,attr"`
	Y0             string   `xml:"y0,attr"`
	Z0             string   `xml:"z0,attr"`
	VX0            string   `xml:"vx0,attr"`
	VY0            string   `xml:"vy0,attr"`
	VZ0            string   `xml:"vz0,attr"`
	AX             string   `xml:"ax,attr"`
	AY             string   `xml:"ay,attr"`
	AZ             string   `xml:"az,attr"`
	BreakY         string   `xml:"break_y,attr"`
	BreakAngle     string   `xml:"break_angle,attr"`
	BreakLength    string   `xml:"break_length,attr"`
	PitchType      string   `xml:"pitch_type,attr"`
	TypeConfidence string   `xml:"type_confidence,attr"`
	Zone           string   `xml:"zone,attr"`
	Nasty          string   `xml:"nasty,attr"`
	SpinDir        string   `xml:"spin_dir,attr"`
	SpinRate       string   `xml:"spin_rate,attr"`
	CC             string   `xml:"cc,attr"`
	MT             string   `xml:"mt,attr"`
}
