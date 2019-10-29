package partner

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/devigned/tab"
)

const (
	// CloudPartnerResource is the AAD resource for the Cloud Partner Portal
	CloudPartnerResource = "https://cloudpartner.azure.com"

	// DefaultHost is the default host name for the Cloud Partner Portal
	DefaultHost = "https://cloudpartner.azure.com/"
)

type (
	// Client is the HTTP client for the Cloud Partner Portal
	Client struct {
		Authorizer autorest.Authorizer
		APIVersion string
		Host       string
		mwStack    []MiddlewareFunc
	}

	// ClientOption is a variadic optional configuration func
	ClientOption func(c *Client) error

	// MiddlewareFunc allows a consumer of the Client to inject handlers within the request / response pipeline
	//
	// The example below adds the atom xml content type to the request, calls the next middleware and returns the
	// result.
	//
	// addAtomXMLContentType MiddlewareFunc = func(next RestHandler) RestHandler {
	//		return func(ctx context.Context, req *http.Request) (res *http.Response, e error) {
	//			if req.Method != http.MethodGet && req.Method != http.MethodHead {
	//				req.Header.Add("content-Type", "application/atom+xml;type=entry;charset=utf-8")
	//			}
	//			return next(ctx, req)
	//		}
	//	}
	MiddlewareFunc func(next RestHandler) RestHandler

	// RestHandler is used to transform a request and response within the http pipeline
	RestHandler func(ctx context.Context, req *http.Request) (*http.Response, error)

	// ListOffersParams is the parameters for listing offers
	ListOffersParams struct {
		PublisherID string
	}

	// ShowOfferParams is the parameters for showing an offer
	ShowOfferParams struct {
		PublisherID string
		OfferID     string
	}

	// ShowOfferByVersionParams is the parameters for showing an offer by version
	ShowOfferByVersionParams struct {
		PublisherID string
		OfferID     string
		Version     int
	}

	// ShowOfferBySlotParams is the parameters for showing an offer for a given slot
	ShowOfferBySlotParams struct {
		PublisherID string
		OfferID     string
		SlotID      string
	}

	// SimpleTokenProvider makes it easy to authorize with a string bearer token
	SimpleTokenProvider struct{}
)

var (
	httpLogger MiddlewareFunc = func(next RestHandler) RestHandler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			requestDump, err := httputil.DumpRequest(req, true)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(requestDump))

			res, err := next(ctx, req)
			if err != nil {
				return res, err
			}

			resDump, err := httputil.DumpResponse(res, true)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(resDump))

			return res, err
		}
	}
)

// New creates a new Cloud Provider Portal client
func New(apiVersion string, opts ...ClientOption) (*Client, error) {
	c := &Client{
		Host:       DefaultHost,
		APIVersion: apiVersion,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.Authorizer == nil {
		var a autorest.Authorizer
		if os.Getenv("AZURE_TOKEN") != "" {
			a = new(SimpleTokenProvider)
		} else {
			settings, err := auth.GetSettingsFromEnvironment()
			if err != nil {
				return nil, err
			}
			settings.Values[auth.Resource] = CloudPartnerResource

			a, err = settings.GetAuthorizer()
			if err != nil {
				return nil, err
			}
		}

		c.Authorizer = a
	}

	return c, nil
}

// PutOffer will PUT an offer to the API and return the offer
func (c *Client) PutOffer(ctx context.Context, offer *Offer) (*Offer, error) {
	offerJSON, err := json.Marshal(offer)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("api/publishers/%s/offers/%s?api-version=%s", offer.PublisherID, offer.ID, c.APIVersion)
	res, err := c.execute(ctx, http.MethodPut, path, bytes.NewReader(offerJSON))
	defer closeResponse(ctx, res)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(fmt.Sprintf("uri: %s, status: %d, body: %s", res.Request.URL, res.StatusCode, body))
	}

	var newOffer Offer
	if err := json.Unmarshal(body, &newOffer); err != nil {
		return nil, err
	}

	return &newOffer, nil
}

// GetOfferBySlot will get an offer by publisher and offer ID and version
func (c *Client) GetOfferBySlot(ctx context.Context, params ShowOfferBySlotParams) (*Offer, error) {
	path := fmt.Sprintf("api/publishers/%s/offers/%s/slot/%s?api-version=%s", params.PublisherID, params.OfferID, params.SlotID, c.APIVersion)
	res, err := c.execute(ctx, http.MethodGet, path, nil)
	defer closeResponse(ctx, res)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(fmt.Sprintf("uri: %s, status: %d, body: %s", res.Request.URL, res.StatusCode, body))
	}

	var offer Offer
	if err := json.Unmarshal(body, &offer); err != nil {
		return nil, err
	}

	return &offer, nil
}

// GetOfferByVersion will get an offer by publisher and offer ID and version
func (c *Client) GetOfferByVersion(ctx context.Context, params ShowOfferByVersionParams) (*Offer, error) {
	path := fmt.Sprintf("api/publishers/%s/offers/%s/versions/%d?api-version=%s", params.PublisherID, params.OfferID, params.Version, c.APIVersion)
	res, err := c.execute(ctx, http.MethodGet, path, nil)
	defer closeResponse(ctx, res)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(fmt.Sprintf("uri: %s, status: %d, body: %s", res.Request.URL, res.StatusCode, body))
	}

	var offer Offer
	if err := json.Unmarshal(body, &offer); err != nil {
		return nil, err
	}

	return &offer, nil
}

// GetOffer will get an offer by publisher and offer ID
func (c *Client) GetOffer(ctx context.Context, params ShowOfferParams) (*Offer, error) {
	path := fmt.Sprintf("api/publishers/%s/offers/%s?api-version=%s", params.PublisherID, params.OfferID, c.APIVersion)
	res, err := c.execute(ctx, http.MethodGet, path, nil)
	defer closeResponse(ctx, res)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(fmt.Sprintf("uri: %s, status: %d, body: %s", res.Request.URL, res.StatusCode, body))
	}

	var offers Offer
	if err := json.Unmarshal(body, &offers); err != nil {
		return nil, err
	}

	return &offers, nil
}

// GetOfferStatus gets the status of a given offer
func (c *Client) GetOfferStatus(ctx context.Context, params ShowOfferParams) (*OfferStatus, error) {
	path := fmt.Sprintf("api/publishers/%s/offers/%s/status?api-version=%s", params.PublisherID, params.OfferID, c.APIVersion)
	res, err := c.execute(ctx, http.MethodGet, path, nil)
	defer closeResponse(ctx, res)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(fmt.Sprintf("uri: %s, status: %d, body: %s", res.Request.URL, res.StatusCode, body))
	}

	var status OfferStatus
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, err
	}

	return &status, nil
}

// ListOffers will get all of the offers for a given publisher ID
func (c *Client) ListOffers(ctx context.Context, params ListOffersParams) ([]Offer, error) {
	path := fmt.Sprintf("api/publishers/%s/offers?api-version=%s", params.PublisherID, c.APIVersion)
	res, err := c.execute(ctx, http.MethodGet, path, nil)
	defer closeResponse(ctx, res)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(fmt.Sprintf("uri: %s, status: %d, body: %s", res.Request.URL, res.StatusCode, body))
	}

	var offers []Offer
	if err := json.Unmarshal(body, &offers); err != nil {
		return nil, err
	}

	return offers, nil
}

// ListPublishers will get all of the publishers
func (c *Client) ListPublishers(ctx context.Context) ([]Publisher, error) {
	path := fmt.Sprintf("api/publishers?api-version=%s", c.APIVersion)
	res, err := c.execute(ctx, http.MethodGet, path, nil)
	defer closeResponse(ctx, res)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(fmt.Sprintf("uri: %s, status: %d, body: %s", res.Request.URL, res.StatusCode, body))
	}

	var publishers []Publisher
	if err := json.Unmarshal(body, &publishers); err != nil {
		return nil, err
	}

	return publishers, nil
}

func (c *Client) execute(ctx context.Context, method string, entityPath string, body io.Reader, mw ...MiddlewareFunc) (*http.Response, error) {
	req, err := http.NewRequest(method, c.Host+strings.TrimPrefix(entityPath, "/"), body)
	if err != nil {
		tab.For(ctx).Error(err)
		return nil, err
	}

	final := func(_ RestHandler) RestHandler {
		return func(reqCtx context.Context, request *http.Request) (*http.Response, error) {
			client := &http.Client{
				Timeout: 60 * time.Second,
			}
			request = request.WithContext(reqCtx)
			request.Header.Set("Content-Type", "application/json")
			request, err := autorest.CreatePreparer(c.Authorizer.WithAuthorization()).Prepare(request)
			if err != nil {
				return nil, err
			}

			return client.Do(request)
		}
	}

	mwStack := []MiddlewareFunc{final}
	if os.Getenv("DEBUG") == "true" {
		mwStack = append(mwStack, httpLogger)
	}

	sl := len(c.mwStack) - 1
	for i := sl; i >= 0; i-- {
		mwStack = append(mwStack, c.mwStack[i])
	}

	for i := len(mw) - 1; i >= 0; i-- {
		if mw[i] != nil {
			mwStack = append(mwStack, mw[i])
		}
	}

	var h RestHandler
	for _, mw := range mwStack {
		h = mw(h)
	}

	return h(ctx, req)
}

func closeResponse(ctx context.Context, res *http.Response) {
	if res == nil {
		return
	}

	if err := res.Body.Close(); err != nil {
		tab.For(ctx).Error(err)
	}
}

// WithAuthorization will inject the AZURE_TOKEN env var as the bearer token for API auth
//
// This is useful if you want to use a token from az cli.
// `AZURE_TOKEN=$(az account get-access-token --resource https://cloudpartner.azure.com --query "accessToken" -o tsv) pub publishers list`
func (s SimpleTokenProvider) WithAuthorization() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AZURE_TOKEN")))
			return r, nil
		})
	}
}
