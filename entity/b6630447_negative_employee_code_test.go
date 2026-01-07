package entity_test

import (
	"testing"

	"example.com/finalexam/entity"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestInvalidEmployeesCodeCase(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	employee := entity.Employees{
		Name:         "wwwkdwo",
		Salary:       15000,
		EmployeeCode: "EEE-1024",
	}
	ok, err := govalidator.ValidateStruct(employee)
	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).NotTo(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
}