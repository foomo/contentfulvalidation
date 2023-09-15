// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package webhook

import (
	go_context "context"
	go_net_http "net/http"

	github_com_foomo_contentfulvalidation_validator "github.com/foomo/contentfulvalidation/validator"
	gotsrpc "github.com/foomo/gotsrpc/v2"
	pkg_errors "github.com/pkg/errors"
)

type WebhookGoTSRPCClient interface {
	UpdateCache(ctx go_context.Context, sysType github_com_foomo_contentfulvalidation_validator.SysType, modelType github_com_foomo_contentfulvalidation_validator.ModelType, modelID github_com_foomo_contentfulvalidation_validator.ModelID) (clientErr error)
}

type HTTPWebhookGoTSRPCClient struct {
	URL      string
	EndPoint string
	Client   gotsrpc.Client
}

func NewDefaultWebhookGoTSRPCClient(url string) *HTTPWebhookGoTSRPCClient {
	return NewWebhookGoTSRPCClient(url, "/services/contenfulvalidation/webhook")
}

func NewWebhookGoTSRPCClient(url string, endpoint string) *HTTPWebhookGoTSRPCClient {
	return NewWebhookGoTSRPCClientWithClient(url, endpoint, nil)
}

func NewWebhookGoTSRPCClientWithClient(url string, endpoint string, client *go_net_http.Client) *HTTPWebhookGoTSRPCClient {
	return &HTTPWebhookGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
		Client:   gotsrpc.NewClientWithHttpClient(client),
	}
}
func (tsc *HTTPWebhookGoTSRPCClient) UpdateCache(ctx go_context.Context, sysType github_com_foomo_contentfulvalidation_validator.SysType, modelType github_com_foomo_contentfulvalidation_validator.ModelType, modelID github_com_foomo_contentfulvalidation_validator.ModelID) (clientErr error) {
	args := []interface{}{sysType, modelType, modelID}
	reply := []interface{}{}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UpdateCache", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call webhook.WebhookGoTSRPCProxy UpdateCache")
	}
	return
}