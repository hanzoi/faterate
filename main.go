package main

import "fmt"

func main() {
	//先定义三人的姓名、性别、身高、体重、年龄信息
	var name [3]string
	var sex [3]string
	var tall [3]float64
	var weight [3]float64
	var age [3]int

	//为三人的各项信息赋值
	for i := 0; i < 3; i++ {
		fmt.Printf("请输入第%d个人的姓名：", i+1)
		fmt.Scanln(&name[i])
		fmt.Printf("请输入第%d个人的性别：", i+1)
		fmt.Scanln(&sex[i])
		fmt.Printf("请输入第%d个人的身高：", i+1)
		fmt.Scanln(&tall[i])
		fmt.Printf("请输入第%d个人的体重：", i+1)
		fmt.Scanln(&weight[i])
		fmt.Printf("请输入第%d个人的年龄：", i+1)
		fmt.Scanln(&age[i])
	}
	fmt.Println()

	//根据三人的性别确定性别权重
	var sexWeight [3]int
	for i := 0; i < 3; i++ {
		if sex[i] == "男" {
			sexWeight[i] = 1
		} else if sex[i] == "女" {
			sexWeight[i] = 0
		} else {
			fmt.Println("性别错误，请输入正确的性别")
			continue
		}
	}

	//分别计算三人的BMI、体脂率
	var bmi [3]float64
	var fatRate [3]float64
	for i := 0; i < 3; i++ {
		bmi[i] = weight[i] / (tall[i] * tall[i])
		fatRate[i] = (1.2*bmi[i] + 0.23*float64(age[i]) - 5.4 - 10.8*float64(sexWeight[i])) / 100
	}

	//分别为三人匹配体脂率建议
	var suggest [3]string
	for i := 0; i < 3; i++ {
		if sex[i] == "男" {
			//编写男性体脂状态表
			if age[i] >= 18 && age[i] <= 39 {
				if fatRate[i] <= 0.1 {
					suggest[i] = "偏瘦，多吃多锻炼，增强体质"
				} else if fatRate[i] > 0.1 && fatRate[i] <= 0.16 {
					suggest[i] = "标准，太棒了，要保持哦"
				} else if fatRate[i] > 0.16 && fatRate[i] <= 0.21 {
					suggest[i] = "偏胖，少吃晚饭多散步，加油哦"
				} else if fatRate[i] > 0.21 && fatRate[i] <= 0.26 {
					suggest[i] = "肥胖，少吃多运动"
				} else {
					suggest[i] = "非常肥胖，建议医院检查"
				}
			} else if age[i] >= 40 && age[i] <= 59 {
				if fatRate[i] <= 0.11 {
					suggest[i] = "偏瘦，多吃多锻炼，增强体质"
				} else if fatRate[i] > 0.12 && fatRate[i] <= 0.17 {
					suggest[i] = "标准，太棒了，要保持哦"
				} else if fatRate[i] > 0.18 && fatRate[i] <= 0.22 {
					suggest[i] = "偏胖，少吃晚饭多散步，加油哦"
				} else if fatRate[i] > 0.22 && fatRate[i] <= 0.27 {
					suggest[i] = "肥胖，少吃多运动"
				} else {
					suggest[i] = "非常肥胖，建议医院检查"
				}
			} else if age[i] >= 60 {
				if fatRate[i] <= 0.13 {
					suggest[i] = "偏瘦，多吃多锻炼，增强体质"
				} else if fatRate[i] > 0.13 && fatRate[i] <= 0.19 {
					suggest[i] = "标准，太棒了，要保持哦"
				} else if fatRate[i] > 0.19 && fatRate[i] <= 0.24 {
					suggest[i] = "偏胖，少吃晚饭多散步，加油哦"
				} else if fatRate[i] > 0.24 && fatRate[i] <= 0.29 {
					suggest[i] = "肥胖，少吃多运动"
				} else {
					suggest[i] = "非常肥胖，建议医院检查"
				}
			} else {
				suggest[i] = "抱歉，未成年人体脂率无法计算"
			}
		} else if sex[i] == "女" {
			if age[i] >= 18 && age[i] <= 39 {
				if fatRate[i] <= 0.2 {
					suggest[i] = "偏瘦，多吃多锻炼，增强体质"
				} else if fatRate[i] > 0.2 && fatRate[i] <= 0.27 {
					suggest[i] = "标准，太棒了，要保持哦"
				} else if fatRate[i] > 0.27 && fatRate[i] <= 0.34 {
					suggest[i] = "偏胖，少吃晚饭多散步，加油哦"
				} else if fatRate[i] > 0.34 && fatRate[i] <= 0.39 {
					suggest[i] = "肥胖，少吃多运动"
				} else {
					suggest[i] = "非常肥胖，建议医院检查"
				}
			} else if age[i] >= 40 && age[i] <= 59 {
				if fatRate[i] <= 0.21 {
					suggest[i] = "偏瘦，多吃多锻炼，增强体质"
				} else if fatRate[i] > 0.21 && fatRate[i] <= 0.28 {
					suggest[i] = "标准，太棒了，要保持哦"
				} else if fatRate[i] > 0.28 && fatRate[i] <= 0.35 {
					suggest[i] = "偏胖，少吃晚饭多散步，加油哦"
				} else if fatRate[i] > 0.35 && fatRate[i] <= 0.4 {
					suggest[i] = "肥胖，少吃多运动"
				} else {
					suggest[i] = "非常肥胖，建议医院检查"
				}
			} else if age[i] >= 60 {
				if fatRate[i] <= 0.22 {
					suggest[i] = "偏瘦，多吃多锻炼，增强体质"
				} else if fatRate[i] > 0.22 && fatRate[i] <= 0.29 {
					suggest[i] = "标准，太棒了，要保持哦"
				} else if fatRate[i] > 0.29 && fatRate[i] <= 0.36 {
					suggest[i] = "偏胖，少吃晚饭多散步，加油哦"
				} else if fatRate[i] > 0.36 && fatRate[i] <= 0.41 {
					suggest[i] = "肥胖，少吃多运动"
				} else {
					suggest[i] = "非常肥胖，建议医院检查"
				}
			} else {
				suggest[i] = "抱歉，未成年人体脂率无法计算"
			}
		}
	}

	//输出三人的姓名、BMI、体脂率、建议
	for i := 0; i < 3; i++ {
		fmt.Printf("%s 的 bmi 为 %.2f 体脂率为 %.2f 【%s】\n", name[i], bmi[i], fatRate[i], suggest[i])
	}
	fmt.Println()

	//输出总人数、平均体脂率
	var sum int
	var fatRateSum float64
	for i, v := range fatRate {
		fatRateSum += v
		sum += i
	}
	fmt.Printf("总人数为 %d，平均体脂率为 %f", sum, fatRateSum/float64(sum))
}
