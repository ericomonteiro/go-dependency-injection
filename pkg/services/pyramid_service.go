package services

type PyramidService struct{}

func NewPyramidService() *PyramidService {
	return &PyramidService{}
}

func (p *PyramidService) PrintPyramidWithLines(lines int) {
	for i := 0; i < lines; i++ {
		for j := 0; j < lines-i; j++ {
			print(" ")
		}
		for j := 0; j <= i; j++ {
			print("* ")
		}
		println()
	}
}
