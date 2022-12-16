package crawler

import (
	"context"
	"log"
	"time"

	"github.com/maxgio92/cloudevents-vpn-provisioner/pkg/vpn"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
)

func Receive(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	// Here is where your code to process the event will go.
	// In this example we will log the event msg
	log.Printf("Event received. \n%s\n", event)

	switch event.Type() {
	case TypeCrawlingPending:
		eventData := &CrawlingPending{}
		if err := event.DataAs(eventData); err != nil {
			log.Printf("Error while extracting cloudevent Data: %s\n", err.Error())

			return nil, cloudevents.NewHTTPResult(400, "failed to c onvert data: %s", err)
		}
		log.Printf("A new Crawling request is pending: %s", eventData.Message)

		respEvent, err := vpn.NewPendingEvent(eventData.VPNProvider, Source)
		if err != nil {
			return nil, cloudevents.NewHTTPResult(500, "failed to set VPN request data: %s", err)
		}
		log.Printf("A VPN service on %s has been requested.", eventData.VPNProvider)

		return respEvent, nil
	case vpn.TypeVpnReady:
		eventData := &vpn.VPNReady{}
		if err := event.DataAs(eventData); err != nil {
			log.Printf("Error while extracting cloudevent Data: %s\n", err.Error())

			return nil, cloudevents.NewHTTPResult(400, "failed to convert data: %s", err)
		}
		log.Printf("A new VPN is ready, provided by: %s", eventData.Provider)
		log.Printf("A new Crawling request is now ready to be satisfied")

		respEvent, err := doCrawl()
		if err != nil {
			return nil, cloudevents.NewHTTPResult(500, "failed to set Crawling response data: %s", err)
		}
		log.Printf("Crawling done")

		return respEvent, nil
	}

	return nil, nil
}

func doCrawl() (*event.Event, error) {

	// Insert logic here.
	time.Sleep(time.Second * 5)

	return newCrawlingDoneEvent()
}

func newCrawlingDoneEvent() (*event.Event, error) {
	crawlingDone := CrawlingDone{
		Message: "A crawler request has been completed.",
		Result:  "this is the result",
	}

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(Source)
	e.SetType(TypeCrawlingComplete)
	if err := e.SetData(cloudevents.ApplicationJSON, crawlingDone); err != nil {
		return nil, err
	}

	return &e, nil
}
