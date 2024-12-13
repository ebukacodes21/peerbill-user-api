package gapi

import (
	"context"

	"github.com/ebukacodes21/peerbill-user-api/pb"
	// "github.com/ebukacodes21/peerbill-user-api/validation"
	// "google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (s *GServer) GetDataPlans(ctx context.Context, req *pb.DataLookupRequest) (*pb.DataLookupResponse, error) {
	// violations := validateGetDataplansRequests(req)
	// if violations != nil {
	// 	return nil, invalidArgumentError(violations)
	// }

	// url := "https://min-api.cryptocompare.com/data/price?"
	// rate, err := fetchExchangeRate(url, req)
	// if err != nil {
	// 	return nil, err
	// }

	// resp := &pb.RateResponse{
	// 	Rate: rate,
	// }

	return nil, nil
}

// func fetchDataplans(url string, rr *pb.DataLookupRequest) (float32, error) {
// 	// Construct the API URL with the crypto and fiat symbols
// 	apiURL := fmt.Sprintf("%sfsym=%s&tsyms=%s", url, rr.Crypto, rr.Fiat)

// 	// Fetch the rate from the API
// 	resp, err := http.Get(apiURL)
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to fetch exchange rate from API: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to read response body: %v", err)
// 	}

// 	// Unmarshal the response into a map of fiat currencies to rates
// 	var data map[string]float32
// 	if err := json.Unmarshal(body, &data); err != nil {
// 		return 0, fmt.Errorf("failed to unmarshal response body: %v", err)
// 	}

// 	// // Get the rate for the requested fiat currency (fiat is a string like "NGN")
// 	rateStr, ok := data[rr.Fiat]
// 	if !ok {
// 		return 0, fmt.Errorf("rate not found in response for fiat %s", rr.Fiat)
// 	}

// 	return rateStr, nil
// }

// func validateGetDataplansRequests(req *pb.DataLookupRequest) (violation []*errdetails.BadRequest_FieldViolation) {
// 	if err := validation.IsSupported(req.Network); err != nil {
// 		violation = append(violation, fieldViolation("network", err))
// 	}
// 	return violation
// }
