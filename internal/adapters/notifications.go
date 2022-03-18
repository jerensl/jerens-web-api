package adapters

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func NewFirebaseMessagingConnection() (*messaging.Client, error) {
	var opts []option.ClientOption
	ctx := context.Background()

	if file := os.Getenv("SERVICE_ACCOUNT_FILE"); file != "" {
		opts = append(opts, option.WithCredentialsFile(file))
	}
	config := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT")}
	firebaseApp, err := firebase.NewApp(context.Background(), config, opts...)
	if err != nil {
		logrus.Fatalf("error initializing app: %v\n", err)
	}

	client, err := firebaseApp.Messaging(ctx)
	if err != nil {
		logrus.Fatalf("error getting Messaging client: %v\n", err)
	}

	return client, nil
}

type Messaging struct {
	MessagingClient *messaging.Client
}

func (m *Messaging) SendNotification(token []string, messageClient string) error {
	if len(token) < 0 {
		return errors.New("Unable to get token list")
	}

	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "Jerens App",
			Body: messageClient,
		},
		Tokens: token,
	}

	br, err := m.MessagingClient.SendMulticast(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%d messages were sent successfully\n", br.SuccessCount)
	return nil
}