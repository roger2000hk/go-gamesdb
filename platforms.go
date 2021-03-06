package thegamesdb

//go:generate generategamesdb -output get_platforms_list.go -method GetPlatformsList -type GetPlatformListResponse -endpoint GetPlatformsList.php
//go:generate generategamesdb -output get_platform.go -method GetPlatform -type GetPlatformResponse -endpoint GetPlatform.php
//go:generate generategamesdb -output get_platform_games.go -method GetPlatformGames -type GetPlatformGamesResponse -endpoint GetPlatformGames.php

type GetPlatformGamesResponse struct {
	Games []GameEntity `xml:"Game"`
}

type GetPlatformListResponse struct {
	PlatformTag Platforms `xml:"Platforms"`
}

type GetPlatformResponse struct {
	Platform PlatformEntity `xml:"Platform"`
}

type Platforms struct {
	Platforms []ConcisePlatformEntity `xml:"Platform"`
}

type ConcisePlatformEntity struct {
	Id    string `xml:"id"`
	Name  string `xml:"name"`
	Alias string `xml:"alias"`
}

type PlatformEntity struct {
	Id             string `xml:"id"`
	Platform       string `xml:"Platform"`
	Console        string `xml:"console"`
	Controller     string `xml:"controller"`
	Overview       string `xml:"overview"`
	Developer      string `xml:"developer"`
	Manufacturer   string `xml:"manufacturer"`
	CPU            string `xml:"cpu"`
	Memory         string `xml:"memory"`
	Graphics       string `xml:"graphics"`
	Sound          string `xml:"sound"`
	Display        string `xml:"display"`
	Media          string `xml:"media"`
	MaxControllers string `xml:"maxcontrollers"`
	Ratings        string `xml:"Rating"`
}
