package cert

import (
	"bytes"
	"encoding/json"
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

func GetCertInfo(config *rest.Config) (*certinfo.Certificate, error) {
	return certinfo.ParseCertificatePEM(config.CertData)
}

func Summary(c certinfo.Certificate) string {
	c.NotBefore = c.NotBefore.Local()
	c.NotAfter = c.NotAfter.Local()
	tpl, _ := template.New("summary").Parse(summaryCertInfoTPL)
	b := bytes.NewBuffer(nil)
	_ = tpl.Execute(b, c)

	return strings.TrimSpace(b.String())
}

func ToJSON(c certinfo.Certificate) string {
	d, _ := json.MarshalIndent(c, "", " ")
	return strings.TrimSpace(string(d))
}
