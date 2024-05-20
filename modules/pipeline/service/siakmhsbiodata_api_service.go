package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	"tracerstudy-datapipeline-service/common/config"
	"tracerstudy-datapipeline-service/modules/pipeline/entity"
)

const (
	apiMaxRetries = 3
	sleepTime     = 500 * time.Millisecond
)

type MhsBiodataService struct {
	cfg config.Config
}

type MhsBiodataServiceUseCase interface {
	FetchMhsBiodataByNimFromSiakApi(nim string) (*entity.MhsSiakBiodata, error)
}

func NewMhsBiodataService(cfg config.Config) *MhsBiodataService {
	return &MhsBiodataService{
		cfg: cfg,
	}
}

func (svc *MhsBiodataService) FetchMhsBiodataByNimFromSiakApi(nim string) (*entity.MhsSiakBiodata, error) {
	payload := map[string]string{"nim": nim}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while marshalling payload:", err)
		if _, isUnsupportedTypeError := err.(*json.UnsupportedTypeError); isUnsupportedTypeError {
			return nil, err
		}
		return nil, err
	}

	apiUrl := svc.cfg.SIAK_API.URL
	apiKey := svc.cfg.SIAK_API.KEY

	for attempt := 1; attempt <= apiMaxRetries; attempt++ {
		reqHttp, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while creating HTTP request:", err)
			return nil, err
		}

		reqHttp.Header.Set("Api-Key", apiKey)
		reqHttp.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(reqHttp)
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while sending HTTP request:", err)

			if attempt == apiMaxRetries {
				log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached:", err)
				return nil, err
			}

			time.Sleep(sleepTime)
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if attempt == apiMaxRetries {
				log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached:", resp.StatusCode, resp.Body)
				return nil, err
			}

			time.Sleep(sleepTime)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while reading HTTP response body:", err)
			return nil, err
		}

		var apiResponse []entity.MhsSiakBiodata
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while unmarshalling HTTP response body:", err)
			return nil, err
		}

		if len(apiResponse) == 0 {
			log.Println("WARNING: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Resource not found: nim", nim)
			return nil, err
		}

		return &apiResponse[0], nil
	}

	log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached without success")
	return nil, err
}
