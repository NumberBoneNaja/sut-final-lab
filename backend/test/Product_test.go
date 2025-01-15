package test

import (
	"final-lab/entity"
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestProducts(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`test success`,func(t *testing.T)  {
		product:=entity.Products{
			Name: "test",
			Price:200,
			SKU:"pdd",
		}
		ok,err := govalidator.ValidateStruct(product)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		// g.Expect(err.Error()).To(Equal("Price must be between 1 and 1000"))
	})


	t.Run(`Price not success`,func(t *testing.T)  {
		product:=entity.Products{
			Name: "test",
			Price:-1,
			SKU:"pdd",
		}
		ok,err := govalidator.ValidateStruct(product)
		if !ok {
			fmt.Println("err is",err)
		}
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be between 1 and 1000"))
	})

	t.Run(`Price not `,func(t *testing.T)  {
		product:=entity.Products{
			Name: "test",
			Price:0,
			SKU:"pdd",
		}
		ok,err := govalidator.ValidateStruct(product)
		if !ok {
			fmt.Println("err is",err)
		}
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price is required"))
	})
    t.Run(`SKU not pattern`,func(t *testing.T)  {
		product:=entity.Products{
			Name: "test",
			Price:0,
			SKU:"pdd",
		}
		ok,err := govalidator.ValidateStruct(product)
		if !ok {
			fmt.Println("err is",err)
		}
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price is required"))
	})
	
}