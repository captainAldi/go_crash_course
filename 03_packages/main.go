package main

//untuk multi import tambahkan () dan separator nya spasi
//taruh di line berbeda untuk formating yg benar
import (
	"fmt"
	"math"

	"github.com/captainAldi/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}
