package apitest

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

type apiFeature struct {
	relyingParty RelyingParty
	responses    []SendRequestAsyncResult
}

func (a *apiFeature) iCallTheSendRequestAPI() error {
	a.iCallTheSendRequestAPIWithReferenceID("meow")
	return nil
}

func (a *apiFeature) iShouldGetARequestID() error {
	if len(a.responses[0].RequestID) == 0 {
		return fmt.Errorf("Did not get a request ID")
	}
	return nil
}

func (a *apiFeature) iCallTheSendRequestAPIWithReferenceID(referenceID string) error {
	params := SendRequestAsyncParams{referenceID}
	response := a.relyingParty.SendRequestAsync(params)
	a.responses = append(a.responses, response)
	return nil
}

func (a *apiFeature) receivedRequestIDsShouldBeTheSame() error {
	if len(a.responses) != 2 {
		return fmt.Errorf("Expected 2 responses but got %d", len(a.responses))
	}
	if a.responses[0].RequestID != a.responses[1].RequestID {
		return fmt.Errorf(
			"Expected the two response IDs to be the same ('%s' != '%s')",
			a.responses[0].RequestID,
			a.responses[1].RequestID,
		)
	}
	return nil
}

func (a *apiFeature) receivedRequestIDsShouldBeDifferent() error {
	if len(a.responses) != 2 {
		return fmt.Errorf("Expected 2 responses but got %d", len(a.responses))
	}
	if a.responses[0].RequestID == a.responses[1].RequestID {
		return fmt.Errorf(
			"Expected the two response IDs to be the different (both are '%s')",
			a.responses[0].RequestID,
		)
	}
	return nil
}

func (a *apiFeature) reset(_ interface{}) {
	a.relyingParty = NewMockRelyingParty()
	a.responses = []SendRequestAsyncResult{}
}

func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}
	s.BeforeScenario(api.reset)
	s.Step(`^I call the send request API$`, api.iCallTheSendRequestAPI)
	s.Step(`^I should get a request ID$`, api.iShouldGetARequestID)
	s.Step(`^I call the send request API with reference ID "([^"]*)"$`, api.iCallTheSendRequestAPIWithReferenceID)
	s.Step(`^received request IDs should be the same$`, api.receivedRequestIDsShouldBeTheSame)
	s.Step(`^received request IDs should be different$`, api.receivedRequestIDsShouldBeDifferent)
}
