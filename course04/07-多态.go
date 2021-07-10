package main

import (
	"fmt"
)

type fighter interface {
	attack() (bloodloss int)
	defend()
}

// 骑兵
type rider struct{}

func (r *rider) attack() (bloodloss int) {
	fmt.Println("铁蹄滚滚，尸横遍野")
	return 100
}

func (r *rider) defend() {
	fmt.Println("骑兵防守")
}

// 弓箭手
type archer struct{}

func (a *archer) attack() (bloodloss int) {
	fmt.Println("射箭")
	return 50
}

func (a *archer) defend() {
	fmt.Println("弓箭手防守")
}

type master struct{}

func (m *master) attack() (bloodloss int) {
	fmt.Println("天灵灵地灵灵")
	return 50
}

func (m *master) defend() {
	fmt.Println("法师防守")
}

func main() {
	fighters := make([]fighter, 0)
	fighters = append(fighters, &rider{})
	fighters = append(fighters, &archer{})
	fighters = append(fighters, &master{})

	/*
		让用户发将令
		第一位数字代表骑兵，第二位数字代表弓箭手，第三位数字代表法师
		999 全体进攻
		000 全体防守
	*/
	var cmd string
	for {
		fmt.Println("大将军，请传将令：(输入exit退出)")
		_, _ = fmt.Scan(&cmd)

		switch {
		case cmd == "exit":
			goto GAMEOVER
		case cmd == "999":
			//fmt.Println("全体进攻")
			for _, v := range fighters {
				v.attack()
			}
		case cmd == "000":
			//fmt.Println("全体防守")
			for _, v := range fighters {
				v.defend()
			}
		default:
			for _, f := range fighters {
				if r, ok := f.(*rider); ok {
					if cmd[0] == '9' {
						r.attack()
					} else {
						r.defend()
					}
				}
				if a, ok := f.(*archer); ok {
					if cmd[1] == '9' {
						a.attack()
					} else {
						a.defend()
					}
				}
				if m, ok := f.(*master); ok {
					if cmd[2] == '9' {
						m.attack()
					} else {
						m.defend()
					}
				}
			}
		}
	}
	GAMEOVER:fmt.Println("GAMEOVER")
}

//type Hoop struct {
//	name string
//}
//
//func (h *Hoop) born() {
//	fmt.Println(h.name + "出生了")
//}
//
//func (h *Hoop) die() {
//	fmt.Println(h.name + "死了")
//}

//func main01() {
//	var hoop = Hoop{"lili"}
//	hoop.born()
//	hoop.die()
//	// 统一包下的变量访问要同时编译
//	Say()
//	fmt.Println(Name)
//	time.Sleep(10*time.Second)
//}
