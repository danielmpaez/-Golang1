package main

import (
	"flag"
	"log"
	//"os"
	//"fmt"
	//"strconv"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// "github.com/gorilla/websocket"
	// "golang.org/x/net/proxy"
	"time"
	"github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/postgres"
 "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
   }
   // BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) {
	//uuid := uuid.NewV4()
	
	
	//return scope.SetColumn("ID", uuid)
   }
   // User is the model for the user table.
   type User5 struct {
	Name   string    //`gorm:"column:name;size:128;not null;"`
	UserID uuid.UUID //`gorm:"type:uuid;not null"`
	Email   string    //`gorm:"column:email;size:128;not null;"`
	Profile  Profile
   }
   // Profile is the model for the profile table.
   type Profile struct {

	Name   string   // `gorm:"column:name;size:128;not null;"`
	UserID uuid.UUID //`gorm:"type:uuid;not null"`
   }
  
   //var profile = &Profile{Name: "New User", UserID: user.ID,}
	//	if db.Create(&profile).Error != nil {
	 //	log.Panic("Unable to create profile.")
	//}
   var fetchedUser = &User5{}
	//if db.Where("id = ?", profile.UserID).Preload("Profile").First(&fetchedUser).RecordNotFound() {
	 //log.Panic("Unable to find created user.")
	//}
  // fmt.Printf("User: %+v\n", fetchedUser)
   //}
func main() {

	log.Printf("Hola!")

	//
	//mqtt
	//
	//topic := flag.String("topic", "esp/test", "The topic name to/from which to publish/subscribe")
	broker := flag.String("broker", "tcp://broker.shiftr.io:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password := flag.String("password", "try", "The password (optional)")
	user := flag.String("user", "try", "The User (optional)")
	id := flag.String("id", "fercho-command-center", "The ClientID (optional)")
	cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	//qos := flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")
	// num := flag.Int("num", 1, "The number of messages to publish or subscribe (default 1)")
	// payload := flag.String("message", "", "The message text to publish (default empty)")
	// action := flag.String("action", "", "Action publish or subscribe (required)")
	// store := flag.String("store", ":memory:", "The Store Directory (default use memory store)")
	flag.Parse()

	opts := MQTT.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetClientID(*id)
	
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetCleanSession(*cleansess)

	// simple mqtt pub
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())}

	db, err := gorm.Open("postgres", "host=hansken.db.elephantsql.com port=5432 user=vtwlajng dbname=vtwlajng password=	Sty59HjeuNLpFjjhRA5HNok1gHc58lVs")
	if err != nil {
	 panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&User5{}, &Profile{})
	/*user := &User{SomeFlag: false}
	if db.Create(&user).Error != nil {
	 log.Panic("Unable to create user.")	
	}*/
	log.Println("Sample Publisher Started")

	bot, err := tgbotapi.NewBotAPI(("1061491150:AAHd2hlo9dPkxLajLWpsBzhCNc6XD_jg79w"))
	if err != nil {
		panic(err) // You should add better error handling than this!
	}

	bot.Debug = true // Has the library display every request and response.

	// Create a new UpdateConfig struct with an offset of 0.
	// Future requests can pass a higher offset to ensure there aren't duplicates.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we want to keep the connection open longer and wait for incoming updates.
	// This reduces the number of requests that are made while improving response time.
	updateConfig.Timeout = 60

	// Now we can pass our UpdateConfig struct to the library to start getting updates.
	// The GetUpdatesChan method is opinionated and as such, it is reasonable to implement
	// your own version of it. It is easier to use if you have no special requirements though.
	updates, err := bot.GetUpdatesChan(updateConfig)

	// Now we're ready to start going through the updates we're given.
	// Because we have a channel, we can range over it.
	for update := range updates {
		// There are many types of updates. We only care about messages right now,
		// so we should ignore any other kinds.
		if update.Message == nil {
			continue
		}

		// Sample #1
		// // Because we have to create structs for every kind of request,
		// // there's a number of helper functions to make creating common
		// // types easier. Here, we're using the NewMessage helper which
		// // returns a MessageConfig struct.
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		// // As there's too many fields for each Config to specify in a single
		// // function call, we need to modify the result the helper gave us.
		// msg.ReplyToMessageID = update.Message.MessageID

		// // We're ready to send our message!
		// // The Send method is for Configs that return a Message struct.
		// // Sending Messages (among many other types) return a Message.
		// // In this case, we don't care about the returned Message.
		// // We only need to make sure our message went through successfully.
		// if _, err := bot.Send(msg); err != nil {
		// 	panic(err) // Again, this is a bad way to handle errors.
		// }

		// var byte payload = 1
		// Sample #2
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			/*case "calentar":
				if update.Message.CommandArguments() != "" {
					var payload []byte = []byte{0}
					input := update.Message.CommandArguments()
					payload[0] = input[0] - 48
					log.Printf("Calentar el agua a %v", payload)
					token := client.Publish(*topic, byte(*qos), false, payload) // *payload)
					token.Wait()
					msg.Text = "Ahora va..."
				} else {
					msg.Text = "A cuantos grados?"
				}*/

			case "signup":
				if update.Message.CommandArguments() != "" {
					//var payload []byte = []byte{0}
					input := update.Message.CommandArguments()
					//user = User5{input}
					user := &User5{Email: input}
						if db.Create(&user).Error != nil {
						log.Panic("Unable to create user.")
						}
					 //newuser := User5{ Email:input} db.NewRecord(newuser)
					log.Printf(input)

					//func (newuser *User5) BeforeCreate(scope *gorm.Scope) 

					msg.Text = "Usuario creado"
				} else {
					msg.Text = "Ingresa tu email"
					
						}
						/*func (m *Message) Command() string {
							command := m.CommandWithAt()
						
							if i := strings.Index(command, "@"); i != -1 {
								command = command[:i]
							}
						
							return command*/
						}
				bot.Send(msg)
				}
			/*case "help":
				msg.Text = "type /sayhi or /status."
			case "sayhi":
				msg.Text = "Hi :)"			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"*/
				
			}
			//bot.Send(msg)
		}
