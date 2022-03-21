// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package ports

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// Message defines model for Message.
type Message struct {
	// Notification message
	Message string `json:"message"`

	// Notification title
	Title string `json:"title"`
}

// NewSubscriber defines model for NewSubscriber.
type NewSubscriber struct {
	// Client Token
	Token string `json:"token"`
}

// SendNotificationJSONBody defines parameters for SendNotification.
type SendNotificationJSONBody Message

// SubscribeNotificationJSONBody defines parameters for SubscribeNotification.
type SubscribeNotificationJSONBody NewSubscriber

// SendNotificationJSONRequestBody defines body for SendNotification for application/json ContentType.
type SendNotificationJSONRequestBody SendNotificationJSONBody

// SubscribeNotificationJSONRequestBody defines body for SubscribeNotification for application/json ContentType.
type SubscribeNotificationJSONRequestBody SubscribeNotificationJSONBody