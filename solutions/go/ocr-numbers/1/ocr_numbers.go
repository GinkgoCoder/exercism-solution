package ocr

import (
    "strings"
)

// 定义0-9的数字模式（每个数字3列，前3行）
var digitPatterns = map[string]string{
    " _ " + "| |" + "|_|": "0",
    "   " + "  |" + "  |": "1",
    " _ " + " _|" + "|_ ": "2",
    " _ " + " _|" + " _|": "3",
    "   " + "|_|" + "  |": "4",
    " _ " + "|_ " + " _|": "5",
    " _ " + "|_ " + "|_|": "6",
    " _ " + "  |" + "  |": "7",
    " _ " + "|_|" + "|_|": "8",
    " _ " + "|_|" + " _|": "9",
}

// 识别单个数字（3列x3行）
func recognizeDigit(top, middle, bottom string) string {
    pattern := top + middle + bottom
    if digit, ok := digitPatterns[pattern]; ok {
        return digit
    }
    return "?"
}

// Recognize 识别整个输入
func Recognize(input string) []string {
    lines := strings.Split(input, "\n")
    
    // 如果第一行是空的（因为输入以\n开头），跳过它
    if len(lines) > 0 && lines[0] == "" {
        lines = lines[1:]
    }
    
    // 验证输入：行数必须是4的倍数
    if len(lines)%4 != 0 {
        return []string{}
    }
    
    var results []string
    
    // 每4行处理一个数字行
    for i := 0; i < len(lines); i += 4 {
        line1 := lines[i]
        line2 := lines[i+1]
        line3 := lines[i+2]
        // lines[i+3] 是空行
        
        // 补齐行长度（以最长的为准）
        maxLen := max(len(line1), len(line2), len(line3))
        
        // 验证：每行长度必须是3的倍数
        if maxLen%3 != 0 {
            return []string{}
        }
        
        line1 = padRight(line1, maxLen)
        line2 = padRight(line2, maxLen)
        line3 = padRight(line3, maxLen)
        
        // 识别这一行的所有数字
        numDigits := maxLen / 3
        var rowResult string
        
        for j := 0; j < numDigits; j++ {
            col := j * 3
            top := line1[col : col+3]
            middle := line2[col : col+3]
            bottom := line3[col : col+3]
            
            rowResult += recognizeDigit(top, middle, bottom)
        }
        
        results = append(results, rowResult)
    }
    
    return results
}

// 辅助函数：右侧填充空格
func padRight(s string, length int) string {
    if len(s) >= length {
        return s
    }
    return s + strings.Repeat(" ", length-len(s))
}

// 辅助函数：获取最大值
func max(a, b, c int) int {
    result := a
    if b > result {
        result = b
    }
    if c > result {
        result = c
    }
    return result
}
