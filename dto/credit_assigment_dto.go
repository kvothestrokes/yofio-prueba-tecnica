package dto

type CreditAssigmentRequest struct {
	Investment int32 `json:"investment"`
}

type CreditAssigmentResponse struct {
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`
}

type CreditStatisticsResponse struct {
	TotalAssignations         int     `json:"total_assignations"`
	TotalSuccessAssignation   int     `json:"total_success_assignation"`
	TotalFailedAssignation    int     `json:"total_failed_assignation"`
	AverageSuccessAssignation float32 `json:"average_success_assignation"`
	AverageFailedAssignation  float32 `json:"average_failed_assignation"`
}
