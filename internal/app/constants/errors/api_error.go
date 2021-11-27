package errors

import "errors"

// Default return login
const (
	FailMessage                           = "Failed to fetch data"
	CreateFailMessage                     = "Failed to create data"
	UpdateFailMessage                     = "Failed to update data"
	DeleteFailMessage                     = "Failed to delete data"
	FailedFieldsAssociationMessage        = "Failed while associating fields from request"
	FailedToConvertObj                    = "Failed on object conversion"
	FailedToStoreAuthenticationKeyOnCache = "Failed to store authentication key on cache"
	FailedToGetAuthenticationOnCache      = "Failed to get authentication on cache"
	FailedToParseAuthenticationFromCache  = "Failed to parse authentication from cache"
	FailedToMarshalAuthenticationOnCache  = "Failed to marshal authentication on cache"
	FailedToDeleteAuthenticationOnCache   = "Failed to delete authentication on cache"
	FailedToCreateConsumer                = "Failed to create consumer"
	FailedToDeleteConsumer                = "Failed to delete consumer"
	FailedToRetrieveConsumerKey           = "Failed to retrieve consumer"
	FailedToUpdateInventoryAmount         = "Failed to update inventory amount"
)

// ApiError will be used on API Errors
type ApiError struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (e ApiError) GetError() error {
	return errors.New(e.Error)
}
