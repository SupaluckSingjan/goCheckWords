package sayword 

import "fmt" 

func SayWord(x string) string{
	var sayWord string = x

	switch sayWord {
	case "Monday":
		fmt.Println("วันจันทร์")
	case "Tuesday":
		fmt.Println("วันอังคาร")
	case "Wednesday":
		fmt.Println("วันพุธ")
	case "Thursday":
		fmt.Println("วันพฤหัส")
	case "Friday":
		fmt.Println("วันศุกร์")
	case "Saturday":
		fmt.Println("วันเสาร์")
	case "Sanday":
		fmt.Println("วันอาทิตย์")
	default:
		fmt.Println("ไม่มีคำพูดที่ตรงกัน")
	}

	return x

}