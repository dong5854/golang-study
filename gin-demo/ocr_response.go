package main

type ocrResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		IdType          string `json:"idType"`
		UserName        string `json:"userName"`
		DriverNo        string `json:"driverNo"`
		JuminNo1        string `json:"juminNo1"`
		JuminNo2        string `json:"juminNo2"`
		MaskedJuminNo2  string `json:"_juminNo2"`
		IssueDate       string `json:"issueDate"`
		ImageBase64Mask string `json:"image_base64_mask,omitempty"`
	}
	TransactionId string `json:"transaction_id"`
}

type ocrData struct {
}
