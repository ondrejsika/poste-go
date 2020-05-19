package poste

type ListDomainsResponse struct {
	Page         int
	Paging       int
	LastPage     int
	ResultsCount int
	Results      []struct {
		Home             string
		Name             string
		Created          string
		Updated          string
		DomainBin        bool
		DomainBinAddress string
		Forward          bool
		ForwardDomain    string
		ForceRoute       bool
		ForceRouteHost   string
		ReferenceID      string
	}
}
