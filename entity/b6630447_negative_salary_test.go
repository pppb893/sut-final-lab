package entity_test

import (
	"testing"

	"example.com/finalexam/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestSalalyOutOfRangeCase(t *testing.T) {
	g := NewGomegaWithT(t)
	employee := entity.Employees{
		Name:         "wwwkdwo",
		Salary:       999999999999999999,
		EmployeeCode: "HR-1024",
	}
	ok, err := govalidator.ValidateStruct(employee)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
}
