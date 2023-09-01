// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package webhook

import (
	io "io"
	http "net/http"
	time "time"

	github_com_foomo_contentfulvalidation_validator "github.com/foomo/contentfulvalidation/validator"
	gotsrpc "github.com/foomo/gotsrpc/v2"
)

const (
	WebhookGoTSRPCProxyUpdateCache = "UpdateCache"
)

type WebhookGoTSRPCProxy struct {
	EndPoint string
	service  Webhook
}

func NewDefaultWebhookGoTSRPCProxy(service Webhook) *WebhookGoTSRPCProxy {
	return NewWebhookGoTSRPCProxy(service, "/services/contenfulvalidation/webhook")
}

func NewWebhookGoTSRPCProxy(service Webhook, endpoint string) *WebhookGoTSRPCProxy {
	return &WebhookGoTSRPCProxy{
		EndPoint: endpoint,
		service:  service,
	}
}

// ServeHTTP exposes your service
func (p *WebhookGoTSRPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	} else if r.Method != http.MethodPost {
		gotsrpc.ErrorMethodNotAllowed(w)
		return
	}
	defer io.Copy(io.Discard, r.Body) // Drain Request Body

	funcName := gotsrpc.GetCalledFunc(r, p.EndPoint)
	callStats, _ := gotsrpc.GetStatsForRequest(r)
	callStats.Func = funcName
	callStats.Package = "github.com/foomo/contentfulvalidation/webhook"
	callStats.Service = "Webhook"
	switch funcName {
	case WebhookGoTSRPCProxyUpdateCache:
		var (
			args []interface{}
			rets []interface{}
		)
		var (
			arg_sysType   github_com_foomo_contentfulvalidation_validator.SysType
			arg_modelType github_com_foomo_contentfulvalidation_validator.ModelType
			arg_modelID   github_com_foomo_contentfulvalidation_validator.ModelID
		)
		args = []interface{}{&arg_sysType, &arg_modelType, &arg_modelID}
		if err := gotsrpc.LoadArgs(&args, callStats, r); err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		p.service.UpdateCache(arg_sysType, arg_modelType, arg_modelID)
		callStats.Execution = time.Since(executionStart)
		rets = []interface{}{}
		if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {
			gotsrpc.ErrorCouldNotReply(w)
			return
		}
		gotsrpc.Monitor(w, r, args, rets, callStats)
		return
	default:
		gotsrpc.ClearStats(r)
		gotsrpc.ErrorFuncNotFound(w)
	}
}
