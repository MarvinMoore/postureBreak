package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"strconv"
	"github.com/vigneshuvi/GoDateFormat"

	"github.com/bwmarrin/discordgo"
)
var startTime int
// Variables used for command line parameters
var (
	Token string
)
func timerStart(s *discordgo.Session, m *discordgo.MessageCreate) int {
	//1 hour ticker
    
    go func() {
    ticker := time.NewTicker(1 * time.Minute)
	for _ = range ticker.C {
	fmt.Println("Smoke Break")
	s.ChannelMessageSend(m.ChannelID, "Posture Check!")

	}
    }()
 	s.ChannelMessageSend(m.ChannelID, "Starting the timer!")
	time := time.Now()
	var timeFormatted string
	timeFormatted = time.Format(GoDateFormat.ConvertFormat("HH:MM:SS"))
	//text = strings.ReplaceAll(timeFormatted, ":", "")
	//timeStarted, err := strconv.Atoi(text)
	runes := []rune(timeFormatted)
    // ... Convert back into a string from rune slice.
    safeSubstring := string(runes[3:5])
	startHour, err := strconv.Atoi(safeSubstring)
	if(err!=nil){
	fmt.Println(err)
	}
	fmt.Println(safeSubstring + " Start Hour")
	return startHour
}


func timeLeft(startHour int,s *discordgo.Session, m *discordgo.MessageCreate){
	var timeLeft int
	var currHour int
	var timeFormatted string
	time := time.Now()
	timeFormatted =time.Format(GoDateFormat.ConvertFormat("HH:MM:SS"))
	runes := []rune(timeFormatted)
	safeSubstring := string(runes[3:5])
	timeLeft=startHour-currHour
	currHour, err := strconv.Atoi(safeSubstring)
	fmt.Println("Start Hour:",startHour)
	fmt.Println("Current Hour:",currHour)
	timeLeft=startHour-currHour
	fmt.Println("Subtraction Answer :",timeLeft)
	if(err!=nil){
		fmt.Println(err)
	}
	if(timeLeft<0){
	//negative
	//fmt.Println("Negative timeleft %s", timeLeft)
	timeLeft=timeLeft+60
	message:= " Minutes left!"
	concatenated:= strconv.Itoa(timeLeft) + message
    s.ChannelMessageSend(m.ChannelID, concatenated)

	}else{
	s.ChannelMessageSend(m.ChannelID, "60 Minutes Left!")
	}


}



func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	
	if m.Content == "!start" {
		
		startTime=timerStart(s, m)
		fmt.Println(startTime)
		
	}

	
	if m.Content == "!time" {
		timeLeft(startTime,s,m)	
	}
}