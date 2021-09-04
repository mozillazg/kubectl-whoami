package cert

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
	"text/template"

	"github.com/cloudflare/cfssl/certinfo"
	"k8s.io/client-go/rest"
)

const summaryCertInfoTPL = `
User(subject.common_name): {{ .Subject.CommonName }}
Organization(subject.organization): {{ .Subject.Organization }}
Groups(subject.names):
{{- range .Subject.Names }}
 - {{ . }}
{{- end }}
NotBefore(not_before): {{ .NotBefore }}
NotAfter(not_after): {{ .NotAfter }}
`

var rePEM = regexp.MustCompile(`-----BEGIN CERTIFICATE-----[^-]+-----END CERTIFICATE-----`)

func GetCertInfo(config *rest.Config) ([]certinfo.Certificate, error) {
	var certs []certinfo.Certificate
	for _, b := range rePEM.FindAll(config.CertData, -1) {
		cert, err := certinfo.ParseCertificatePEM(b)
		if err != nil {
			return certs, err
		}
		certs = append(certs, *cert)
	}
	return certs, nil
}

func Summary(c certinfo.Certificate) string {
	c.NotBefore = c.NotBefore.Local()
	c.NotAfter = c.NotAfter.Local()
	tpl, _ := template.New("summary").Parse(summaryCertInfoTPL)
	b := bytes.NewBuffer(nil)
	_ = tpl.Execute(b, c)

	return strings.TrimSpace(b.String())
}

func ToJSON(certs []certinfo.Certificate) string {
	d, _ := json.MarshalIndent(certs, "", " ")
	return strings.TrimSpace(string(d))
}
