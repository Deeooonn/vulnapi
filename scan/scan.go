package scan

import (
	"fmt"
	"regexp"

	"github.com/cerberauth/vulnapi/internal/operation"
	"github.com/cerberauth/vulnapi/report"
)

type ScanOptions struct {
	IncludeScans []string
	ExcludeScans []string
	Reporter     *report.Reporter
}

type Scan struct {
	*ScanOptions

	Operations      operation.Operations
	OperationsScans []OperationScan
}

func NewScan(operations operation.Operations, opts *ScanOptions) (*Scan, error) {
	if len(operations) == 0 {
		return nil, fmt.Errorf("a scan must have at least one operation")
	}

	if opts == nil {
		opts = &ScanOptions{}
	}

	if opts.Reporter == nil {
		opts.Reporter = report.NewReporter()
	}

	return &Scan{
		ScanOptions: opts,

		Operations:      operations,
		OperationsScans: []OperationScan{},
	}, nil
}

func (s *Scan) GetOperationsScans() []OperationScan {
	return s.OperationsScans
}

func (s *Scan) AddOperationScanHandler(handler *OperationScanHandler) *Scan {
	if !s.shouldAddScan(handler.ID) {
		return s
	}

	for _, operation := range s.Operations {
		s.OperationsScans = append(s.OperationsScans, OperationScan{
			Operation:   operation,
			ScanHandler: handler,
		})
	}
	return s
}

func (s *Scan) AddScanHandler(handler *OperationScanHandler) *Scan {
	if !s.shouldAddScan(handler.ID) {
		return s
	}

	s.OperationsScans = append(s.OperationsScans, OperationScan{
		Operation:   s.Operations[0],
		ScanHandler: handler,
	})
	return s
}

func (s *Scan) Execute(scanCallback func(operationScan *OperationScan)) (*report.Reporter, []error, error) {
	if scanCallback == nil {
		scanCallback = func(operationScan *OperationScan) {}
	}

	var errors []error
	for _, scan := range s.OperationsScans {
		if scan.ScanHandler == nil {
			continue
		}

		report, err := scan.ScanHandler.Handler(scan.Operation, scan.Operation.SecuritySchemes[0]) // TODO: handle multiple security schemes
		if err != nil {
			errors = append(errors, err)
		}

		if report != nil {
			s.Reporter.AddReport(report)
		}

		scanCallback(&scan)
	}

	return s.Reporter, errors, nil
}

func (s *Scan) shouldAddScan(scanID string) bool {
	// Check if the scan should be excluded
	if len(s.ExcludeScans) > 0 && contains(s.ExcludeScans, scanID) {
		return false
	}

	// Check if the scan should be included
	if len(s.IncludeScans) > 0 && !contains(s.IncludeScans, scanID) {
		return false
	}

	return true
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}

		match, _ := regexp.MatchString(s, item)
		if match {
			return true
		}
	}
	return false
}
