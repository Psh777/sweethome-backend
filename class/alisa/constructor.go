package alisa

import "../../types"

func CreateDevice(dbDevise types.DBDevice, caps []Capabilitie) Device {
	return Device{
		ID:           dbDevise.ID,
		Name:         dbDevise.Name,
		Description:  dbDevise.Description,
		Room:         dbDevise.Room,
		Type:         dbDevise.AlisaType,
		Capabilities: caps,
	}
}

func CreateDeviceAnswer(token string, devices []Device) Answer {
	payload := Payload{
		UserID:  "psh",
		Devices: devices,
	}
	return Answer{
		RequestID: token,
		Payload:   payload,
	}
}