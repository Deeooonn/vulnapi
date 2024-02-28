package jwt

import (
	"github.com/cerberauth/vulnapi/internal/auth"
	"github.com/cerberauth/vulnapi/internal/request"
	"github.com/cerberauth/vulnapi/internal/scan"
	"github.com/cerberauth/vulnapi/jwt"
	"github.com/cerberauth/vulnapi/report"
)

const (
	AlgNoneVulnerabilitySeverityLevel = 9
	AlgNoneVulnerabilityName          = "JWT None Algorithm"
	AlgNoneVulnerabilityDescription   = "JWT with none algorithm is accepted allowing to bypass authentication."
)

func AlgNoneJwtScanHandler(operation *request.Operation, ss auth.SecurityScheme) (*report.ScanReport, error) {
	r := report.NewScanReport()
	token := ss.GetValidValueWriter().(*jwt.JWTWriter)

	newToken, err := token.WithAlgNone()
	if err != nil {
		return r, err
	}
	ss.SetAttackValue(newToken)
	vsa, err := scan.ScanURL(operation, &ss)
	if err != nil {
		return r, err
	}
	r.AddScanAttempt(vsa).End()

	if err := scan.DetectNotExpectedResponse(vsa.Response); err != nil {
		r.AddVulnerabilityReport(&report.VulnerabilityReport{
			SeverityLevel: AlgNoneVulnerabilitySeverityLevel,
			Name:          AlgNoneVulnerabilityName,
			Description:   AlgNoneVulnerabilityDescription,
			Operation:     operation,
		})
	}

	return r, nil
}
