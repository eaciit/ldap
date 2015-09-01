package ldap

import (
	"github.com/rbns/asn1-ber"
)

// PasswdModifyRequestValue ::= SEQUENCE {
//      userIdentity    [0]  OCTET STRING OPTIONAL
//      oldPasswd       [1]  OCTET STRING OPTIONAL
//      newPasswd       [2]  OCTET STRING OPTIONAL }

type PasswordModifyRequest struct {
	UserIdentity	string
	OldPasswd	string
	NewPasswd	string
}

func (r *PasswordModifyRequest) Encode() (*ber.Packet, error) {
	p := ber.Encode(ber.ClassApplication, ber.TypeConstructed, uint8(ApplicationExtendedRequest), nil, "PasswordModifyRequest")
	p.AppendChild(ber.NewString(ber.ClassContext, ber.TypePrimative, 0, "1.3.6.1.4.1.4203.1.11.1", "Password Modify Request"))

	octetString := ber.Encode(ber.ClassContext, ber.TypePrimative, 1, nil, "Octet String")
	value := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "PasswordModifyRequestValue")
	
	if r.UserIdentity != "" {
		userIdentity := ber.NewString(ber.ClassContext, ber.TypePrimative, 0, string(r.UserIdentity), "userIdentity")
		value.AppendChild(userIdentity)
	}

	if r.OldPasswd != "" {
		oldPasswd := ber.NewString(ber.ClassContext, ber.TypePrimative, 1, string(r.OldPasswd), "oldPasswd")
		value.AppendChild(oldPasswd)
	}

	if r.NewPasswd != "" {
		newPasswd := ber.NewString(ber.ClassContext, ber.TypePrimative, 2, string(r.NewPasswd), "newPasswd")
		value.AppendChild(newPasswd)
	}

	octetString.AppendChild(value)
	p.AppendChild(octetString)

	return p, nil
}

func (l *Connection) Passwd(req *PasswordModifyRequest) error {
	messageID, ok := l.nextMessageID()
	if !ok {
		return newError(ErrorClosing, "MessageID channel is closed.")
	}

	encodedReq, err := req.Encode()
	if err != nil {
		return err
	}

	packet, err := requestBuildPacket(messageID, encodedReq, nil)
	
	return l.sendReqRespPacket(messageID, packet)	
}
