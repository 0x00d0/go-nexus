package v3

type Proxy struct {
	RemoteUrl string `json:"remote_url"`
}

type Attributes struct {
	Proxy Proxy `json:"proxy"`
}

type Repository struct {
	Name       string     `json:"name"`
	Format     string     `json:"format"`
	Type       string     `json:"type"`
	Url        string     `json:"url"`
	Attributes Attributes `json:"attributes"`
}

type ListRepository []*Repository

// ListRepositories List repositories
func (n *Nexus) ListRepositories() (ListRepository, error) {
	endpoint := "/service/rest/v1/repositories"

	var repository ListRepository
	_, err := n.do(endpoint, &repository, nil)
	if err != nil {
		return repository, err
	}
	return repository, nil
}

type Storage struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictConnectTypeValidation bool   `json:"strictConnectTypeValidation"`
	WritePolicy                 string `json:"writePolicy"`
}

type CleanUp struct {
	PolicyNames []string `json:"policyNames"`
}

type Component struct {
	ProprietaryComponents bool `json:"proprietaryComponents"`
}

type HelmHosted struct {
	Name      string    `json:"name"`
	Online    bool      `json:"online"`
	Storage   Storage   `json:"storage"`
	CleanUp   CleanUp   `json:"cleanUp"`
	Component Component `json:"component"`
}

// CreateHelmHosted Create Helm hosted repository
func (n *Nexus) CreateHelmHosted(hosted HelmHosted) error {
	endpoint := "/service/rest/v1/repositories"
	_, err := n.do(endpoint, hosted, nil)
	if err != nil {
		return err
	}
	return nil
}
