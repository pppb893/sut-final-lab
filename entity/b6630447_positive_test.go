package entity_test

import (
	"testing"

	"example.com/finalexam/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestSuccessCase(t *testing.T){
	g := NewGomegaWithT(t)
	employee := entity.Employees{
		Name: "wwwkdwo",
		Salary: 15000,
		EmployeeCode: "HR-1024",
	}
	ok, err := govalidator.ValidateStruct(employee)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}