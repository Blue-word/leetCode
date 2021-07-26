package main

import (
	"fmt"
	"math"
	"strconv"
)

func gcd(a int, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func stringRev(s string) string {
	str := []rune(s)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}

func intRev(a int) string {
	str := strconv.Itoa(a)
	strs := []rune(str)
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	return string(strs)
}

func recursion(n int) int {
	if n == 2 {
		return 1
	} else if n < 2 {
		return 0
	}
	return n/3 + recursion(n/3+n%3)
}

func tuzi(n int) int {
	if n < 3 {
		return 1
	}
	n1, n2 := 0, 1
	for i := 0; i < n-2; i++ {
		n1, n2 = n2, n1+n2
	}
	return n1 + n2
}

func doLast(stack []string) int {
	var tempS []string
	for i := 0; i < len(stack); i++ {
		cur := stack[i]
		if cur == "+" || cur == "-" || cur == "*" || cur == "/" {
			num1, _ := strconv.Atoi(tempS[len(tempS)-1])
			num2, _ := strconv.Atoi(tempS[len(tempS)-2])
			tempS = tempS[:len(tempS)-2]
			var numR int
			switch cur {
			case "+":
				numR = num2 + num1
			case "-":
				numR = num2 - num1
			case "*":
				numR = num2 * num1
			case "/":
				numR = num2 / num1
			}
			tempS = append(tempS, strconv.Itoa(numR))
		} else {
			tempS = append(tempS, cur)
		}
	}
	if len(tempS) > 0 {
		tempR, _ := strconv.Atoi(tempS[len(tempS)-1])
		return tempR
	}
	return 0
}

func midToLast(oldStr string) []string {
	var str string
	for i := 0; i < len(oldStr); i++ {
		if oldStr[i] == '{' {
			str += "("
		} else if oldStr[i] == '}' {
			str += ")"
		} else if oldStr[i] == '[' {
			str += "("
		} else if oldStr[i] == ']' {
			str += ")"
		} else if oldStr[i] == '-' && i == 0 {
			str += "0-"
		} else if oldStr[i] == '-' && i > 0 {
			if (oldStr[i-1] < '0' || oldStr[i-1] > '9') && oldStr[i-1] != ')' &&
				oldStr[i-1] != ']' && oldStr[i-1] != '}' {
				str += "0-"
			} else {
				str += string(oldStr[i])
			}
		} else {
			str += string(oldStr[i])
		}
	}
	var stack1 []string // 符号
	var stack2 []string // 数字
	var i int
Loop:
	for i < len(str) {
		cur := str[i]
		asd := string(cur)
		fmt.Println(asd)
		switch {
		case cur >= '0' && cur <= '9':
			var tempB []byte
			for ; i < len(str); i++ {
				if str[i] < '0' || str[i] > '9' {
					i--
					break
				}
				tempB = append(tempB, str[i])
			}
			stack2 = append(stack2, string(tempB))
		case cur == '(':
			stack1 = append(stack1, string(cur))
		case cur == ')':
			for len(stack1) > 0 {
				tempS := stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-1]
				if tempS == "(" {
					break
				}
				stack2 = append(stack2, tempS)
			}
		default:
			var s1 string
			if len(stack1) > 0 {
				s1 = stack1[len(stack1)-1]
			}
			if s1 == "" {
				stack1 = append(stack1, string(cur))
			} else if s1 == "(" || s1 == ")" {
				stack1 = append(stack1, string(cur))
			} else if priority(string(cur)) <= priority(s1) {
				for len(stack1) > 0 && priority(string(cur)) <= priority(stack1[len(stack1)-1]) {
					stack2 = append(stack2, stack1[len(stack1)-1])
					stack1 = stack1[:len(stack1)-1]
				}
				stack1 = append(stack1, string(cur))
			} else if priority(string(cur)) > priority(s1) {
				stack1 = append(stack1, string(cur))
			}
		}
		i++
		goto Loop
	}
	for len(stack1) > 0 {
		stack2 = append(stack2, stack1[len(stack1)-1])
		stack1 = stack1[:len(stack1)-1]
	}
	return stack2
}

func priority(s string) int {
	switch s {
	case "+":
		return 1
	case "-":
		return 1
	case "*":
		return 2
	case "/":
		return 2
	}
	return 0
}

func recursion1(n, m int) int {
	if n == 0 || m == 0 {
		return 1
	}
	return recursion1(n, m-1) + recursion1(n-1, m)
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func longestCommonSubsequence(str1, str2 string) int {
	dp := make([][]int, len(str1)+1)
	for i := 0; i <= len(str1); i++ {
		dp[i] = make([]int, len(str2)+1)
	}

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[len(str1)][len(str2)]
}

//最长公共子串是： dp[i][j] -- 表示以str1[i]和str2[j]为结尾的最长公共子串
//当str1[i] == str2[j]时，dp[i][j] = dp[i-1][j-1] + 1; 否则，dp[i][j] = 0;
//最优解为max(dp[i][j]),其中0<=i<len1, 0<=j<len2;
func longestCommonSubsequence1(str1, str2 string) int {
	dp := make([][]int, len(str1)+1)
	for i := 0; i <= len(str1); i++ {
		dp[i] = make([]int, len(str2)+1)
	}
	maxLen := 0
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return maxLen
}

func initTable(n, m int) [][]bool {
	data := make([][]bool, n)
	for i := 0; i < n; i++ {
		data[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			data[i][j] = true
		}
	}
	return data
}

func recursion2(m, n int) int {
	if m < 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	return recursion2(m-n, n) + recursion2(m, n-1)
}

func main() {
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//a := input.Text()
	//input.Scan()
	//b := input.Text()
	//longestCommonSubsequence(a, b)

	//a := 0
	//b := 0
	//for {
	//	n, _ := fmt.Scan(&a)
	//	if n == 0 {
	//		break
	//	} else {
	//		fmt.Scan(&b)
	//		fmt.Println(recursion2(a, b))
	//	}
	//}

	//m := 0
	//for {
	//	n, _ := fmt.Scan(&m)
	//	if n == 0 {
	//		break
	//	} else {
	//		var res []string
	//		a1 := m*m - m + 1
	//		for i := 0; i < m; i++ {
	//			res = append(res, strconv.Itoa(a1+2*i))
	//		}
	//		fmt.Println(strings.Join(res, "+"))
	//	}
	//}

	/*input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var num int
		bytes := []byte(input.Text())
		for i := 0; i < len(bytes); i++ {
			if bytes[i] >= 'A' && bytes[i] <= 'Z' {
				num++
			}
		}
		fmt.Println(num)
	}*/

	//lengthOfLongestSubstring("abcabcbb")
	minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"})
}

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func minOperations(logs []string) int {
	result := 0
	for _, v := range logs {
		op := v[:2]
		switch op {
		case "..":
			if result > 0 {
				result--
			}
		case "./":
		default:
			result++
		}
	}
	return result
}
