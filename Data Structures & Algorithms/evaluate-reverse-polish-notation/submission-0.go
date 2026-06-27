func evalRPN(tokens []string) int {
   for len(tokens) > 1 {
       for i := 0; i < len(tokens); i++ {
           if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
               a, _ := strconv.Atoi(tokens[i-2])
               b, _ := strconv.Atoi(tokens[i-1])
               var result int

               switch tokens[i] {
               case "+":
                   result = a + b
               case "-":
                   result = a - b
               case "*":
                   result = a * b
               case "/":
                   result = a / b
               }

               temp := make([]string, 0)
               temp = append(temp, tokens[:i-2]...)
               temp = append(temp, strconv.Itoa(result))
               temp = append(temp, tokens[i+1:]...)
               tokens = temp
               break
           }
       }
   }

   result, _ := strconv.Atoi(tokens[0])
   return result
}