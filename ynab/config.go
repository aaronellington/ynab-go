package ynab

import "net/http"

type ConfigFunc func(s *Service)

func SetRoundTripper(rt http.RoundTripper) ConfigFunc {
	return func(s *Service) {
		s.client.httpClient.Transport = rt
	}
}
