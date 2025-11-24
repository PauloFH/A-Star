package data

import (
	"sort"
)

var Graph map[string][]Edge
var Heuristics map[string]int
var Positions map[string]struct{ X, Y float32 }
var SortedCities []string

func InitData() {
	Graph = map[string][]Edge{
		"Arad":           {{"Zerind", 75}, {"Sibiu", 140}, {"Timisoara", 118}},
		"Zerind":         {{"Arad", 75}, {"Oradea", 71}},
		"Oradea":         {{"Zerind", 71}, {"Sibiu", 151}},
		"Sibiu":          {{"Arad", 140}, {"Oradea", 151}, {"Fagaras", 99}, {"Rimnicu Vilcea", 80}},
		"Timisoara":      {{"Arad", 118}, {"Lugoj", 111}},
		"Lugoj":          {{"Timisoara", 111}, {"Mehadia", 70}},
		"Mehadia":        {{"Lugoj", 70}, {"Drobeta", 75}},
		"Drobeta":        {{"Mehadia", 75}, {"Craiova", 120}},
		"Craiova":        {{"Drobeta", 120}, {"Rimnicu Vilcea", 146}, {"Pitesti", 138}},
		"Rimnicu Vilcea": {{"Sibiu", 80}, {"Craiova", 146}, {"Pitesti", 97}},
		"Fagaras":        {{"Sibiu", 99}, {"Bucharest", 211}},
		"Pitesti":        {{"Rimnicu Vilcea", 97}, {"Craiova", 138}, {"Bucharest", 101}},
		"Bucharest":      {{"Fagaras", 211}, {"Pitesti", 101}, {"Giurgiu", 90}, {"Urziceni", 85}},
		"Giurgiu":        {{"Bucharest", 90}},
		"Urziceni":       {{"Bucharest", 85}, {"Hirsova", 98}, {"Vaslui", 142}},
		"Hirsova":        {{"Urziceni", 98}, {"Eforie", 86}},
		"Eforie":         {{"Hirsova", 86}},
		"Vaslui":         {{"Urziceni", 142}, {"Iasi", 92}},
		"Iasi":           {{"Vaslui", 92}, {"Neamt", 87}},
		"Neamt":          {{"Iasi", 87}},
	}

	Heuristics = map[string]int{
		"Arad": 366, "Bucharest": 0, "Craiova": 160, "Drobeta": 242,
		"Eforie": 161, "Fagaras": 176, "Giurgiu": 77, "Hirsova": 151,
		"Iasi": 226, "Lugoj": 244, "Mehadia": 241, "Neamt": 234,
		"Oradea": 380, "Pitesti": 100, "Rimnicu Vilcea": 193, "Sibiu": 253,
		"Timisoara": 329, "Urziceni": 80, "Vaslui": 199, "Zerind": 374,
	}

	Positions = map[string]struct{ X, Y float32 }{
		"Arad": {60, 150}, "Zerind": {90, 70}, "Oradea": {180, 50},
		"Timisoara": {70, 280}, "Lugoj": {180, 330}, "Mehadia": {185, 390},
		"Drobeta": {180, 470}, "Craiova": {300, 500}, "Sibiu": {250, 200},
		"Rimnicu Vilcea": {330, 250}, "Fagaras": {400, 210}, "Pitesti": {450, 350},
		"Bucharest": {580, 420}, "Giurgiu": {540, 520}, "Urziceni": {650, 380},
		"Neamt": {560, 80}, "Iasi": {630, 120}, "Vaslui": {700, 200},
		"Hirsova": {720, 380}, "Eforie": {750, 480},
	}

	SortedCities = make([]string, 0, len(Positions))
	for k := range Positions {
		SortedCities = append(SortedCities, k)
	}
	sort.Strings(SortedCities)
}
