package main

import (
	"strconv"
	"strings"
)

type price struct {
	ore      int
	clay     int
	obsidian int
}

type robotPrice struct {
	
	price    price
}

type blueprint struct{
	id int
	ore robotPrice
	clay robotPrice
	geode robotPrice
}

func parseRobots(s string) blueprint {

	blueprint := blueprint{}
	s = strings.ReplaceAll(s, "\n", " ")
	t := strings.Split(s, ":")
	id, _ := strconv.Atoi(strings.Fields(t[0])[1])




}
