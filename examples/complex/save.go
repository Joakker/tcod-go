package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/Joakker/tcod-go"
)

type SaveData struct {
	TurnCount uint64
	Player    Player
	Enemies   []*Enemy
}

func SaveGame() {
	s := SaveData{Player: player, Enemies: EnemyList,TurnCount:TurnCount}
	data, err := json.MarshalIndent(s, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Create("save.json")
	file.Write(data)
	file.Close()
}

func LoadGame() {
	file, err := os.Open("save.json")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	file.Close()

	s := SaveData{}

	if err = json.Unmarshal(data, &s); err != nil {
		log.Fatal(err)
	}

	player = s.Player
	EnemyList = s.Enemies
	TurnCount = s.TurnCount
	DataConsole = tcod.NewConsole(12, WinH-2)
	GameConsole = tcod.NewConsole(WinW-(DataConsole.GetW()+3), WinH-2)
}
