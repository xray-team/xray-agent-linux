package nginx

const (
	CollectorName = "Nginx"
)

// Metrics
const (
	ResourceName = "NginxStubStatus"

	MetricStubStatusActive   = "Active"
	MetricStubStatusAccepts  = "Accepts"
	MetricStubStatusHandled  = "Handled"
	MetricStubStatusRequests = "Requests"
	MetricStubStatusReading  = "Reading"
	MetricStubStatusWriting  = "Writing"
	MetricStubStatusWaiting  = "Waiting"
)
