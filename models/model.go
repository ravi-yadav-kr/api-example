package models

type Employee struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Age         string `json:"age"`
    DeptName    string `json:"deptName"`
    PhoneNumber string `json:"phoneNumber"`
    Address     string `json:"address"`
    JoiningDate string `json:"joiningDate"`
}
