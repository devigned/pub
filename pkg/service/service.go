package service

import (
	"context"

	"github.com/devigned/pub/pkg/format"
	"github.com/devigned/pub/pkg/partner"
)

type (
	// Registry holds the factories and services needed for command execution
	Registry struct {
		CloudPartnerServicerFactory func() (CloudPartnerServicer, error)
		PrinterFactory              func() format.Printer
	}

	// CommandServicer provides all functionality needed for command execution
	CommandServicer interface {
		GetCloudPartnerService() (CloudPartnerServicer, error)
		GetPrinter() format.Printer
	}

	// CloudPartnerServicer provides Azure Cloud Partner functionality
	CloudPartnerServicer interface {
		ListOffers(ctx context.Context, params partner.ListOffersParams) ([]partner.Offer, error)
		GetOfferBySlot(ctx context.Context, params partner.ShowOfferBySlotParams) (*partner.Offer, error)
		GetOfferByVersion(ctx context.Context, params partner.ShowOfferByVersionParams) (*partner.Offer, error)
		GetOffer(ctx context.Context, params partner.ShowOfferParams) (*partner.Offer, error)
		GoLiveWithOffer(ctx context.Context, params partner.GoLiveParams) (string, error)
		GetOfferStatus(ctx context.Context, params partner.ShowOfferParams) (*partner.OfferStatus, error)
		PutOffer(ctx context.Context, offer *partner.Offer) (*partner.Offer, error)
		PublishOffer(ctx context.Context, params partner.PublishOfferParams) (string, error)

		ListOperations(ctx context.Context, params partner.ListOperationsParams) ([]partner.Operation, error)
		CancelOperation(ctx context.Context, params partner.CancelOperationParams) (string, error)
		GetOperationByURI(ctx context.Context, opURI string) (*partner.OperationDetail, error)
		GetOperation(ctx context.Context, params partner.GetOperationParams) (*partner.OperationDetail, error)

		ListPublishers(ctx context.Context) ([]partner.Publisher, error)
	}
)

// GetCloudPartnerService will return a CloudPartnerServicer for interacting with Azure
func (r *Registry) GetCloudPartnerService() (CloudPartnerServicer, error) {
	return r.CloudPartnerServicerFactory()
}

// GetPrinter will return a printer for printing command output
func (r *Registry) GetPrinter() format.Printer {
	return r.PrinterFactory()
}
