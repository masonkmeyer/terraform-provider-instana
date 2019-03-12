package api

//InstanaDataObject is a marker interface for any data object provided by any resource of the Instana REST API
type InstanaDataObject interface {
	GetID() string
	Validate() error
}

//RestClient interface to access REST resources of the Instana API
type RestClient interface {
	GetOne(id string, resourcePath string) ([]byte, error)
	GetAll(resourcePath string) ([]byte, error)
	Put(data InstanaDataObject, resourcePath string) error
	Delete(resourceID string, resourceBasePath string) error
}
