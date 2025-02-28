package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	BASE_URL = "https://notify.eskiz.uz"
	EMAIL    = "info@softbooking.uz"
	PASSWORD = "3q6qnTOKd8WsQAKYDPo51zSFYIL1x20neYT7aI7l"
)

func GetToken() error {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("email", EMAIL); err != nil {
		return fmt.Errorf("emailni yozishda xatolik: %w", err)
	}
	if err := writer.WriteField("password", PASSWORD); err != nil {
		return fmt.Errorf("passwordni yozishda xatolik: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("multipart yopishda xatolik: %w", err)
	}

	req, err := http.NewRequest("POST", BASE_URL+"/api/auth/login", body)
	if err != nil {
		return fmt.Errorf("so‘rov yaratishda xatolik: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("so‘rov yuborishda xatolik: %w", err)
	}
	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("javobni o‘qishda xatolik: %w", err)
	}

	fmt.Println("Status kodi:", resp.Status)
	fmt.Println("Javob:", string(bodyResp))
	return nil
}
