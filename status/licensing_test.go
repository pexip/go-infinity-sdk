package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetLicensing(t *testing.T) {
	client := &mockClient.Client{}
	expectedLicensingResponse := &LicensingResponse{
		Objects: []Licensing{
			{
				AudioCount:          50,
				AudioTotal:          100,
				PortCount:           25,
				PortTotal:           50,
				SystemCount:         1,
				SystemTotal:         1,
				VMRCount:            10,
				VMRTotal:            20,
				TeamsCount:          15,
				TeamsTotal:          30,
				GHMCount:            5,
				GHMTotal:            10,
				OTJCount:            8,
				OTJTotal:            15,
				SchedulingCount:     3,
				SchedulingTotal:     5,
				TelehealthCount:     2,
				TelehealthTotal:     5,
				CustomLayoutsActive: true,
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/licensing/", mock.AnythingOfType("*status.LicensingResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*LicensingResponse)
		*result = *expectedLicensingResponse
	})

	service := New(client)
	result, err := service.GetLicensing(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, &expectedLicensingResponse.Objects[0], result)
	assert.Equal(t, 50, result.AudioCount)
	assert.Equal(t, 100, result.AudioTotal)
	assert.Equal(t, 25, result.PortCount)
	assert.Equal(t, 50, result.PortTotal)
	assert.Equal(t, 1, result.SystemCount)
	assert.Equal(t, 1, result.SystemTotal)
	assert.Equal(t, 10, result.VMRCount)
	assert.Equal(t, 20, result.VMRTotal)
	assert.Equal(t, 15, result.TeamsCount)
	assert.Equal(t, 30, result.TeamsTotal)
	assert.Equal(t, 5, result.GHMCount)
	assert.Equal(t, 10, result.GHMTotal)
	assert.Equal(t, 8, result.OTJCount)
	assert.Equal(t, 15, result.OTJTotal)
	assert.Equal(t, 3, result.SchedulingCount)
	assert.Equal(t, 5, result.SchedulingTotal)
	assert.Equal(t, 2, result.TelehealthCount)
	assert.Equal(t, 5, result.TelehealthTotal)
	assert.True(t, result.CustomLayoutsActive)
	client.AssertExpectations(t)
}

func TestService_GetLicensing_EmptyResponse(t *testing.T) {
	client := &mockClient.Client{}
	emptyLicensing := &LicensingResponse{
		Objects: []Licensing{},
	}

	client.On("GetJSON", t.Context(), "status/v1/licensing/", mock.AnythingOfType("*status.LicensingResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*LicensingResponse)
		*result = *emptyLicensing
	})

	service := New(client)
	result, err := service.GetLicensing(t.Context())

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "no licensing data returned")
	client.AssertExpectations(t)
}
