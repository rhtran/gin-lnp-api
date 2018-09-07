package grpc

/**
service DeviceService {
  // List all registered devices
  rpc GetAllDevices(Empty) returns (Devices) {}
  // Get a device by ID
  rpc GetDeviceByID(ID) returns (Device) {}
  // Update a device&#8217;s state
  rpc SwitchDevice(UpdateDevice) returns (Device) {}
  // Register a new device
  rpc RegisterDevice(Device) returns (Device) {}

  rpc findByOcn (OcnReq) returns (OcnRes);
  rpc findByOcns (OcnsReq) returns (OcnsRes);

}
 */

type GrpcOcnService struct {

}
