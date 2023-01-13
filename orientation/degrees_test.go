package orientation

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DegreesTests", func() {
	Context("ScaleDegrees", func() {
		It("Should scale degrees to range -180 to 180", func() {
			cases := []struct {
				in       float64
				scale360 bool
				exp      float64
			}{
				{90, false, 90},
				{180, false, 179.99},
				{190, false, -170},
				{360, false, 0},
				{720, false, 0},
				{347, false, -13},
				{-50, false, -50},
				{-410, false, -50},
			}

			for _, c := range cases {
				got := ScaleDegrees(c.in, c.scale360)
				Expect(got).To(Equal(c.exp))
			}
		})
		It("Should scale degrees to range 0 to 360", func() {
			cases := []struct {
				in       float64
				scale360 bool
				exp      float64
			}{
				{190, true, 190},
				{360, true, 0},
				{720, true, 0},
				{347, true, 347},
				{-50, true, 310},
				{-410, true, 310},
				{-730, true, 350},
			}

			for _, c := range cases {
				got := ScaleDegrees(c.in, c.scale360)
				Expect(got).To(Equal(c.exp))
			}
		})
		It("Should scale degrees with floating point decimals", func() {
			casesLossy := []struct {
				in       float64
				scale360 bool
				exp      float64
			}{
				{4.2783, false, 4.2783},
				{4.2783, true, 4.2783},
				{759.53, false, 39.53},
				{759.53, true, 39.53},
				{360.00001984784285, true, 0.00001984784285},
			}

			for _, c := range casesLossy {
				got := ScaleDegrees(c.in, c.scale360)
				Expect(got).To(BeNumerically("~", c.exp, 0.00001))
			}
		})
	})
	Context("DegreeWrapDelta", func() {
		It("Should return wrapped degree delta correctly", func() {
			cases := []struct {
				a   float64
				b   float64
				exp float64
			}{
				{-45, 45, 90},
				{45, -45, -90},
				{90.0, 91, 1},
				{170.0, -170, 20},
				{-170.0, 170, -20},
				{-90.0, 90, -180},
				{90.0, -90, -180},
				{-90.0, 170, -100},
				{170, -90, 100},
			}

			for _, c := range cases {
				got := DegreeWrapDelta(c.a, c.b)
				// fmt.Println("got:[", got, "]", "expected [", c.exp, "]")
				Expect(got).To(Equal(c.exp))
			}
		})
	})
	Context("DegreeDelta", func() {
		It("Should return degree delta correctly", func() {
			cases := []struct {
				a   float64
				b   float64
				min float64
				max float64
				exp float64
			}{
				{-45, 45, -180, 180, 90},
				{45, -45, -180, 180, -90},
				{90.0, 91, -180, 180, 1},
				{200, -170, -217, 217, -370},
				{200, 215, -217, 217, 15},
				{200, -200, -217, 217, -400},
				{170.0, -170, -180, 180, -340},
				{-170.0, 170, -180, 180, 340},
				{-90.0, 90, -180, 180, 180},
				{90.0, -90, -180, 180, -180},
				{-90.0, 170, -180, 180, 260},
				{170, -90, -180, 180, -260},
				{62, -217.5, -217.5, 217.5, -279.5},
				{62, 217.5, -217.5, 217.5, 155.5},
				{170, 210, -217, 217, 40},
				{180, 210, -217, 217, 30},
				{190, 210, -217, 217, 20},
				{200, 210, -217, 217, 10},
				{210, 210, -217, 217, 0},
				{0, 180, -180, 180, 180},
				{0, 190, -180, 180, 190},
				{0, -170, -180, 180, -170},
				{0, 190, -217, 217, 190},
			}

			for _, c := range cases {
				got := DegreeDelta(c.a, c.b, c.min, c.max)
				// fmt.Println("got:[", got, "]", "expected [", c.exp, "]")
				Expect(got).To(Equal(c.exp))
			}
		})
	})
	Context("DegreeClampDelta", func() {
		It("Should return clamped degree delta correctly", func() {
			cases := []struct {
				a   float64
				b   float64
				min float64
				max float64
				exp float64
			}{
				{-45, 45, -180, 180, 90},
				{45, -45, -180, 180, -90},
				{90.0, 91, -180, 180, 1},
				{200, -170, -217, 217, -370},
				{200, 215, -217, 217, 15},
				{200, -200, -217, 217, -400},
				{170.0, -170, -180, 180, -340},
				{-170.0, 170, -180, 180, 340},
				{-90.0, 90, -180, 180, 180},
				{90.0, -90, -180, 180, -180},
				{-90.0, 170, -180, 180, 260},
				{170, -90, -180, 180, -260},
				{62, -217.5, -217.5, 217.5, -279.5},
				{62, 217.5, -217.5, 217.5, 155.5},
				{170, 210, -217, 217, 40},
				{180, 210, -217, 217, 30},
				{190, 210, -217, 217, 20},
				{200, 210, -217, 217, 10},
				{210, 210, -217, 217, 0},
				{0, 180, -180, 180, 180},
				{0, 190, -180, 180, 180},
				{0, -170, -180, 180, -170},
				{0, 190, -217, 217, 190},
			}

			for _, c := range cases {
				got := DegreeClampDelta(c.a, c.b, c.min, c.max)
				// fmt.Println(x, "got:[", got, "]", "expected [", c.exp, "]")
				Expect(got).To(Equal(c.exp))
			}
		})
	})
	Context("DegreeWrapCompare", func() {
		It("Should compare wrapped degrees correctly", func() {
			cases := []struct {
				a         float64
				b         float64
				tolerance float64
				exp       bool
			}{
				{90.0, 90.05, 0.1, true},
				{90.0, 90.11, 0.1, false},
				{90.0, 89.95, 0.1, true},
				{90.0, 89.89, 0.1, false},
				{179.9, 180.1, 1.0, true},
				{179.9, 181.1, 1.0, false},
				{-179.9, -180.1, 1.0, true},
				{-179.9, -181.1, 1.0, false},
				{179.9, -179.9, 1.0, true},
				{179.9, -178.9, 1.0, false},
				{-179.9, 179.9, 1.0, true},
				{-179.9, 178.9, 1.0, false},
				{179.9, 180.1, 1.0, true},
				{179.9, 181.1, 1.0, false},
				{-179.9, -180.1, 1.0, true},
				{-179.9, -181.1, 1.0, false},
				{179.9, -179.9, 1.0, true},
				{179.9, -178.9, 1.0, false},
				{-179.9, 179.9, 1.0, true},
				{-179.9, 178.9, 1.0, false},
				{179.9, 180.1, 1.0, true},
				{-179.9, -180.1, 1.0, true},
				{179.9, -179.9, 1.0, true},
				{-179.9, 179.9, 1.0, true},
				{-179.9, 180.1, 1.0, true},
				{179.9, -180.1, 1.0, true},
				{179.9, 179.9, 1.0, true},
				{-179.9, -179.9, 1.0, true},
				{-179.9, 180.1, 1.0, true},
				{179.9, -180.1, 1.0, true},
				{179.9, 179.9, 1.0, true},
				{-179.9, -179.9, 1.0, true},
				{-179.92, 179.99, 0.8, true},
				{-179.92, 179.99, 0.8, true},
				{90.0, -90, 1.0, false},
				{-90.0, 90, 1.0, false},
				{90.0, -90, 1.0, false},
				{-90.0, 90, 1.0, false},
				{90.0, -90, 1.0, false},
				{-90.0, 90, 1.0, false},
			}

			for _, c := range cases {
				got := DegreeWrapCompare(c.a, c.b, c.tolerance)
				// fmt.Println("got:[", got, "]", "expected [", c.exp, "]")
				Expect(got).To(Equal(c.exp))
			}
		})
	})
	Context("DegreeCompare", func() {
		It("Should compare degrees correctly", func() {
			cases := []struct {
				a         float64
				b         float64
				min       float64
				max       float64
				tolerance float64
				exp       bool
			}{
				{90.0, 90.05, -180, 180, 0.1, true},
				{90.0, 90.11, -180, 180, 0.1, false},
				{90.0, 89.95, -180, 180, 0.1, true},
				{90.0, 89.89, -180, 180, 0.1, false},
				{179.9, 180.1, -217, 217, 1.0, true},
				{179.9, 181.1, -180, 180, 1.0, false},
				{-179.9, -180.1, -217, 217, 1.0, true},
				{-179.9, -181.1, -180, 180, 1.0, false},
				{179.9, -179.9, -217, 217, 1.0, false},
				{179.9, -179.9, -180, 180, 1.0, false},
				{179.9, -178.9, -180, 180, 1.0, false},
				{-179.9, 179.9, -217, 217, 1.0, false},
				{-179.9, 179.9, -180, 180, 1.0, false},
				{-179.9, 178.9, -180, 180, 1.0, false},
				{179.9, 180.1, -217, 217, 1.0, true},
				{179.9, 181.1, -180, 180, 1.0, false},
				{-179.9, -180.1, -217, 217, 1.0, true},
				{-179.9, -181.1, -180, 180, 1.0, false},
				{179.9, -178.9, -180, 180, 1.0, false},
				{-179.9, 178.9, -180, 180, 1.0, false},
				{179.9, 180.1, -217, 217, 1.0, true},
				{-179.9, -180.1, -217, 217, 1.0, true},
				{-179.9, 180.1, -180, 180, 1.0, false},
				{179.9, -180.1, -180, 180, 1.0, false},
				{179.9, 179.9, -180, 180, 1.0, true},
				{-179.9, -179.9, -180, 180, 1.0, true},
				{-179.9, 180.1, -180, 180, 1.0, false},
				{179.9, -180.1, -180, 180, 1.0, false},
				{179.9, 179.9, -180, 180, 1.0, true},
				{-179.92, 179.99, -217, 217, 0.8, false},
				{-179.92, 179.99, -180, 180, 0.8, false},
				{90.0, -90, -180, 180, 1.0, false},
				{-90.0, 90, -180, 180, 1.0, false},
				{90.0, -90, -180, 180, 1.0, false},
				{-90.0, 90, -180, 180, 1.0, false},
				{90.0, -90, -180, 180, 1.0, false},
				{-90.0, 90, -180, 180, 1.0, false},
			}

			for _, c := range cases {
				got := DegreeCompare(c.a, c.b, c.min, c.max, c.tolerance)
				// fmt.Println(x, "got:[", got, "]", "expected [", c.exp, "]")
				Expect(got).To(Equal(c.exp))
			}
		})
	})
	Context("DegreeClampCompare", func() {
		It("Should compare clamped degrees correctly", func() {
			cases := []struct {
				a         float64
				b         float64
				min       float64
				max       float64
				tolerance float64
				exp       bool
			}{
				{90.0, 90.05, -180, 180, 0.1, true},
				{90.0, 90.11, -180, 180, 0.1, false},
				{90.0, 89.95, -180, 180, 0.1, true},
				{90.0, 89.89, -180, 180, 0.1, false},
				{179.9, 180.1, -217, 217, 1.0, true},
				{179.9, 181.1, -180, 180, 1.0, true},
				{-179.9, -180.1, -217, 217, 1.0, true},
				{-179.9, -181.1, -180, 180, 1.0, true},
				{179.9, -179.9, -217, 217, 1.0, false},
				{179.9, -179.9, -180, 180, 1.0, false},
				{179.9, -178.9, -180, 180, 1.0, false},
				{-179.9, 179.9, -217, 217, 1.0, false},
				{-179.9, 179.9, -180, 180, 1.0, false},
				{-179.9, 178.9, -180, 180, 1.0, false},
				{179.9, 180.1, -217, 217, 1.0, true},
				{179.9, 181.1, -180, 180, 1.0, true},
				{-179.9, -180.1, -217, 217, 1.0, true},
				{-179.9, -181.1, -180, 180, 1.0, true},
				{179.9, -178.9, -180, 180, 1.0, false},
				{-179.9, 178.9, -180, 180, 1.0, false},
				{179.9, 180.1, -217, 217, 1.0, true},
				{-179.9, -180.1, -217, 217, 1.0, true},
				{-179.9, 180.1, -180, 180, 1.0, false},
				{179.9, -180.1, -180, 180, 1.0, false},
				{179.9, 179.9, -180, 180, 1.0, true},
				{-179.9, -179.9, -180, 180, 1.0, true},
				{-179.9, 180.1, -180, 180, 1.0, false},
				{179.9, -180.1, -180, 180, 1.0, false},
				{179.9, 179.9, -180, 180, 1.0, true},
				{-179.92, 179.99, -217, 217, 0.8, false},
				{-179.92, 179.99, -180, 180, 0.8, false},
				{90.0, -90, -180, 180, 1.0, false},
				{-90.0, 90, -180, 180, 1.0, false},
				{90.0, -90, -180, 180, 1.0, false},
				{-90.0, 90, -180, 180, 1.0, false},
				{90.0, -90, -180, 180, 1.0, false},
				{-90.0, 90, -180, 180, 1.0, false},
			}

			for _, c := range cases {
				got := DegreeClampCompare(c.a, c.b, c.min, c.max, c.tolerance)
				// fmt.Println(x, "got:[", got, "]", "expected [", c.exp, "]")
				Expect(got).To(Equal(c.exp))
			}
		})
	})
})

func TestDegrees(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DegreesTests")
}
