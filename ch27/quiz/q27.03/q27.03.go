package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Attacker interface {
	Name() string
}

type Defender interface {
	DealDamage(att Attacker, damage int)
}

type Player struct {
	name string
	hp   int
}

type Monster struct {
	name string
	hp   int
}

func (p *Player) Name() string {
	return p.name
}

func (m *Monster) Name() string {
	return m.name
}

func (p *Player) Attack(d Defender) {
	d.DealDamage(p, 100)
}

func (p *Player) DealDamage(att Attacker, damage int) {
	p.hp -= damage
	if p.hp <= 0 {
		fmt.Printf("%s가 플레이어 %s를 죽였습니다\n", att.Name(), p.Name())
	}
}

func (m *Monster) Attack(d Defender) {
	d.DealDamage(m, 100)
}

func (m *Monster) DealDamage(att Attacker, damage int) {
	m.hp -= damage
	if m.hp <= 0 {
		fmt.Printf("플레이어 %s가 몬스터 %s를 죽였습니다\n", att.Name(), m.Name())
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := &Player{"Go", 500}
	m := &Monster{"SOLID", 500}

	for {
		randValue := rand.Intn(2)
		if p.hp > 0 && m.hp > 0 {
			fmt.Printf("Player: %s (%d)\tMonster: %s (%d)\n", p.name, p.hp, m.name, m.hp)
			if randValue == 0 {
				p.Attack(m)
			} else {
				m.Attack(p)
			}
		} else {
			break
		}
	}
}
