package report

// Class struct of coverage report's class tag
type Class struct {
	Name      string       `xml:"name,attr"`
	Namespace string       `xml:"namespace,attr"`
	Package   string       `xml:"fullPackage,attr"`
	Metrics   ClassMetrics `xml:"metrics"`
}

// ClassMetrics struct of coverage report's metrics tag in class tag
type ClassMetrics struct {
	Methods           int32 `xml:"methods,attr"`
	CoveredMethods    int32 `xml:"coveredmethods,attr"`
	Statements        int32 `xml:"statements,attr"`
	CoveredStatements int32 `xml:"coveredstatements,attr"`
	Elements          int32 `xml:"elements,attr"`
	CoveredElements   int32 `xml:"coveredelements,attr"`
}
