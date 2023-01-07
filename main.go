package main

import (
	"bufio"
  "fmt"
  "os"
  "strings"
  "errors"
)

func main() {
  //Считываем выражение из консоли и определяем оператор и операнды
  operator, operandAStr,  operandBStr:=initString()
 //Инициализируем первый операнд
  operandA, err,nA:=operand(operandAStr)
  //Инициализируем второй операнд
  operandB, err1, nB:=operand(operandBStr)
  
if err!=nil || err1!=nil {
    fmt.Println("Ошибка! Введен неверный операнд.")
    os.Exit(0)
    }
//Определяем систему счисления, производим вычисление результата и выполняем нужные преобразования
  n:=nA+nB
  if n == 2 || n== 4 {
  result:=operate(operator, operandA,  operandB)
    if n==4 {
      if result>0 {
     arabToRoman(result)
      } else {
      fmt.Println("Ошибка! В римской системе нет отрицательных чисел.")
    }
   } else if n==2 {
     fmt.Println("Результат: ", result)
     }
   } else {
  fmt.Println("Ошибка! Используются одновременно разные системы счисления ")
  }
  os.Exit(0)
}

func initString()(string, string, string){
var operator, operandA, operandB string
var reader = bufio.NewReader(os.Stdin)
  fmt.Println ("Введите выражение: ")
  text, _ := reader.ReadString('\n')
  text=strings.ReplaceAll(text, " ", "")
  text= strings.ToUpper(text)
  for i:=0; i < len(text); i++ {
  if string(text[i]) == "+" ||
  string(text[i]) == "-"  ||
  string(text[i]) == "*"  ||
  string(text[i]) == "/"   {
    operator= string(text[i])
    n:= i
    for i=0; i<n; i++ {
      operandA+=string(text[i])  
      }
    for i=n+1; i<(len(text)-1); i++ {
      operandB+=string(text[i]) 
      } 
    }
    }
  return operator, operandA, operandB
}

// выбор оператора и его проверка
func operand (m string) (int, error, int){
  switch m {
   case "1": return 1, nil, 1
   case "2" : return 2, nil, 1
   case "3" : return 3, nil, 1
   case "4" : return 4, nil, 1
   case "5" : return 5, nil, 1
   case "6": return 6, nil, 1
   case "7" : return 7, nil, 1
   case "8" : return 8, nil, 1
   case "9" : return 9, nil, 1
   case "10": return 10, nil, 1
   case "I" : return 1, nil, 2
   case "II" : return 2, nil, 2
   case "III" : return 3, nil, 2
   case "IV": return 4, nil, 2
   case "V" : return 5, nil, 2
   case "VI" : return 6, nil, 2
   case "VII" : return 7, nil, 2
   case "VIII": return 8, nil, 2
   case "IX" : return 9, nil, 2
   case "X" : return 10, nil, 2
   default: return -1, errors.New("Ошибка! Введен неверный оператор."), 0
    }
}

// выбор знака операции и его проверка
func operate (m string, a,b int) (int){
  var c int
  switch m {
   case "+": {c=a+b}
   case "-" : {c=a-b}
   case "*" : {c=a*b}
   case "/" : {c=a/b}
   default: {fmt.Println("Ошибка! Введено неверное выражение.")
     os.Exit(0)
     }
   }
return c 
}
  
func arabToRoman (result int)(){
var (
table = [][] string {{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
{"", "C"}})
var roman string
	if result >0 {
    digit := 100
    for i := 2; i >= 0; i-- {
		  d := result / digit
		  roman += table[i][d]
	  	result %= digit
		  digit /= 10
      }} else {
     fmt.Println("Ошибка! Выражение введено неверное.")
     os.Exit(0)
     }  
  fmt.Println ("Ответ: ",roman)
  os.Exit(0)
}