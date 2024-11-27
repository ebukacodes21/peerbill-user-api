package gapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"peerbill-user-api/pb"
	"peerbill-user-api/validation"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (s *GServer) GetRates(ctx context.Context, req *pb.RateRequest) (*pb.RateResponse, error) {
	violations := validateRateRequests(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	url := "https://min-api.cryptocompare.com/data/price?"
	rate, err := fetchExchangeRate(url, req)
	if err != nil {
		return nil, err
	}

	resp := &pb.RateResponse{
		Rate: rate,
	}

	return resp, nil
}

func fetchExchangeRate(url string, rr *pb.RateRequest) (float32, error) {
	// Construct the API URL with the crypto and fiat symbols
	apiURL := fmt.Sprintf("%sfsym=%s&tsyms=%s", url, rr.Crypto, rr.Fiat)

	// Fetch the rate from the API
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch exchange rate from API: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the response into a map of fiat currencies to rates
	var data map[string]float32
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// // Get the rate for the requested fiat currency (fiat is a string like "NGN")
	rateStr, ok := data[rr.Fiat]
	if !ok {
		return 0, fmt.Errorf("rate not found in response for fiat %s", rr.Fiat)
	}

	return rateStr, nil
}

func validateRateRequests(req *pb.RateRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validation.ValidateCrypto(req.Crypto); err != nil {
		violations = append(violations, fieldViolation("crypto", err))
	}

	if err := validation.ValidateFiat(req.Fiat); err != nil {
		violations = append(violations, fieldViolation("fiat", err))
	}

	return violations
}
