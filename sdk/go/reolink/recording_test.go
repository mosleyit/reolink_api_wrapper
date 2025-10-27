package reolink

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRecordingAPI_GetRec(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{
			"cmd": "GetRec",
			"code": 0,
			"value": {
				"Rec": {
					"channel": 0,
					"overwrite": 1,
					"postRec": "30 Seconds",
					"preRec": 1,
					"schedule": {
						"enable": 1,
						"table": "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
					}
				}
			}
		}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	rec, err := client.Recording.GetRec(ctx, 0)
	if err != nil {
		t.Fatalf("GetRec failed: %v", err)
	}

	if rec.Channel != 0 {
		t.Errorf("expected channel 0, got %d", rec.Channel)
	}

	if rec.Overwrite != 1 {
		t.Errorf("expected overwrite 1, got %d", rec.Overwrite)
	}

	if rec.PostRec != "30 Seconds" {
		t.Errorf("expected postRec '30 Seconds', got '%s'", rec.PostRec)
	}
}

func TestRecordingAPI_SetRec(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetRec", "code": 0}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	rec := Rec{
		Channel:   0,
		Overwrite: 1,
		PostRec:   "30 Seconds",
		PreRec:    1,
		Schedule: RecSchedule{
			Enable: 1,
			Table:  "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
		},
	}

	err := client.Recording.SetRec(ctx, rec)
	if err != nil {
		t.Fatalf("SetRec failed: %v", err)
	}
}

func TestRecordingAPI_GetRecV20(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{
			"cmd": "GetRecV20",
			"code": 0,
			"value": {
				"Rec": {
					"channel": 0,
					"overwrite": 1,
					"postRec": "1 Minute",
					"preRec": 1,
					"saveDay": 30,
					"schedule": {
						"enable": 1,
						"channel": 0,
						"table": {
							"MD": "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
							"TIMING": "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
						}
					}
				}
			}
		}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	rec, err := client.Recording.GetRecV20(ctx, 0)
	if err != nil {
		t.Fatalf("GetRecV20 failed: %v", err)
	}

	if rec.Channel != 0 {
		t.Errorf("expected channel 0, got %d", rec.Channel)
	}

	if rec.SaveDay != 30 {
		t.Errorf("expected saveDay 30, got %d", rec.SaveDay)
	}

	if rec.PostRec != "1 Minute" {
		t.Errorf("expected postRec '1 Minute', got '%s'", rec.PostRec)
	}
}

func TestRecordingAPI_SetRecV20(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetRecV20", "code": 0}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	rec := Rec{
		Channel:   0,
		Overwrite: 1,
		PostRec:   "1 Minute",
		PreRec:    1,
		SaveDay:   30,
		Schedule: RecSchedule{
			Enable:  1,
			Channel: 0,
			Table: RecScheduleTable{
				MD:     "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
				TIMING: "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			},
		},
	}

	err := client.Recording.SetRecV20(ctx, rec)
	if err != nil {
		t.Fatalf("SetRecV20 failed: %v", err)
	}
}

func TestRecordingAPI_Search(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{
			"cmd": "Search",
			"code": 0,
			"value": {
				"SearchResult": [
					{
						"channel": 0,
						"fileName": "RecM01_20201221_121551_121553.mp4",
						"fileSize": 1024000,
						"startTime": "2020-12-21T12:15:51Z",
						"endTime": "2020-12-21T12:15:53Z",
						"type": "MD"
					}
				]
			}
		}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	startTime := time.Date(2020, 12, 21, 12, 0, 0, 0, time.UTC)
	endTime := time.Date(2020, 12, 21, 13, 0, 0, 0, time.UTC)

	results, err := client.Recording.Search(ctx, 0, startTime, endTime, "main")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].FileName != "RecM01_20201221_121551_121553.mp4" {
		t.Errorf("expected fileName 'RecM01_20201221_121551_121553.mp4', got '%s'", results[0].FileName)
	}

	if results[0].Type != "MD" {
		t.Errorf("expected type 'MD', got '%s'", results[0].Type)
	}
}

func TestRecordingAPI_Download(t *testing.T) {
	client := NewClient("192.168.1.100", WithCredentials("admin", "password"), WithHTTPS(true))
	client.token = "test-token-123"

	url := client.Recording.Download("Mp4Record/2020-12-21/RecM01_20201221_121551_121553.mp4", "recording.mp4")

	expected := "https://192.168.1.100/cgi-bin/api.cgi?cmd=Download&source=Mp4Record/2020-12-21/RecM01_20201221_121551_121553.mp4&output=recording.mp4&token=test-token-123"
	if url != expected {
		t.Errorf("expected URL '%s', got '%s'", expected, url)
	}
}

func TestRecordingAPI_Playback(t *testing.T) {
	client := NewClient("192.168.1.100", WithCredentials("admin", "password"), WithHTTPS(true))
	client.token = "test-token-456"

	url := client.Recording.Playback("Mp4Record/2020-12-22/RecM01_20201222_075939_080140.mp4", "playback.mp4")

	expected := "https://192.168.1.100/cgi-bin/api.cgi?cmd=Playback&source=Mp4Record/2020-12-22/RecM01_20201222_075939_080140.mp4&output=playback.mp4&token=test-token-456"
	if url != expected {
		t.Errorf("expected URL '%s', got '%s'", expected, url)
	}
}

func TestRecordingAPI_NvrDownload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "NvrDownload", "code": 0}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := context.Background()
	params := map[string]interface{}{
		"channel": 0,
		"source":  "test.mp4",
	}

	err := client.Recording.NvrDownload(ctx, params)
	if err != nil {
		t.Fatalf("NvrDownload failed: %v", err)
	}
}
