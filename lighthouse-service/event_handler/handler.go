package event_handler

import (
	"context"
	"fmt"
	"github.com/keptn/go-utils/pkg/sdk/connector/types"
	"net/http"

	keptnapi "github.com/keptn/go-utils/pkg/api/utils"
	"github.com/keptn/go-utils/pkg/sdk/connector/controlplane"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	logger "github.com/sirupsen/logrus"

	keptn "github.com/keptn/go-utils/pkg/lib"
	keptncommon "github.com/keptn/go-utils/pkg/lib/keptn"
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
)

type EvaluationEventHandler interface {
	HandleEvent(ctx context.Context) error
}

func NewEventHandler(ctx context.Context, event cloudevents.Event) (EvaluationEventHandler, error) {
	logger.Debug("Received event: " + event.Type())

	eventSender, ok := ctx.Value(types.EventSenderKey).(controlplane.EventSender)
	if !ok {
		return nil, fmt.Errorf("could not get eventSender from context")
	}
	keptnHandler, err := keptnv2.NewKeptn(&event, keptncommon.KeptnOpts{EventSender: &CPEventSender{Sender: eventSender}})
	if err != nil {
		return nil, err
	}

	configurationServiceEndpoint, err := keptncommon.GetServiceEndpoint("RESOURCE_SERVICE")
	if err != nil {
		return nil, err
	}
	resourceHandler := keptnapi.NewResourceHandler(configurationServiceEndpoint.String())
	serviceHandler := keptnapi.NewServiceHandler(configurationServiceEndpoint.String())

	switch event.Type() {
	case keptnv2.GetTriggeredEventType(keptnv2.EvaluationTaskName):
		return &StartEvaluationHandler{
			Event:             event,
			KeptnHandler:      keptnHandler,
			SLIProviderConfig: K8sSLIProviderConfig{},
			SLOFileRetriever: SLOFileRetriever{
				ResourceHandler: resourceHandler,
				ServiceHandler:  serviceHandler,
			},
		}, nil
	case keptnv2.GetFinishedEventType(keptnv2.GetSLITaskName):
		return &EvaluateSLIHandler{
			Event:        event,
			HTTPClient:   &http.Client{},
			KeptnHandler: keptnHandler,
			SLOFileRetriever: SLOFileRetriever{
				ResourceHandler: resourceHandler,
				ServiceHandler:  serviceHandler,
			},
			EventStore: keptnHandler.EventHandler,
		}, nil
	case keptn.ConfigureMonitoringEventType:
		return NewConfigureMonitoringHandler(event, logger.StandardLogger())
	default:
		logger.Info("received unhandled event type")
		return nil, nil
	}
}
