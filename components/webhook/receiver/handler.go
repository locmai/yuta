package receiver

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/alertmanager/template"
	"github.com/sirupsen/logrus"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var alerts template.Data
	err := json.NewDecoder(r.Body).Decode(&alerts)
	if err != nil {
		logrus.Errorf("cannot parse content because of %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = logAlertsFromAlertmanager(alerts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func logAlertsFromAlertmanager(alerts template.Data) error {
	logrus.Print(alerts.CommonAnnotations)
	logrus.Print(alerts.CommonLabels)
	logrus.Print(alerts.GroupLabels)
	for _, alert := range alerts.Alerts {
		logrus.Print(alert.Labels)
		logrus.Print(alert.Annotations)

		logrus.Printf("status", alert.Status, "startsAt", alert.StartsAt, "endsAt", alert.EndsAt, "generatorURL", alert.GeneratorURL, "externalURL", alerts.ExternalURL, "receiver", alerts.Receiver, "fingerprint", alert.Fingerprint)
	}

	return nil
}
