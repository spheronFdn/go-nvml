// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/spheronFdn/nvml/pkg/nvml"
	"sync"
)

// Ensure, that Unit does implement nvml.Unit.
// If this is not the case, regenerate this file with moq.
var _ nvml.Unit = &Unit{}

// Unit is a mock implementation of nvml.Unit.
//
//	func TestSomethingThatUsesUnit(t *testing.T) {
//
//		// make and configure a mocked nvml.Unit
//		mockedUnit := &Unit{
//			GetDevicesFunc: func() ([]nvml.Device, nvml.Return) {
//				panic("mock out the GetDevices method")
//			},
//			GetFanSpeedInfoFunc: func() (nvml.UnitFanSpeeds, nvml.Return) {
//				panic("mock out the GetFanSpeedInfo method")
//			},
//			GetLedStateFunc: func() (nvml.LedState, nvml.Return) {
//				panic("mock out the GetLedState method")
//			},
//			GetPsuInfoFunc: func() (nvml.PSUInfo, nvml.Return) {
//				panic("mock out the GetPsuInfo method")
//			},
//			GetTemperatureFunc: func(n int) (uint32, nvml.Return) {
//				panic("mock out the GetTemperature method")
//			},
//			GetUnitInfoFunc: func() (nvml.UnitInfo, nvml.Return) {
//				panic("mock out the GetUnitInfo method")
//			},
//			SetLedStateFunc: func(ledColor nvml.LedColor) nvml.Return {
//				panic("mock out the SetLedState method")
//			},
//		}
//
//		// use mockedUnit in code that requires nvml.Unit
//		// and then make assertions.
//
//	}
type Unit struct {
	// GetDevicesFunc mocks the GetDevices method.
	GetDevicesFunc func() ([]nvml.Device, nvml.Return)

	// GetFanSpeedInfoFunc mocks the GetFanSpeedInfo method.
	GetFanSpeedInfoFunc func() (nvml.UnitFanSpeeds, nvml.Return)

	// GetLedStateFunc mocks the GetLedState method.
	GetLedStateFunc func() (nvml.LedState, nvml.Return)

	// GetPsuInfoFunc mocks the GetPsuInfo method.
	GetPsuInfoFunc func() (nvml.PSUInfo, nvml.Return)

	// GetTemperatureFunc mocks the GetTemperature method.
	GetTemperatureFunc func(n int) (uint32, nvml.Return)

	// GetUnitInfoFunc mocks the GetUnitInfo method.
	GetUnitInfoFunc func() (nvml.UnitInfo, nvml.Return)

	// SetLedStateFunc mocks the SetLedState method.
	SetLedStateFunc func(ledColor nvml.LedColor) nvml.Return

	// calls tracks calls to the methods.
	calls struct {
		// GetDevices holds details about calls to the GetDevices method.
		GetDevices []struct {
		}
		// GetFanSpeedInfo holds details about calls to the GetFanSpeedInfo method.
		GetFanSpeedInfo []struct {
		}
		// GetLedState holds details about calls to the GetLedState method.
		GetLedState []struct {
		}
		// GetPsuInfo holds details about calls to the GetPsuInfo method.
		GetPsuInfo []struct {
		}
		// GetTemperature holds details about calls to the GetTemperature method.
		GetTemperature []struct {
			// N is the n argument value.
			N int
		}
		// GetUnitInfo holds details about calls to the GetUnitInfo method.
		GetUnitInfo []struct {
		}
		// SetLedState holds details about calls to the SetLedState method.
		SetLedState []struct {
			// LedColor is the ledColor argument value.
			LedColor nvml.LedColor
		}
	}
	lockGetDevices      sync.RWMutex
	lockGetFanSpeedInfo sync.RWMutex
	lockGetLedState     sync.RWMutex
	lockGetPsuInfo      sync.RWMutex
	lockGetTemperature  sync.RWMutex
	lockGetUnitInfo     sync.RWMutex
	lockSetLedState     sync.RWMutex
}

// GetDevices calls GetDevicesFunc.
func (mock *Unit) GetDevices() ([]nvml.Device, nvml.Return) {
	if mock.GetDevicesFunc == nil {
		panic("Unit.GetDevicesFunc: method is nil but Unit.GetDevices was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetDevices.Lock()
	mock.calls.GetDevices = append(mock.calls.GetDevices, callInfo)
	mock.lockGetDevices.Unlock()
	return mock.GetDevicesFunc()
}

// GetDevicesCalls gets all the calls that were made to GetDevices.
// Check the length with:
//
//	len(mockedUnit.GetDevicesCalls())
func (mock *Unit) GetDevicesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetDevices.RLock()
	calls = mock.calls.GetDevices
	mock.lockGetDevices.RUnlock()
	return calls
}

// GetFanSpeedInfo calls GetFanSpeedInfoFunc.
func (mock *Unit) GetFanSpeedInfo() (nvml.UnitFanSpeeds, nvml.Return) {
	if mock.GetFanSpeedInfoFunc == nil {
		panic("Unit.GetFanSpeedInfoFunc: method is nil but Unit.GetFanSpeedInfo was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetFanSpeedInfo.Lock()
	mock.calls.GetFanSpeedInfo = append(mock.calls.GetFanSpeedInfo, callInfo)
	mock.lockGetFanSpeedInfo.Unlock()
	return mock.GetFanSpeedInfoFunc()
}

// GetFanSpeedInfoCalls gets all the calls that were made to GetFanSpeedInfo.
// Check the length with:
//
//	len(mockedUnit.GetFanSpeedInfoCalls())
func (mock *Unit) GetFanSpeedInfoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetFanSpeedInfo.RLock()
	calls = mock.calls.GetFanSpeedInfo
	mock.lockGetFanSpeedInfo.RUnlock()
	return calls
}

// GetLedState calls GetLedStateFunc.
func (mock *Unit) GetLedState() (nvml.LedState, nvml.Return) {
	if mock.GetLedStateFunc == nil {
		panic("Unit.GetLedStateFunc: method is nil but Unit.GetLedState was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetLedState.Lock()
	mock.calls.GetLedState = append(mock.calls.GetLedState, callInfo)
	mock.lockGetLedState.Unlock()
	return mock.GetLedStateFunc()
}

// GetLedStateCalls gets all the calls that were made to GetLedState.
// Check the length with:
//
//	len(mockedUnit.GetLedStateCalls())
func (mock *Unit) GetLedStateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetLedState.RLock()
	calls = mock.calls.GetLedState
	mock.lockGetLedState.RUnlock()
	return calls
}

// GetPsuInfo calls GetPsuInfoFunc.
func (mock *Unit) GetPsuInfo() (nvml.PSUInfo, nvml.Return) {
	if mock.GetPsuInfoFunc == nil {
		panic("Unit.GetPsuInfoFunc: method is nil but Unit.GetPsuInfo was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPsuInfo.Lock()
	mock.calls.GetPsuInfo = append(mock.calls.GetPsuInfo, callInfo)
	mock.lockGetPsuInfo.Unlock()
	return mock.GetPsuInfoFunc()
}

// GetPsuInfoCalls gets all the calls that were made to GetPsuInfo.
// Check the length with:
//
//	len(mockedUnit.GetPsuInfoCalls())
func (mock *Unit) GetPsuInfoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPsuInfo.RLock()
	calls = mock.calls.GetPsuInfo
	mock.lockGetPsuInfo.RUnlock()
	return calls
}

// GetTemperature calls GetTemperatureFunc.
func (mock *Unit) GetTemperature(n int) (uint32, nvml.Return) {
	if mock.GetTemperatureFunc == nil {
		panic("Unit.GetTemperatureFunc: method is nil but Unit.GetTemperature was just called")
	}
	callInfo := struct {
		N int
	}{
		N: n,
	}
	mock.lockGetTemperature.Lock()
	mock.calls.GetTemperature = append(mock.calls.GetTemperature, callInfo)
	mock.lockGetTemperature.Unlock()
	return mock.GetTemperatureFunc(n)
}

// GetTemperatureCalls gets all the calls that were made to GetTemperature.
// Check the length with:
//
//	len(mockedUnit.GetTemperatureCalls())
func (mock *Unit) GetTemperatureCalls() []struct {
	N int
} {
	var calls []struct {
		N int
	}
	mock.lockGetTemperature.RLock()
	calls = mock.calls.GetTemperature
	mock.lockGetTemperature.RUnlock()
	return calls
}

// GetUnitInfo calls GetUnitInfoFunc.
func (mock *Unit) GetUnitInfo() (nvml.UnitInfo, nvml.Return) {
	if mock.GetUnitInfoFunc == nil {
		panic("Unit.GetUnitInfoFunc: method is nil but Unit.GetUnitInfo was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetUnitInfo.Lock()
	mock.calls.GetUnitInfo = append(mock.calls.GetUnitInfo, callInfo)
	mock.lockGetUnitInfo.Unlock()
	return mock.GetUnitInfoFunc()
}

// GetUnitInfoCalls gets all the calls that were made to GetUnitInfo.
// Check the length with:
//
//	len(mockedUnit.GetUnitInfoCalls())
func (mock *Unit) GetUnitInfoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetUnitInfo.RLock()
	calls = mock.calls.GetUnitInfo
	mock.lockGetUnitInfo.RUnlock()
	return calls
}

// SetLedState calls SetLedStateFunc.
func (mock *Unit) SetLedState(ledColor nvml.LedColor) nvml.Return {
	if mock.SetLedStateFunc == nil {
		panic("Unit.SetLedStateFunc: method is nil but Unit.SetLedState was just called")
	}
	callInfo := struct {
		LedColor nvml.LedColor
	}{
		LedColor: ledColor,
	}
	mock.lockSetLedState.Lock()
	mock.calls.SetLedState = append(mock.calls.SetLedState, callInfo)
	mock.lockSetLedState.Unlock()
	return mock.SetLedStateFunc(ledColor)
}

// SetLedStateCalls gets all the calls that were made to SetLedState.
// Check the length with:
//
//	len(mockedUnit.SetLedStateCalls())
func (mock *Unit) SetLedStateCalls() []struct {
	LedColor nvml.LedColor
} {
	var calls []struct {
		LedColor nvml.LedColor
	}
	mock.lockSetLedState.RLock()
	calls = mock.calls.SetLedState
	mock.lockSetLedState.RUnlock()
	return calls
}
