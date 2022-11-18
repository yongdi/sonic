package vo

import "sonic/consts"

type MFAFactorAuth struct {
	QRImage    string         `json:"qrImage"`
	OptAuthUrl string         `json:"optAuthUrl"`
	MFAKey     string         `json:"mfaKey"`
	MFAType    consts.MFAType `json:"mfaType"`
}
