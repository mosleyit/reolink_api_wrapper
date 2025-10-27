package reolink

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideoAPI_GetOsd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetOsd", "code": 0, "value": {"Osd": {"channel": 0, "bgcolor": 0, "osdChannel": {"enable": 1, "name": "Camera1", "pos": "Lower Right"}, "osdTime": {"enable": 1, "pos": "Top Center"}, "watermark": 1}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := t.Context()
	osd, err := client.Video.GetOsd(ctx, 0)
	if err != nil {
		t.Fatalf("GetOsd failed: %v", err)
	}

	if osd.Channel != 0 {
		t.Errorf("Expected Channel 0, got %d", osd.Channel)
	}
	if osd.OsdChannel.Name != "Camera1" {
		t.Errorf("Expected OsdChannel.Name Camera1, got %s", osd.OsdChannel.Name)
	}
	if osd.OsdChannel.Pos != "Lower Right" {
		t.Errorf("Expected OsdChannel.Pos 'Lower Right', got %s", osd.OsdChannel.Pos)
	}
	if osd.OsdTime.Enable != 1 {
		t.Errorf("Expected OsdTime.Enable 1, got %d", osd.OsdTime.Enable)
	}
}

func TestVideoAPI_SetOsd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetOsd", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := t.Context()
	osd := Osd{
		Channel: 0,
		BgColor: 0,
		OsdChannel: OsdChannel{
			Enable: 1,
			Name:   "Camera1",
			Pos:    "Lower Right",
		},
		OsdTime: OsdTime{
			Enable: 1,
			Pos:    "Top Center",
		},
		Watermark: 1,
	}

	err := client.Video.SetOsd(ctx, osd)
	if err != nil {
		t.Fatalf("SetOsd failed: %v", err)
	}
}

func TestVideoAPI_GetImage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetImage", "code": 0, "value": {"Image": {"channel": 0, "bright": 128, "contrast": 128, "saturation": 128, "hue": 128, "sharpen": 128}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := t.Context()
	image, err := client.Video.GetImage(ctx, 0)
	if err != nil {
		t.Fatalf("GetImage failed: %v", err)
	}

	if image.Channel != 0 {
		t.Errorf("Expected Channel 0, got %d", image.Channel)
	}
	if image.Bright != 128 {
		t.Errorf("Expected Bright 128, got %d", image.Bright)
	}
	if image.Contrast != 128 {
		t.Errorf("Expected Contrast 128, got %d", image.Contrast)
	}
	if image.Saturation != 128 {
		t.Errorf("Expected Saturation 128, got %d", image.Saturation)
	}
	if image.Hue != 128 {
		t.Errorf("Expected Hue 128, got %d", image.Hue)
	}
	if image.Sharpen != 128 {
		t.Errorf("Expected Sharpen 128, got %d", image.Sharpen)
	}
}

func TestVideoAPI_SetImage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetImage", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := t.Context()
	image := Image{
		Channel:    0,
		Bright:     150,
		Contrast:   150,
		Saturation: 150,
		Hue:        150,
		Sharpen:    150,
	}

	err := client.Video.SetImage(ctx, image)
	if err != nil {
		t.Fatalf("SetImage failed: %v", err)
	}
}

func TestVideoAPI_GetIsp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetIsp", "code": 0, "value": {"Isp": {"channel": 0, "antiFlicker": "Outdoor", "exposure": "Auto", "gain": {"min": 1, "max": 62}, "dayNight": "Auto", "backLight": "Off", "blc": 128, "drc": 128, "rotation": 0, "mirroring": 0, "nr3d": 50}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := t.Context()
	isp, err := client.Video.GetIsp(ctx, 0)
	if err != nil {
		t.Fatalf("GetIsp failed: %v", err)
	}

	if isp.Channel != 0 {
		t.Errorf("Expected Channel 0, got %d", isp.Channel)
	}
	if isp.AntiFlicker != "Outdoor" {
		t.Errorf("Expected AntiFlicker Outdoor, got %s", isp.AntiFlicker)
	}
	if isp.Exposure != "Auto" {
		t.Errorf("Expected Exposure Auto, got %s", isp.Exposure)
	}
	if isp.Gain.Min != 1 {
		t.Errorf("Expected Gain.Min 1, got %d", isp.Gain.Min)
	}
	if isp.Gain.Max != 62 {
		t.Errorf("Expected Gain.Max 62, got %d", isp.Gain.Max)
	}
	if isp.DayNight != "Auto" {
		t.Errorf("Expected DayNight Auto, got %s", isp.DayNight)
	}
}

func TestVideoAPI_SetIsp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetIsp", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := t.Context()
	isp := Isp{
		Channel:     0,
		AntiFlicker: "Outdoor",
		Exposure:    "Auto",
		Gain: IspGain{
			Min: 1,
			Max: 62,
		},
		DayNight:  "Auto",
		BackLight: "Off",
		Blc:       128,
		Drc:       128,
		Rotation:  0,
		Mirroring: 0,
		Nr3d:      50,
	}

	err := client.Video.SetIsp(ctx, isp)
	if err != nil {
		t.Fatalf("SetIsp failed: %v", err)
	}
}

func TestVideoAPI_GetMask(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "GetMask", "code": 0, "value": {"Mask": {"channel": 0, "enable": 1, "area": [{"screen": {"height": 1080, "width": 1920}, "x": 100, "y": 100, "width": 200, "height": 200}]}}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := t.Context()
	mask, err := client.Video.GetMask(ctx, 0)
	if err != nil {
		t.Fatalf("GetMask failed: %v", err)
	}

	if mask.Channel != 0 {
		t.Errorf("Expected Channel 0, got %d", mask.Channel)
	}
	if mask.Enable != 1 {
		t.Errorf("Expected Enable 1, got %d", mask.Enable)
	}
	if len(mask.Area) != 1 {
		t.Fatalf("Expected 1 area, got %d", len(mask.Area))
	}
	if mask.Area[0].X != 100 {
		t.Errorf("Expected Area[0].X 100, got %d", mask.Area[0].X)
	}
	if mask.Area[0].Width != 200 {
		t.Errorf("Expected Area[0].Width 200, got %d", mask.Area[0].Width)
	}
}

func TestVideoAPI_SetMask(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"cmd": "SetMask", "code": 0, "value": {"rspCode": 200}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL
	client.token = "test-token"

	ctx := t.Context()
	mask := Mask{
		Channel: 0,
		Enable:  1,
		Area: []MaskArea{
			{
				Screen: MaskScreen{
					Height: 1080,
					Width:  1920,
				},
				X:      100,
				Y:      100,
				Width:  200,
				Height: 200,
			},
		},
	}

	err := client.Video.SetMask(ctx, mask)
	if err != nil {
		t.Fatalf("SetMask failed: %v", err)
	}
}

func TestVideoAPI_GetCrop(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetCrop",
			Code: 0,
			Value: json.RawMessage(`{
				"Crop": {
					"channel": 0,
					"screenWidth": 2560,
					"screenHeight": 1920,
					"cropWidth": 640,
					"cropHeight": 480,
					"topLeftX": 960,
					"topLeftY": 720
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := t.Context()
	crop, err := client.Video.GetCrop(ctx, 0)
	if err != nil {
		t.Fatalf("GetCrop failed: %v", err)
	}

	if crop.Channel != 0 {
		t.Errorf("expected Channel 0, got %d", crop.Channel)
	}
	if crop.ScreenWidth != 2560 {
		t.Errorf("expected ScreenWidth 2560, got %d", crop.ScreenWidth)
	}
	if crop.ScreenHeight != 1920 {
		t.Errorf("expected ScreenHeight 1920, got %d", crop.ScreenHeight)
	}
	if crop.CropWidth != 640 {
		t.Errorf("expected CropWidth 640, got %d", crop.CropWidth)
	}
	if crop.CropHeight != 480 {
		t.Errorf("expected CropHeight 480, got %d", crop.CropHeight)
	}
	if crop.TopLeftX != 960 {
		t.Errorf("expected TopLeftX 960, got %d", crop.TopLeftX)
	}
	if crop.TopLeftY != 720 {
		t.Errorf("expected TopLeftY 720, got %d", crop.TopLeftY)
	}
}

func TestVideoAPI_SetCrop(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:   "SetCrop",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := t.Context()
	crop := Crop{
		Channel:      0,
		ScreenWidth:  2560,
		ScreenHeight: 1920,
		CropWidth:    640,
		CropHeight:   480,
		TopLeftX:     960,
		TopLeftY:     720,
	}

	err := client.Video.SetCrop(ctx, crop)
	if err != nil {
		t.Fatalf("SetCrop failed: %v", err)
	}
}

func TestVideoAPI_GetStitch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:  "GetStitch",
			Code: 0,
			Value: json.RawMessage(`{
				"stitch": {
					"distance": 8.1,
					"stitchXMove": 5,
					"stitchYMove": 3
				}
			}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := t.Context()
	stitch, err := client.Video.GetStitch(ctx)
	if err != nil {
		t.Fatalf("GetStitch failed: %v", err)
	}

	if stitch.Distance != 8.1 {
		t.Errorf("expected Distance 8.1, got %f", stitch.Distance)
	}
	if stitch.StitchXMove != 5 {
		t.Errorf("expected StitchXMove 5, got %d", stitch.StitchXMove)
	}
	if stitch.StitchYMove != 3 {
		t.Errorf("expected StitchYMove 3, got %d", stitch.StitchYMove)
	}
}

func TestVideoAPI_SetStitch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []Response{{
			Cmd:   "SetStitch",
			Code:  0,
			Value: json.RawMessage(`{"rspCode": 200}`),
		}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.URL[7:])
	client.baseURL = server.URL

	ctx := t.Context()
	stitch := Stitch{
		Distance:    8.1,
		StitchXMove: 5,
		StitchYMove: 3,
	}

	err := client.Video.SetStitch(ctx, stitch)
	if err != nil {
		t.Fatalf("SetStitch failed: %v", err)
	}
}
