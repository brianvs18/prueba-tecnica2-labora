package card

import (
	"ejercicio-poo/challenge/yu-gi-oh/common"
	"github.com/olekukonko/tablewriter"
	"math/rand"
	"os"
	"strconv"
)

const (
	url = "https://db.ygoprodeck.com/api/v7/cardinfo.php?&startdate=2000-01-01&enddate=2002-08-23&dateregion=tcg"
)

func getAll() ([]Card, error) {
	if len(storage) == 0 {
		cards, err := FetchYuGiOhCards(url)
		if err != nil {
			return nil, err
		}
		storage = cards
	}
	tempCards := make([]Card, len(storage))
	copy(tempCards, storage)
	rand.Shuffle(len(tempCards), func(i, j int) {
		tempCards[i], tempCards[j] = tempCards[j], tempCards[i]
	})
	return tempCards[:15], nil
}

func DisplayAll() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Type", "FrameType", "Description", "Atk", "Def", "Level", "Race", "Attribute"})
	cards, err := getAll()
	if err != nil {
		panic(err)
	}
	for _, card := range cards {
		cardsData := []string{
			strconv.Itoa(card.Id),
			card.Name,
			card.Type,
			card.FrameType,
			card.Description,
			strconv.Itoa(card.Atk),
			strconv.Itoa(card.Def),
			strconv.Itoa(card.Level),
			card.Race,
			card.Attribute,
		}
		table.Append(cardsData)
	}
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{
			tablewriter.Bold, tablewriter.FgHiYellowColor})
	table.SetRowLine(true)
	table.Render()
}

func GetRandomCard() *Card {
	cards, err := getAll()
	if err != nil {
		panic(err)
	}
	for {
		selectedCard := common.GetRandomObject(cards)
		if selectedCard.FrameType != "spell" && selectedCard.FrameType != "trap" {
			return &selectedCard
		}
	}
}
