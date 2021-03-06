package pagerduty

import (
	"encoding/json"
	"io"
)

type IncidentDetail struct {
	ID                    string           `json:"id"`
	IncidentNumber        uint             `json:"incident_number"`
	CreatedOn             string           `json:"created_on"`
	Status                string           `json:"status"`
	HTMLUrl               string           `json:"html_url"`
	Service               string           `json:"service"`
	AssignedToUser        *json.RawMessage `json:"assigned_to_user"`
	AssignedTo            []string         `json:"assigned_to"`
	TriggerSummaryData    *json.RawMessage `json:"trigger_summary_daya"`
	TriggerDeatilsHTMLUrl string           `json:"trigger_details_html_url"`
}

type WebhookPayload struct {
	ID        string           `json:"id"`
	Type      string           `json:"type"`
	CreatedOn string           `json:"created_on"`
	Data      *json.RawMessage `json:"data"`
}

func DecodeWebhook(r io.Reader) (*WebhookPayload, error) {
	var payload WebhookPayload
	if err := json.NewDecoder(r).Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
