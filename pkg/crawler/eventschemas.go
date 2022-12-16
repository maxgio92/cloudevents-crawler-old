package crawler

import (
	"github.com/maxgio92/cloudevents-vpn-provisioner/pkg/vpn"
)

type SocialModule string

type AvatarID string

type CrawlingResult string

// CrawlingPending defines the Data of CloudEvent with type=dev.knative.crawler.pending
type CrawlingPending struct {

	// Message holds the message from the event
	Message string `json:"message,omitempty"`

	// VPNProvider holds the provider of the tunnel to be used for the crawler requests
	VPNProvider vpn.VPNProvider `json:"vpn_provider"`

	// SocialModule holds the social network module name where the avatar resides
	SocialModule SocialModule `json:"social_module"`

	// AvatarID is the unique ID of the avatar
	AvatarID AvatarID `json:"avatar_id"`
}

// CrawlingDone defines the Data of CloudEvent with type=dev.knative.crawler.complete
type CrawlingDone struct {

	// Message holds the message from the event
	Message string `json:"message"`

	// Result holds the data that have been crawled
	Result CrawlingResult `json:"result"`
}
