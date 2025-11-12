package emailgateway

type WebhookEvent string

const (
	WebhookEventDelivered    WebhookEvent = "Delivered"
	WebhookEventOpened       WebhookEvent = "Opened"
	WebhookEventClicked      WebhookEvent = "Clicked"
	WebhookEventBounced      WebhookEvent = "Bounced"
	WebhookEventUnsubscribed WebhookEvent = "Unsubscribed"
	WebhookEventComplained   WebhookEvent = "Complained"
)
