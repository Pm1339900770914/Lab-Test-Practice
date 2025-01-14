package unit

import (
	"Lab-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValidInput(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("all field are valid", func(t *testing.T){
		user := entity.User{
			StudentID: "B6509712",
			Username: "test",
			Password: "test123",
			Email: "email@gmail.com",
			Phone: "0123456789",
			GenderID: 1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestUserName(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("UserName valid", func(t *testing.T){
		user := entity.User{
			StudentID: "B6509712",
			Username: "",
			Password: "test123",
			Email: "email@gmail.com",
			Phone: "0123456789",
			GenderID: 1,
		}

		ok, err := govalidator.ValidateStruct(user)
		
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Username is required"))
	})
}

func TestGmail(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Email is required", func(t *testing.T) {
		user := entity.User{
			StudentID: "B6509712",
			Username: "test",
			Password: "test123",
			Email: "",
			Phone: "0123456789",
			GenderID: 1,	
		}
		
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run("Email invalid", func(t *testing.T) {
		user := entity.User{
			StudentID: "B6509712",
			Username: "test",
			Password: "test123",
			Email: "email",
			Phone: "0123456789",
			GenderID: 1,
		}
		
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
}

func TestStudent(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Student invalid", func(t *testing.T){
		user := entity.User{
			StudentID: "B",
			Username: "test",
			Password: "test123",
			Email: "email@gmail.com",
			Phone: "0123456789",
			GenderID: 1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("StudentID is invalid"))
	})
}

