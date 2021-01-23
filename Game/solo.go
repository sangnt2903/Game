package main

import (
	"fmt"
	"math/rand"
	"time"
)

var c1 = OCharacter{
	Name:        "c1",
	Healthy:     1,
	Defense:     2,
	Attack:      1,
	AttackSpeed: 1,
	Critical:    0,
	Miss:        0,
}

var c2 = OCharacter{
	Name:        "c2",
	Healthy:     1,
	Defense:     2,
	Attack:      1,
	AttackSpeed: 1,
	Critical:    0,
	Miss:        0,
}

/*var c2 = OCharacter{
	Name: "c2",
	Healthy: 500,
	Defense: 1500,
	Attack: 110,
	AttackSpeed: 1,
	Critical: 80,
	Miss: 1,
}*/

func main() {
	done := make(chan bool)
	go Game(&c1, &c2, done)
	go Game(&c2, &c1, done)
	<-done
}

func Game(c *OCharacter, target *OCharacter, done chan bool) {
	for c.Healthy > 0 {
		c.CAttack(target)
		if target.Healthy <= 0 {
			fmt.Println(fmt.Sprintf("%s Killed %s", c.Name, target.Name))
			//fmt.Println(fmt.Sprintf("---VICTORY---%s-------", c.Name))
			done <- true
			return
		}
		t := time.Duration(1.0 / c.AttackSpeed)
		time.Sleep(t * time.Second)
	}
}

type ICharacter interface {
	CAttack(target *OCharacter)
}

type OCharacter struct {
	Name        string
	Healthy     float32
	Defense     int32
	Attack      float32
	AttackSpeed float32
	Critical    float32
	Miss        float32
}

func (O *OCharacter) CAttack(target *OCharacter) {
	if target.Healthy > 0 {
		target.Healthy -= O.GetAttack(*target)
		fmt.Println(fmt.Sprintf("%s Healthy remaing : %f", target.Name, target.Healthy))
	}
}

func (O OCharacter) GetAttack(target OCharacter) float32 {
	if O.IsMissTarget(target) {
		fmt.Println(fmt.Sprintf("%s missed %s", O.Name, target.Name))
		return 0
	}
	if O.IsCriticalTarget() {
		fmt.Println(fmt.Sprintf("%s criticaled %s damages %f", O.Name, target.Name, 2*(O.Attack-O.GetDefense(target))))
		return 2 * (O.Attack - O.GetDefense(target))
	}
	fmt.Println(fmt.Sprintf("%s damaged %s  %f", O.Name, target.Name, O.Attack-O.GetDefense(target)))
	return O.Attack - O.GetDefense(target)
}

func (O OCharacter) GetDefense(target OCharacter) float32 {
	rate := float32(target.Defense) / O.Attack
	if rate > 1 {
		return O.Attack * 0.5
	} else if rate == 1 {
		return O.Attack * 0.3
	} else {
		return O.Attack * 0.15
	}
}

func (O OCharacter) IsCriticalTarget() bool {
	return int(O.Critical) > rand.Intn(100)+1
}

func (O OCharacter) IsMissTarget(target OCharacter) bool {
	return int(target.Miss) > rand.Intn(100)+1
}

var _ ICharacter = &OCharacter{}
