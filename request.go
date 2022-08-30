package tallyGo

import (
	"encoding/json"
	"net/http"
	"time"
)

type WebhookBody struct {
	EventID   string    `json:"eventId"`   //75bd4876-d5f1-4771-9fb6-251c430acbd5
	EventType string    `json:"eventType"` // FORM_RESPONSE
	CreatedAt time.Time `json:"createdAt"` // 2021-08-10T08:00:47.578Z
	Data      Data      `json:"data"`
}

type Data struct {
	SubmissionID string    `json:"submissionId"` //Pn0xNn
	RespondentID string    `json:"respondentId"` // Zw8lwJ
	FormID       string    `json:"formId"`       // pnrB2n ==> this is the WAIVER ID
	FormName     string    `json:"formName"`     // Rise Climbing Waiver
	CreatedAt    time.Time `json:"createdAt"`    // 2021-08-10T08:00:47.578Z
	Fields       []Field   `json:"fields"`
}

func HandleTallyWebhook(r *http.Request, contents *WebhookBody) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(contents)

	if err != nil {
		return err
	}
	return nil
}
