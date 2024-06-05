package svg

import (
	"LuxeGo/internal/lx"
)

//FeConvolveMatrix - 
func FeConvolveMatrix() *FeConvolveMatrixTag {
	return &FeConvolveMatrixTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "feconvolvematrix", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeConvolveMatrixTag struct {
	*ComponentHtmlTag
}

//In - 
func (f *FeConvolveMatrixTag) In(value string) *FeConvolveMatrixTag {
	f.AddAttribute("in", value)
	return f
}

//Order - 
func (f *FeConvolveMatrixTag) Order(value string) *FeConvolveMatrixTag {
	f.AddAttribute("order", value)
	return f
}

//KernelMatrix - 
func (f *FeConvolveMatrixTag) KernelMatrix(value string) *FeConvolveMatrixTag {
	f.AddAttribute("kernelmatrix", value)
	return f
}

//Divisor - 
func (f *FeConvolveMatrixTag) Divisor(value string) *FeConvolveMatrixTag {
	f.AddAttribute("divisor", value)
	return f
}

//Bias - 
func (f *FeConvolveMatrixTag) Bias(value string) *FeConvolveMatrixTag {
	f.AddAttribute("bias", value)
	return f
}

//TargetX - 
func (f *FeConvolveMatrixTag) TargetX(value string) *FeConvolveMatrixTag {
	f.AddAttribute("targetx", value)
	return f
}

//TargetY - 
func (f *FeConvolveMatrixTag) TargetY(value string) *FeConvolveMatrixTag {
	f.AddAttribute("targety", value)
	return f
}

//EdgeMode - 
func (f *FeConvolveMatrixTag) EdgeMode(value string) *FeConvolveMatrixTag {
	f.AddAttribute("edgemode", value)
	return f
}

//KernelUnitLength - 
func (f *FeConvolveMatrixTag) KernelUnitLength(value string) *FeConvolveMatrixTag {
	f.AddAttribute("kernelunitlength", value)
	return f
}

//PreserveAlpha - 
func (f *FeConvolveMatrixTag) PreserveAlpha(value string) *FeConvolveMatrixTag {
	f.AddAttribute("preservealpha", value)
	return f
}
