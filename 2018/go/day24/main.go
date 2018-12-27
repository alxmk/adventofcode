package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	inputRegexp  = regexp.MustCompile(`(?P<units>[0-9]+) units each with (?P<hp>[0-9]+) hit points (\((immune to (?P<immunities>[a-z\ ,]+)[; ]*)?(weak to (?P<weaknesses>[a-z\ ,]+))?\) )?with an attack that does (?P<attack>[0-9]+) (?P<attacktype>[a-z]+) damage at initiative (?P<initiative>[0-9]+)`)
	inputRegexp2 = regexp.MustCompile(`(?P<units>[0-9]+) units each with (?P<hp>[0-9]+) hit points (\(weak to (?P<weaknesses>[a-z\ ,]+)[; ]*)?((immune to (?P<immunities>[a-z\ ,]+))?\) )?with an attack that does (?P<attack>[0-9]+) (?P<attacktype>[a-z]+) damage at initiative (?P<initiative>[0-9]+)`)
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	a := generateUnits(strings.Split(string(data), "\n"), 0)

	for _, g := range a {
		fmt.Println(g.String())
	}

	for {
		fmt.Println(a.String())
		if a.AttackPhase(a.TargetSelectionPhase()) {
			break
		}
	}

	winner, remaining := a.Result()
	log.Println("Part one:", winner, remaining)

	for boost := 30; ; boost++ {
		log.Println("Boost:", boost)
		a := generateUnits(strings.Split(string(data), "\n"), boost)

		for {
			fmt.Println(a.String())
			var phases int
			if a.AttackPhase(a.TargetSelectionPhase()) {
				fmt.Println("Phase", phases)
				break
			}
			phases++
		}

		if winner, remaining := a.Result(); winner == "Immune System:" {
			log.Println("Part two:", remaining)
			break
		}
	}
}

func generateUnits(input []string, boost int) army {
	var a army

	var currentTeam string
	var groupnum int
	var currentBoost int
	for _, line := range input {
		switch line {
		case "":
			continue
		case "Immune System:":
			currentTeam = line
			groupnum = 0
			currentBoost = boost
		case "Infection:":
			currentTeam = line
			groupnum = 0
			currentBoost = 0
		default:
			groupnum++
			groups := inputRegexp.FindStringSubmatch(line)
			// If we have immunities then we'll have 8 groups including whole match
			if len(groups) == 11 {
				a = append(a, &group{
					units:      mustIntConv(groups[1]),
					hp:         mustIntConv(groups[2]),
					immunities: strings.Split(groups[5], ", "),
					weaknesses: strings.Split(groups[7], ", "),
					attack:     mustIntConv(groups[8]) + currentBoost,
					attacktype: groups[9],
					initiative: mustIntConv(groups[10]),
					team:       currentTeam,
					num:        groupnum,
				})
			} else {
				groups = inputRegexp2.FindStringSubmatch(line)
				// fmt.Println(strings.Join(groups, "; "))
				if len(groups) == 11 {
					a = append(a, &group{
						units:      mustIntConv(groups[1]),
						hp:         mustIntConv(groups[2]),
						weaknesses: strings.Split(groups[4], ", "),
						immunities: strings.Split(groups[7], ", "),
						attack:     mustIntConv(groups[8]) + currentBoost,
						attacktype: groups[9],
						initiative: mustIntConv(groups[10]),
						team:       currentTeam,
						num:        groupnum,
					})
				} else {
					panic(fmt.Sprintf("Matched %d groups (%s) on line: %s", len(groups), strings.Join(groups, "; "), line))
				}
			}
		}
	}

	return a
}

func mustIntConv(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

type group struct {
	units      int
	hp         int
	attack     int
	attacktype string
	weaknesses []string
	immunities []string
	initiative int
	team       string
	num        int
}

func (g *group) EffectivePower() int {
	return g.units * g.attack
}

func (g *group) WeakTo(attacktype string) bool {
	for _, w := range g.weaknesses {
		if w == attacktype {
			return true
		}
	}

	return false
}

func (g *group) ImmuneTo(attacktype string) bool {
	for _, i := range g.immunities {
		if i == attacktype {
			return true
		}
	}

	return false
}

func (g *group) Alive() bool {
	return g.units > 0
}

func (g *group) CalculateDamage(target *group) int {
	if target.WeakTo(g.attacktype) {
		return g.EffectivePower() * 2
	}
	if target.ImmuneTo(g.attacktype) {
		return 0
	}
	return g.EffectivePower()
}

func (g *group) Attack(target *group) int {
	damage := g.CalculateDamage(target)
	before := target.units
	// Remove a number of units from the target group based on the damage
	target.units -= damage / target.hp

	// Can't drop below zero
	if target.units < 0 {
		target.units = 0
	}

	return before - target.units

	// fmt.Printf("%s %d attacks defending group %d for %d, killing %d units\n", g.team, g.num, target.num, damage, before-target.units)
}

func (g *group) SelectTarget(enemies army, targets map[*group]*group) (*group, int) {
	var mostDamage int
	var c1 []*group

	// Pick the enemies that the group would do most damage to
outer:
	for _, e := range enemies {
		// Skip our own team and dead bois
		if e.team == g.team || !e.Alive() {
			continue
		}
		// Can't be targetted twice
		for _, t := range targets {
			if e == t {
				continue outer
			}
		}
		damage := g.CalculateDamage(e)
		// Skip any we would do zero damage to
		if damage == 0 {
			continue
		}
		if damage > mostDamage {
			mostDamage = damage
			c1 = []*group{e}
			continue
		}
		if damage == mostDamage {
			c1 = append(c1, e)
		}
	}

	// If we wouldn't do any damage then return nothing
	if len(c1) == 0 {
		return nil, 0
	}

	// If there's no tie then return the group
	if len(c1) == 1 {
		return c1[0], mostDamage
	}

	// Otherwise need to select on effective power
	var mostEP int
	var c2 []*group
	for _, e := range c1 {
		power := e.EffectivePower()
		if power > mostEP {
			mostEP = power
			c2 = []*group{e}
			continue
		}
		if power == mostEP {
			c2 = append(c2, e)
		}
	}

	// If there's no tie then return the group
	if len(c2) == 1 {
		return c2[0], mostDamage
	}

	// Finally select on initiative
	var mostInitiative int
	var c3 []*group
	for _, e := range c2 {
		initiative := e.initiative
		if initiative > mostInitiative {
			mostInitiative = initiative
			c3 = []*group{e}
			continue
		}
		if initiative == mostInitiative {
			c3 = append(c3, e)
		}
	}

	return c3[0], mostDamage
}

type army []*group

type byPower army

func (a byPower) Len() int      { return len(a) }
func (a byPower) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byPower) Less(i, j int) bool {
	if a[i].EffectivePower() == a[j].EffectivePower() {
		return a[i].initiative < a[j].initiative
	}
	return a[i].EffectivePower() < a[j].EffectivePower()
}

type byInitiative army

func (a byInitiative) Len() int      { return len(a) }
func (a byInitiative) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byInitiative) Less(i, j int) bool {
	return a[i].initiative < a[j].initiative
}

func (a army) TargetSelectionPhase() map[*group]*group {
	sort.Sort(sort.Reverse(byPower(a)))

	targets := make(map[*group]*group)
	for _, u := range a {
		if t, _ := u.SelectTarget(a, targets); t != nil {
			targets[u] = t
		}
	}

	return targets
}

func (a army) AttackPhase(targets map[*group]*group) bool {
	printTargets(targets)
	if len(targets) == 0 {
		return true
	}

	sort.Sort(sort.Reverse(byInitiative(a)))

	var unitskilled int

	for _, u := range a {
		// Skip dead bois
		if !u.Alive() {
			continue
		}
		if target := targets[u]; target != nil {
			unitskilled += u.Attack(target)
		}
	}

	return a.TeamsRemaining() == 1 || unitskilled == 0
}

func printTargets(targets map[*group]*group) {
	for k, v := range targets {
		log.Printf("%s => %s", k.String(), v.String())
	}
}

func (a army) TeamsRemaining() int {
	teams := make(map[string]struct{})

	for _, u := range a {
		if u.Alive() {
			teams[u.team] = struct{}{}
		}
	}

	return len(teams)
}

func (a army) Result() (string, int) {
	var sum int
	teamsRemaining := make(map[string]struct{})

	for _, u := range a {
		sum += u.units
		if u.Alive() {
			teamsRemaining[u.team] = struct{}{}
		}
	}

	if len(teamsRemaining) == 1 {
		for k := range teamsRemaining {
			return k, sum
		}
	}

	return "Draw", sum
}

func (a army) String() string {
	teams := make(map[string][]*group)

	for _, u := range a {
		teams[u.team] = append(teams[u.team], u)
	}

	var out strings.Builder
	for t, groups := range teams {
		out.WriteString(fmt.Sprintf("%s\n", t))
		for _, g := range groups {
			out.WriteString(fmt.Sprintf("Group %d contains %d units\n", g.num, g.units))
		}
	}

	return out.String()
}

func (g *group) String() string {
	return fmt.Sprintf(
		"%s %d: %d units each with %d hit points (immune to %s; weak to %s) with an attack that does %d %s damage at initiative %d",
		g.team, g.num, g.units, g.hp, strings.Join(g.immunities, ", "), strings.Join(g.weaknesses, ", "), g.attack, g.attacktype, g.initiative,
	)
}
