package sarbacane

// DESCRIPTION
// delivered Delivered Email
// soft_bounce Bounced Email
// hard_bounce Hard bounced Email
// open Opened Email
// click Clicked Email
// unsubscribe Unsubscribed Email
// complaint Complainted Email
// You will find in the rest of the envelope all or some of the following parameters:

// PARAMETER DESCRIPTION EXAMPLE
// campaignId Id of the campaign "campaignId": "EDE18BYwQwqYA9_BKyTsVr"
// date Event date "date": "2018-01-17T11:21:59.314Z"
// email Recipient Email "email": "recipient.com"
// from Sender Email "from": "sender.com"
// sendId Id of the send "sendId": "BbJ-n53VQ-2_3i5ewR111R"
// subject Message subject "subject": "30% off !"
// category Type of refusal "category": "RELAYING_ISSUES"
// returnCode Return code "returnCode": -1
// smtpDescription SMTP error description "smtpDescription": "There are no DNS entries for the hostname tgipimail.com. I cannot determine where to send this message."
// browser Browser type "browser": "CHROME 63.0.3239.132"
// language Browser language "language": "en,en-US;q=0.9,fr-FR;q
// operatingSystem Operating system "operatingSystem": "WINDOWS"
// webmail Webmail type "webmail": "Unknown"

type SarbacaneWebHook struct {
	email      string `json:"email"`
	campaignId string `json:"email",omitempty`
}
