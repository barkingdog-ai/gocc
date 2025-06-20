package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/barkingdog-ai/gocc"
)

// 測試案例結構
type TestCase struct {
	Name        string
	Conversion  string
	Input       string
	Description string
}

func main() {
	// 定義各種測試案例
	testCases := []TestCase{
		{
			Name:        "簡體轉繁體",
			Conversion:  "s2t",
			Input:       "托福",
			Description: "托福",
		},
	}

	fmt.Println("=== GoCC 中文轉換測試腳本 ===")
	fmt.Println()

	// 統計資訊
	totalTests := len(testCases)
	successCount := 0
	failCount := 0

	for i, testCase := range testCases {
		fmt.Printf("測試 %d/%d: %s\n", i+1, totalTests, testCase.Name)
		fmt.Printf("轉換模式: %s (%s)\n", testCase.Conversion, testCase.Description)
		fmt.Printf("原始文字: %s\n", testCase.Input)

		// 執行轉換
		converter, err := gocc.New(testCase.Conversion)
		if err != nil {
			fmt.Printf("❌ 建立轉換器失敗: %v\n", err)
			failCount++
		} else {
			output, err := converter.Convert(testCase.Input)
			if err != nil {
				fmt.Printf("❌ 轉換失敗: %v\n", err)
				failCount++
			} else {
				fmt.Printf("轉換結果: %s\n", output)
				fmt.Printf("✅ 轉換成功\n")
				successCount++
			}
		}

		fmt.Println(strings.Repeat("-", 60))
		fmt.Println()
	}

	// 顯示測試結果統計
	fmt.Printf("=== 測試結果統計 ===\n")
	fmt.Printf("總測試數: %d\n", totalTests)
	fmt.Printf("成功: %d\n", successCount)
	fmt.Printf("失敗: %d\n", failCount)
	fmt.Printf("成功率: %.1f%%\n", float64(successCount)/float64(totalTests)*100)

	// 互動模式
	if len(os.Args) > 1 && os.Args[1] == "-i" {
		fmt.Println("\n=== 互動測試模式 ===")
		runInteractiveMode()
	}
}

// 互動測試模式
func runInteractiveMode() {
	availableConversions := []string{
		"s2t", "t2s", "s2tw", "tw2s", "s2hk", "hk2s",
		"s2twp", "tw2sp", "t2tw", "t2hk",
	}

	fmt.Println("可用的轉換模式:")
	for i, conv := range availableConversions {
		fmt.Printf("%d. %s\n", i+1, conv)
	}

	fmt.Print("\n請選擇轉換模式 (輸入數字 1-10，或輸入 'quit' 退出): ")

	var choice int
	var input string

	for {
		fmt.Scanln(&input)
		if input == "quit" {
			fmt.Println("退出互動模式")
			break
		}

		// 簡單的數字解析
		if len(input) == 1 && input[0] >= '1' && input[0] <= '9' {
			choice = int(input[0] - '0')
			if choice >= 1 && choice <= len(availableConversions) {
				selectedConversion := availableConversions[choice-1]
				fmt.Printf("\n你選擇了: %s\n", selectedConversion)

				fmt.Print("請輸入要轉換的文字: ")
				var text string
				fmt.Scanln(&text)

				if text != "" {
					converter, err := gocc.New(selectedConversion)
					if err != nil {
						fmt.Printf("❌ 建立轉換器失敗: %v\n", err)
					} else {
						output, err := converter.Convert(text)
						if err != nil {
							fmt.Printf("❌ 轉換失敗: %v\n", err)
						} else {
							fmt.Printf("轉換結果: %s\n", output)
						}
					}
				}
			}
		}

		fmt.Print("\n請選擇轉換模式 (輸入數字 1-10，或輸入 'quit' 退出): ")
	}
}
