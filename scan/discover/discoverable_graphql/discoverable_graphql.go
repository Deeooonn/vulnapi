package discoverablegraphql

import (
	"github.com/cerberauth/vulnapi/internal/auth"
	"github.com/cerberauth/vulnapi/internal/operation"
	"github.com/cerberauth/vulnapi/report"
	"github.com/cerberauth/vulnapi/scan/discover"
)

const (
	DiscoverableGraphQLPathScanID   = "discover.graphql"
	DiscoverableGraphQLPathScanName = "Discoverable GraphQL Path"
)

type DiscoverableGraphQLPathData = discover.DiscoverData

var issue = report.Issue{
	ID:   "discover.discoverable_graphql",
	Name: "Discoverable GraphQL Endpoint",

	Classifications: &report.Classifications{
		OWASP: report.OWASP_2023_SSRF,
	},

	CVSS: report.CVSS{
		Version: 4.0,
		Vector:  "CVSS:4.0/AV:N/AC:L/AT:N/PR:N/UI:N/VC:N/VI:N/VA:N/SC:N/SI:N/SA:N",
		Score:   0,
	},
}

var graphqlSeclistUrl = "https://raw.githubusercontent.com/cerberauth/vulnapi/main/seclist/lists/graphql.txt"

func ScanHandler(op *operation.Operation, securityScheme *auth.SecurityScheme) (*report.ScanReport, error) {
	vulnReport := report.NewIssueReport(issue).WithOperation(op).WithSecurityScheme(securityScheme)
	r := report.NewScanReport(DiscoverableGraphQLPathScanID, DiscoverableGraphQLPathScanName, op)
	r.AddIssueReport(vulnReport)
	return discover.DownloadAndScanURLs("GraphQL", graphqlSeclistUrl, r, vulnReport, op, securityScheme)
}
