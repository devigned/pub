package partner

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/date"
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
		APIVersion  string
	}

	// Offer represents a Cloud Partner Portal offer
	Offer struct {
		TypeID      string          `json:"offerTypeId,omitempty"`
		PublisherID string          `json:"publisherId,omitempty"`
		Status      string          `json:"status,omitempty"`
		ID          string          `json:"id,omitempty"`
		Version     int             `json:"version,omitempty"`
		Definition  OfferDefinition `json:"definition,omitempty"`
		ChangedTime date.Time       `json:"changedTime,omitempty"`
	}

	// OfferDefinition contains offer details
	OfferDefinition struct {
		DisplayText string `json:"displayText,omitempty"`
	}
)

// New creates a new Cloud Provider Portal client
func New(opts ...ClientOption) (*Client, error) {
	c := &Client{
		Host: DefaultHost,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.Authorizer == nil {
		settings, err := auth.GetSettingsFromEnvironment()
		if err != nil {
			return nil, err
		}
		settings.Values[auth.Resource] = CloudPartnerResource

		a, err := settings.GetAuthorizer()
		if err != nil {
			return nil, err
		}
		c.Authorizer = a
	}

	return c, nil
}

// ListOffers will get all of the offers for a given publisher ID
func (c *Client) ListOffers(ctx context.Context, params ListOffersParams) ([]Offer, error) {
	path := fmt.Sprintf("api/publishers/%s/offers?api-version=%s", params.PublisherID, params.APIVersion)
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
		fmt.Println(string(body))
		return nil, err
	}

	return offers, nil
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
