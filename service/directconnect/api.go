package directconnect

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
)

// AllocateConnectionOnInterconnectRequest generates a request for the AllocateConnectionOnInterconnect operation.
func (c *DirectConnect) AllocateConnectionOnInterconnectRequest(input *AllocateConnectionOnInterconnectInput) (req *aws.Request, output *Connection) {
	if opAllocateConnectionOnInterconnect == nil {
		opAllocateConnectionOnInterconnect = &aws.Operation{
			Name:       "AllocateConnectionOnInterconnect",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAllocateConnectionOnInterconnect, input, output)
	output = &Connection{}
	req.Data = output
	return
}

func (c *DirectConnect) AllocateConnectionOnInterconnect(input *AllocateConnectionOnInterconnectInput) (output *Connection, err error) {
	req, out := c.AllocateConnectionOnInterconnectRequest(input)
	output = out
	err = req.Send()
	return
}

var opAllocateConnectionOnInterconnect *aws.Operation

// AllocatePrivateVirtualInterfaceRequest generates a request for the AllocatePrivateVirtualInterface operation.
func (c *DirectConnect) AllocatePrivateVirtualInterfaceRequest(input *AllocatePrivateVirtualInterfaceInput) (req *aws.Request, output *VirtualInterface) {
	if opAllocatePrivateVirtualInterface == nil {
		opAllocatePrivateVirtualInterface = &aws.Operation{
			Name:       "AllocatePrivateVirtualInterface",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAllocatePrivateVirtualInterface, input, output)
	output = &VirtualInterface{}
	req.Data = output
	return
}

func (c *DirectConnect) AllocatePrivateVirtualInterface(input *AllocatePrivateVirtualInterfaceInput) (output *VirtualInterface, err error) {
	req, out := c.AllocatePrivateVirtualInterfaceRequest(input)
	output = out
	err = req.Send()
	return
}

var opAllocatePrivateVirtualInterface *aws.Operation

// AllocatePublicVirtualInterfaceRequest generates a request for the AllocatePublicVirtualInterface operation.
func (c *DirectConnect) AllocatePublicVirtualInterfaceRequest(input *AllocatePublicVirtualInterfaceInput) (req *aws.Request, output *VirtualInterface) {
	if opAllocatePublicVirtualInterface == nil {
		opAllocatePublicVirtualInterface = &aws.Operation{
			Name:       "AllocatePublicVirtualInterface",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAllocatePublicVirtualInterface, input, output)
	output = &VirtualInterface{}
	req.Data = output
	return
}

func (c *DirectConnect) AllocatePublicVirtualInterface(input *AllocatePublicVirtualInterfaceInput) (output *VirtualInterface, err error) {
	req, out := c.AllocatePublicVirtualInterfaceRequest(input)
	output = out
	err = req.Send()
	return
}

var opAllocatePublicVirtualInterface *aws.Operation

// ConfirmConnectionRequest generates a request for the ConfirmConnection operation.
func (c *DirectConnect) ConfirmConnectionRequest(input *ConfirmConnectionInput) (req *aws.Request, output *ConfirmConnectionOutput) {
	if opConfirmConnection == nil {
		opConfirmConnection = &aws.Operation{
			Name:       "ConfirmConnection",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opConfirmConnection, input, output)
	output = &ConfirmConnectionOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) ConfirmConnection(input *ConfirmConnectionInput) (output *ConfirmConnectionOutput, err error) {
	req, out := c.ConfirmConnectionRequest(input)
	output = out
	err = req.Send()
	return
}

var opConfirmConnection *aws.Operation

// ConfirmPrivateVirtualInterfaceRequest generates a request for the ConfirmPrivateVirtualInterface operation.
func (c *DirectConnect) ConfirmPrivateVirtualInterfaceRequest(input *ConfirmPrivateVirtualInterfaceInput) (req *aws.Request, output *ConfirmPrivateVirtualInterfaceOutput) {
	if opConfirmPrivateVirtualInterface == nil {
		opConfirmPrivateVirtualInterface = &aws.Operation{
			Name:       "ConfirmPrivateVirtualInterface",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opConfirmPrivateVirtualInterface, input, output)
	output = &ConfirmPrivateVirtualInterfaceOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) ConfirmPrivateVirtualInterface(input *ConfirmPrivateVirtualInterfaceInput) (output *ConfirmPrivateVirtualInterfaceOutput, err error) {
	req, out := c.ConfirmPrivateVirtualInterfaceRequest(input)
	output = out
	err = req.Send()
	return
}

var opConfirmPrivateVirtualInterface *aws.Operation

// ConfirmPublicVirtualInterfaceRequest generates a request for the ConfirmPublicVirtualInterface operation.
func (c *DirectConnect) ConfirmPublicVirtualInterfaceRequest(input *ConfirmPublicVirtualInterfaceInput) (req *aws.Request, output *ConfirmPublicVirtualInterfaceOutput) {
	if opConfirmPublicVirtualInterface == nil {
		opConfirmPublicVirtualInterface = &aws.Operation{
			Name:       "ConfirmPublicVirtualInterface",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opConfirmPublicVirtualInterface, input, output)
	output = &ConfirmPublicVirtualInterfaceOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) ConfirmPublicVirtualInterface(input *ConfirmPublicVirtualInterfaceInput) (output *ConfirmPublicVirtualInterfaceOutput, err error) {
	req, out := c.ConfirmPublicVirtualInterfaceRequest(input)
	output = out
	err = req.Send()
	return
}

var opConfirmPublicVirtualInterface *aws.Operation

// CreateConnectionRequest generates a request for the CreateConnection operation.
func (c *DirectConnect) CreateConnectionRequest(input *CreateConnectionInput) (req *aws.Request, output *Connection) {
	if opCreateConnection == nil {
		opCreateConnection = &aws.Operation{
			Name:       "CreateConnection",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateConnection, input, output)
	output = &Connection{}
	req.Data = output
	return
}

func (c *DirectConnect) CreateConnection(input *CreateConnectionInput) (output *Connection, err error) {
	req, out := c.CreateConnectionRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateConnection *aws.Operation

// CreateInterconnectRequest generates a request for the CreateInterconnect operation.
func (c *DirectConnect) CreateInterconnectRequest(input *CreateInterconnectInput) (req *aws.Request, output *Interconnect) {
	if opCreateInterconnect == nil {
		opCreateInterconnect = &aws.Operation{
			Name:       "CreateInterconnect",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateInterconnect, input, output)
	output = &Interconnect{}
	req.Data = output
	return
}

func (c *DirectConnect) CreateInterconnect(input *CreateInterconnectInput) (output *Interconnect, err error) {
	req, out := c.CreateInterconnectRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateInterconnect *aws.Operation

// CreatePrivateVirtualInterfaceRequest generates a request for the CreatePrivateVirtualInterface operation.
func (c *DirectConnect) CreatePrivateVirtualInterfaceRequest(input *CreatePrivateVirtualInterfaceInput) (req *aws.Request, output *VirtualInterface) {
	if opCreatePrivateVirtualInterface == nil {
		opCreatePrivateVirtualInterface = &aws.Operation{
			Name:       "CreatePrivateVirtualInterface",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreatePrivateVirtualInterface, input, output)
	output = &VirtualInterface{}
	req.Data = output
	return
}

func (c *DirectConnect) CreatePrivateVirtualInterface(input *CreatePrivateVirtualInterfaceInput) (output *VirtualInterface, err error) {
	req, out := c.CreatePrivateVirtualInterfaceRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreatePrivateVirtualInterface *aws.Operation

// CreatePublicVirtualInterfaceRequest generates a request for the CreatePublicVirtualInterface operation.
func (c *DirectConnect) CreatePublicVirtualInterfaceRequest(input *CreatePublicVirtualInterfaceInput) (req *aws.Request, output *VirtualInterface) {
	if opCreatePublicVirtualInterface == nil {
		opCreatePublicVirtualInterface = &aws.Operation{
			Name:       "CreatePublicVirtualInterface",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreatePublicVirtualInterface, input, output)
	output = &VirtualInterface{}
	req.Data = output
	return
}

func (c *DirectConnect) CreatePublicVirtualInterface(input *CreatePublicVirtualInterfaceInput) (output *VirtualInterface, err error) {
	req, out := c.CreatePublicVirtualInterfaceRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreatePublicVirtualInterface *aws.Operation

// DeleteConnectionRequest generates a request for the DeleteConnection operation.
func (c *DirectConnect) DeleteConnectionRequest(input *DeleteConnectionInput) (req *aws.Request, output *Connection) {
	if opDeleteConnection == nil {
		opDeleteConnection = &aws.Operation{
			Name:       "DeleteConnection",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteConnection, input, output)
	output = &Connection{}
	req.Data = output
	return
}

func (c *DirectConnect) DeleteConnection(input *DeleteConnectionInput) (output *Connection, err error) {
	req, out := c.DeleteConnectionRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteConnection *aws.Operation

// DeleteInterconnectRequest generates a request for the DeleteInterconnect operation.
func (c *DirectConnect) DeleteInterconnectRequest(input *DeleteInterconnectInput) (req *aws.Request, output *DeleteInterconnectOutput) {
	if opDeleteInterconnect == nil {
		opDeleteInterconnect = &aws.Operation{
			Name:       "DeleteInterconnect",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteInterconnect, input, output)
	output = &DeleteInterconnectOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) DeleteInterconnect(input *DeleteInterconnectInput) (output *DeleteInterconnectOutput, err error) {
	req, out := c.DeleteInterconnectRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteInterconnect *aws.Operation

// DeleteVirtualInterfaceRequest generates a request for the DeleteVirtualInterface operation.
func (c *DirectConnect) DeleteVirtualInterfaceRequest(input *DeleteVirtualInterfaceInput) (req *aws.Request, output *DeleteVirtualInterfaceOutput) {
	if opDeleteVirtualInterface == nil {
		opDeleteVirtualInterface = &aws.Operation{
			Name:       "DeleteVirtualInterface",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteVirtualInterface, input, output)
	output = &DeleteVirtualInterfaceOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) DeleteVirtualInterface(input *DeleteVirtualInterfaceInput) (output *DeleteVirtualInterfaceOutput, err error) {
	req, out := c.DeleteVirtualInterfaceRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteVirtualInterface *aws.Operation

// DescribeConnectionsRequest generates a request for the DescribeConnections operation.
func (c *DirectConnect) DescribeConnectionsRequest(input *DescribeConnectionsInput) (req *aws.Request, output *Connections) {
	if opDescribeConnections == nil {
		opDescribeConnections = &aws.Operation{
			Name:       "DescribeConnections",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeConnections, input, output)
	output = &Connections{}
	req.Data = output
	return
}

func (c *DirectConnect) DescribeConnections(input *DescribeConnectionsInput) (output *Connections, err error) {
	req, out := c.DescribeConnectionsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeConnections *aws.Operation

// DescribeConnectionsOnInterconnectRequest generates a request for the DescribeConnectionsOnInterconnect operation.
func (c *DirectConnect) DescribeConnectionsOnInterconnectRequest(input *DescribeConnectionsOnInterconnectInput) (req *aws.Request, output *Connections) {
	if opDescribeConnectionsOnInterconnect == nil {
		opDescribeConnectionsOnInterconnect = &aws.Operation{
			Name:       "DescribeConnectionsOnInterconnect",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeConnectionsOnInterconnect, input, output)
	output = &Connections{}
	req.Data = output
	return
}

func (c *DirectConnect) DescribeConnectionsOnInterconnect(input *DescribeConnectionsOnInterconnectInput) (output *Connections, err error) {
	req, out := c.DescribeConnectionsOnInterconnectRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeConnectionsOnInterconnect *aws.Operation

// DescribeInterconnectsRequest generates a request for the DescribeInterconnects operation.
func (c *DirectConnect) DescribeInterconnectsRequest(input *DescribeInterconnectsInput) (req *aws.Request, output *DescribeInterconnectsOutput) {
	if opDescribeInterconnects == nil {
		opDescribeInterconnects = &aws.Operation{
			Name:       "DescribeInterconnects",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeInterconnects, input, output)
	output = &DescribeInterconnectsOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) DescribeInterconnects(input *DescribeInterconnectsInput) (output *DescribeInterconnectsOutput, err error) {
	req, out := c.DescribeInterconnectsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeInterconnects *aws.Operation

// DescribeLocationsRequest generates a request for the DescribeLocations operation.
func (c *DirectConnect) DescribeLocationsRequest(input *DescribeLocationsInput) (req *aws.Request, output *DescribeLocationsOutput) {
	if opDescribeLocations == nil {
		opDescribeLocations = &aws.Operation{
			Name:       "DescribeLocations",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeLocations, input, output)
	output = &DescribeLocationsOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) DescribeLocations(input *DescribeLocationsInput) (output *DescribeLocationsOutput, err error) {
	req, out := c.DescribeLocationsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeLocations *aws.Operation

// DescribeVirtualGatewaysRequest generates a request for the DescribeVirtualGateways operation.
func (c *DirectConnect) DescribeVirtualGatewaysRequest(input *DescribeVirtualGatewaysInput) (req *aws.Request, output *DescribeVirtualGatewaysOutput) {
	if opDescribeVirtualGateways == nil {
		opDescribeVirtualGateways = &aws.Operation{
			Name:       "DescribeVirtualGateways",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeVirtualGateways, input, output)
	output = &DescribeVirtualGatewaysOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) DescribeVirtualGateways(input *DescribeVirtualGatewaysInput) (output *DescribeVirtualGatewaysOutput, err error) {
	req, out := c.DescribeVirtualGatewaysRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeVirtualGateways *aws.Operation

// DescribeVirtualInterfacesRequest generates a request for the DescribeVirtualInterfaces operation.
func (c *DirectConnect) DescribeVirtualInterfacesRequest(input *DescribeVirtualInterfacesInput) (req *aws.Request, output *DescribeVirtualInterfacesOutput) {
	if opDescribeVirtualInterfaces == nil {
		opDescribeVirtualInterfaces = &aws.Operation{
			Name:       "DescribeVirtualInterfaces",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeVirtualInterfaces, input, output)
	output = &DescribeVirtualInterfacesOutput{}
	req.Data = output
	return
}

func (c *DirectConnect) DescribeVirtualInterfaces(input *DescribeVirtualInterfacesInput) (output *DescribeVirtualInterfacesOutput, err error) {
	req, out := c.DescribeVirtualInterfacesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeVirtualInterfaces *aws.Operation

type AllocateConnectionOnInterconnectInput struct {
	Bandwidth      *string `locationName:"bandwidth" type:"string" required:"true"`
	ConnectionName *string `locationName:"connectionName" type:"string" required:"true"`
	InterconnectID *string `locationName:"interconnectId" type:"string" required:"true"`
	OwnerAccount   *string `locationName:"ownerAccount" type:"string" required:"true"`
	VLAN           *int64  `locationName:"vlan" type:"integer" required:"true"`

	metadataAllocateConnectionOnInterconnectInput `json:"-", xml:"-"`
}

type metadataAllocateConnectionOnInterconnectInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AllocatePrivateVirtualInterfaceInput struct {
	ConnectionID                         *string                               `locationName:"connectionId" type:"string" required:"true"`
	NewPrivateVirtualInterfaceAllocation *NewPrivateVirtualInterfaceAllocation `locationName:"newPrivateVirtualInterfaceAllocation" type:"structure" required:"true"`
	OwnerAccount                         *string                               `locationName:"ownerAccount" type:"string" required:"true"`

	metadataAllocatePrivateVirtualInterfaceInput `json:"-", xml:"-"`
}

type metadataAllocatePrivateVirtualInterfaceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AllocatePublicVirtualInterfaceInput struct {
	ConnectionID                        *string                              `locationName:"connectionId" type:"string" required:"true"`
	NewPublicVirtualInterfaceAllocation *NewPublicVirtualInterfaceAllocation `locationName:"newPublicVirtualInterfaceAllocation" type:"structure" required:"true"`
	OwnerAccount                        *string                              `locationName:"ownerAccount" type:"string" required:"true"`

	metadataAllocatePublicVirtualInterfaceInput `json:"-", xml:"-"`
}

type metadataAllocatePublicVirtualInterfaceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ConfirmConnectionInput struct {
	ConnectionID *string `locationName:"connectionId" type:"string" required:"true"`

	metadataConfirmConnectionInput `json:"-", xml:"-"`
}

type metadataConfirmConnectionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ConfirmConnectionOutput struct {
	ConnectionState *string `locationName:"connectionState" type:"string"`

	metadataConfirmConnectionOutput `json:"-", xml:"-"`
}

type metadataConfirmConnectionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ConfirmPrivateVirtualInterfaceInput struct {
	VirtualGatewayID   *string `locationName:"virtualGatewayId" type:"string" required:"true"`
	VirtualInterfaceID *string `locationName:"virtualInterfaceId" type:"string" required:"true"`

	metadataConfirmPrivateVirtualInterfaceInput `json:"-", xml:"-"`
}

type metadataConfirmPrivateVirtualInterfaceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ConfirmPrivateVirtualInterfaceOutput struct {
	VirtualInterfaceState *string `locationName:"virtualInterfaceState" type:"string"`

	metadataConfirmPrivateVirtualInterfaceOutput `json:"-", xml:"-"`
}

type metadataConfirmPrivateVirtualInterfaceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ConfirmPublicVirtualInterfaceInput struct {
	VirtualInterfaceID *string `locationName:"virtualInterfaceId" type:"string" required:"true"`

	metadataConfirmPublicVirtualInterfaceInput `json:"-", xml:"-"`
}

type metadataConfirmPublicVirtualInterfaceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ConfirmPublicVirtualInterfaceOutput struct {
	VirtualInterfaceState *string `locationName:"virtualInterfaceState" type:"string"`

	metadataConfirmPublicVirtualInterfaceOutput `json:"-", xml:"-"`
}

type metadataConfirmPublicVirtualInterfaceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Connection struct {
	Bandwidth       *string `locationName:"bandwidth" type:"string"`
	ConnectionID    *string `locationName:"connectionId" type:"string"`
	ConnectionName  *string `locationName:"connectionName" type:"string"`
	ConnectionState *string `locationName:"connectionState" type:"string"`
	Location        *string `locationName:"location" type:"string"`
	OwnerAccount    *string `locationName:"ownerAccount" type:"string"`
	PartnerName     *string `locationName:"partnerName" type:"string"`
	Region          *string `locationName:"region" type:"string"`
	VLAN            *int64  `locationName:"vlan" type:"integer"`

	metadataConnection `json:"-", xml:"-"`
}

type metadataConnection struct {
	SDKShapeTraits bool `type:"structure"`
}

type Connections struct {
	Connections []*Connection `locationName:"connections" type:"list"`

	metadataConnections `json:"-", xml:"-"`
}

type metadataConnections struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateConnectionInput struct {
	Bandwidth      *string `locationName:"bandwidth" type:"string" required:"true"`
	ConnectionName *string `locationName:"connectionName" type:"string" required:"true"`
	Location       *string `locationName:"location" type:"string" required:"true"`

	metadataCreateConnectionInput `json:"-", xml:"-"`
}

type metadataCreateConnectionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateInterconnectInput struct {
	Bandwidth        *string `locationName:"bandwidth" type:"string" required:"true"`
	InterconnectName *string `locationName:"interconnectName" type:"string" required:"true"`
	Location         *string `locationName:"location" type:"string" required:"true"`

	metadataCreateInterconnectInput `json:"-", xml:"-"`
}

type metadataCreateInterconnectInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreatePrivateVirtualInterfaceInput struct {
	ConnectionID               *string                     `locationName:"connectionId" type:"string" required:"true"`
	NewPrivateVirtualInterface *NewPrivateVirtualInterface `locationName:"newPrivateVirtualInterface" type:"structure" required:"true"`

	metadataCreatePrivateVirtualInterfaceInput `json:"-", xml:"-"`
}

type metadataCreatePrivateVirtualInterfaceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreatePublicVirtualInterfaceInput struct {
	ConnectionID              *string                    `locationName:"connectionId" type:"string" required:"true"`
	NewPublicVirtualInterface *NewPublicVirtualInterface `locationName:"newPublicVirtualInterface" type:"structure" required:"true"`

	metadataCreatePublicVirtualInterfaceInput `json:"-", xml:"-"`
}

type metadataCreatePublicVirtualInterfaceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteConnectionInput struct {
	ConnectionID *string `locationName:"connectionId" type:"string" required:"true"`

	metadataDeleteConnectionInput `json:"-", xml:"-"`
}

type metadataDeleteConnectionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteInterconnectInput struct {
	InterconnectID *string `locationName:"interconnectId" type:"string" required:"true"`

	metadataDeleteInterconnectInput `json:"-", xml:"-"`
}

type metadataDeleteInterconnectInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteInterconnectOutput struct {
	InterconnectState *string `locationName:"interconnectState" type:"string"`

	metadataDeleteInterconnectOutput `json:"-", xml:"-"`
}

type metadataDeleteInterconnectOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteVirtualInterfaceInput struct {
	VirtualInterfaceID *string `locationName:"virtualInterfaceId" type:"string" required:"true"`

	metadataDeleteVirtualInterfaceInput `json:"-", xml:"-"`
}

type metadataDeleteVirtualInterfaceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteVirtualInterfaceOutput struct {
	VirtualInterfaceState *string `locationName:"virtualInterfaceState" type:"string"`

	metadataDeleteVirtualInterfaceOutput `json:"-", xml:"-"`
}

type metadataDeleteVirtualInterfaceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeConnectionsInput struct {
	ConnectionID *string `locationName:"connectionId" type:"string"`

	metadataDescribeConnectionsInput `json:"-", xml:"-"`
}

type metadataDescribeConnectionsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeConnectionsOnInterconnectInput struct {
	InterconnectID *string `locationName:"interconnectId" type:"string" required:"true"`

	metadataDescribeConnectionsOnInterconnectInput `json:"-", xml:"-"`
}

type metadataDescribeConnectionsOnInterconnectInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeInterconnectsInput struct {
	InterconnectID *string `locationName:"interconnectId" type:"string"`

	metadataDescribeInterconnectsInput `json:"-", xml:"-"`
}

type metadataDescribeInterconnectsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeInterconnectsOutput struct {
	Interconnects []*Interconnect `locationName:"interconnects" type:"list"`

	metadataDescribeInterconnectsOutput `json:"-", xml:"-"`
}

type metadataDescribeInterconnectsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeLocationsInput struct {
	metadataDescribeLocationsInput `json:"-", xml:"-"`
}

type metadataDescribeLocationsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeLocationsOutput struct {
	Locations []*Location `locationName:"locations" type:"list"`

	metadataDescribeLocationsOutput `json:"-", xml:"-"`
}

type metadataDescribeLocationsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeVirtualGatewaysInput struct {
	metadataDescribeVirtualGatewaysInput `json:"-", xml:"-"`
}

type metadataDescribeVirtualGatewaysInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeVirtualGatewaysOutput struct {
	VirtualGateways []*VirtualGateway `locationName:"virtualGateways" type:"list"`

	metadataDescribeVirtualGatewaysOutput `json:"-", xml:"-"`
}

type metadataDescribeVirtualGatewaysOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeVirtualInterfacesInput struct {
	ConnectionID       *string `locationName:"connectionId" type:"string"`
	VirtualInterfaceID *string `locationName:"virtualInterfaceId" type:"string"`

	metadataDescribeVirtualInterfacesInput `json:"-", xml:"-"`
}

type metadataDescribeVirtualInterfacesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeVirtualInterfacesOutput struct {
	VirtualInterfaces []*VirtualInterface `locationName:"virtualInterfaces" type:"list"`

	metadataDescribeVirtualInterfacesOutput `json:"-", xml:"-"`
}

type metadataDescribeVirtualInterfacesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Interconnect struct {
	Bandwidth         *string `locationName:"bandwidth" type:"string"`
	InterconnectID    *string `locationName:"interconnectId" type:"string"`
	InterconnectName  *string `locationName:"interconnectName" type:"string"`
	InterconnectState *string `locationName:"interconnectState" type:"string"`
	Location          *string `locationName:"location" type:"string"`
	Region            *string `locationName:"region" type:"string"`

	metadataInterconnect `json:"-", xml:"-"`
}

type metadataInterconnect struct {
	SDKShapeTraits bool `type:"structure"`
}

type Location struct {
	LocationCode *string `locationName:"locationCode" type:"string"`
	LocationName *string `locationName:"locationName" type:"string"`

	metadataLocation `json:"-", xml:"-"`
}

type metadataLocation struct {
	SDKShapeTraits bool `type:"structure"`
}

type NewPrivateVirtualInterface struct {
	ASN                  *int64  `locationName:"asn" type:"integer" required:"true"`
	AmazonAddress        *string `locationName:"amazonAddress" type:"string"`
	AuthKey              *string `locationName:"authKey" type:"string"`
	CustomerAddress      *string `locationName:"customerAddress" type:"string"`
	VLAN                 *int64  `locationName:"vlan" type:"integer" required:"true"`
	VirtualGatewayID     *string `locationName:"virtualGatewayId" type:"string" required:"true"`
	VirtualInterfaceName *string `locationName:"virtualInterfaceName" type:"string" required:"true"`

	metadataNewPrivateVirtualInterface `json:"-", xml:"-"`
}

type metadataNewPrivateVirtualInterface struct {
	SDKShapeTraits bool `type:"structure"`
}

type NewPrivateVirtualInterfaceAllocation struct {
	ASN                  *int64  `locationName:"asn" type:"integer" required:"true"`
	AmazonAddress        *string `locationName:"amazonAddress" type:"string"`
	AuthKey              *string `locationName:"authKey" type:"string"`
	CustomerAddress      *string `locationName:"customerAddress" type:"string"`
	VLAN                 *int64  `locationName:"vlan" type:"integer" required:"true"`
	VirtualInterfaceName *string `locationName:"virtualInterfaceName" type:"string" required:"true"`

	metadataNewPrivateVirtualInterfaceAllocation `json:"-", xml:"-"`
}

type metadataNewPrivateVirtualInterfaceAllocation struct {
	SDKShapeTraits bool `type:"structure"`
}

type NewPublicVirtualInterface struct {
	ASN                  *int64               `locationName:"asn" type:"integer" required:"true"`
	AmazonAddress        *string              `locationName:"amazonAddress" type:"string" required:"true"`
	AuthKey              *string              `locationName:"authKey" type:"string"`
	CustomerAddress      *string              `locationName:"customerAddress" type:"string" required:"true"`
	RouteFilterPrefixes  []*RouteFilterPrefix `locationName:"routeFilterPrefixes" type:"list" required:"true"`
	VLAN                 *int64               `locationName:"vlan" type:"integer" required:"true"`
	VirtualInterfaceName *string              `locationName:"virtualInterfaceName" type:"string" required:"true"`

	metadataNewPublicVirtualInterface `json:"-", xml:"-"`
}

type metadataNewPublicVirtualInterface struct {
	SDKShapeTraits bool `type:"structure"`
}

type NewPublicVirtualInterfaceAllocation struct {
	ASN                  *int64               `locationName:"asn" type:"integer" required:"true"`
	AmazonAddress        *string              `locationName:"amazonAddress" type:"string" required:"true"`
	AuthKey              *string              `locationName:"authKey" type:"string"`
	CustomerAddress      *string              `locationName:"customerAddress" type:"string" required:"true"`
	RouteFilterPrefixes  []*RouteFilterPrefix `locationName:"routeFilterPrefixes" type:"list" required:"true"`
	VLAN                 *int64               `locationName:"vlan" type:"integer" required:"true"`
	VirtualInterfaceName *string              `locationName:"virtualInterfaceName" type:"string" required:"true"`

	metadataNewPublicVirtualInterfaceAllocation `json:"-", xml:"-"`
}

type metadataNewPublicVirtualInterfaceAllocation struct {
	SDKShapeTraits bool `type:"structure"`
}

type RouteFilterPrefix struct {
	CIDR *string `locationName:"cidr" type:"string"`

	metadataRouteFilterPrefix `json:"-", xml:"-"`
}

type metadataRouteFilterPrefix struct {
	SDKShapeTraits bool `type:"structure"`
}

type VirtualGateway struct {
	VirtualGatewayID    *string `locationName:"virtualGatewayId" type:"string"`
	VirtualGatewayState *string `locationName:"virtualGatewayState" type:"string"`

	metadataVirtualGateway `json:"-", xml:"-"`
}

type metadataVirtualGateway struct {
	SDKShapeTraits bool `type:"structure"`
}

type VirtualInterface struct {
	ASN                   *int64               `locationName:"asn" type:"integer"`
	AmazonAddress         *string              `locationName:"amazonAddress" type:"string"`
	AuthKey               *string              `locationName:"authKey" type:"string"`
	ConnectionID          *string              `locationName:"connectionId" type:"string"`
	CustomerAddress       *string              `locationName:"customerAddress" type:"string"`
	CustomerRouterConfig  *string              `locationName:"customerRouterConfig" type:"string"`
	Location              *string              `locationName:"location" type:"string"`
	OwnerAccount          *string              `locationName:"ownerAccount" type:"string"`
	RouteFilterPrefixes   []*RouteFilterPrefix `locationName:"routeFilterPrefixes" type:"list"`
	VLAN                  *int64               `locationName:"vlan" type:"integer"`
	VirtualGatewayID      *string              `locationName:"virtualGatewayId" type:"string"`
	VirtualInterfaceID    *string              `locationName:"virtualInterfaceId" type:"string"`
	VirtualInterfaceName  *string              `locationName:"virtualInterfaceName" type:"string"`
	VirtualInterfaceState *string              `locationName:"virtualInterfaceState" type:"string"`
	VirtualInterfaceType  *string              `locationName:"virtualInterfaceType" type:"string"`

	metadataVirtualInterface `json:"-", xml:"-"`
}

type metadataVirtualInterface struct {
	SDKShapeTraits bool `type:"structure"`
}