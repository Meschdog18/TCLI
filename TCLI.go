package main

import(
	"fmt"
	"net/http"
	"log"
	"net/url"
	"encoding/json" 
	"os"
	"bufio"
	"strings"
	"strconv"
	"github.com/dghubble/oauth1"
)

const bearertoken = "Bearer <bearertoken>"

func main(){
fmt.Println("____________________________________")
fmt.Println("|  TCLI - V.1 | By Jordan Mesches  |")
fmt.Println("====================================")
fmt.Println("Type /help for help and information")

input()
}
func input(){
	reader:=bufio.NewReader(os.Stdin)


	for{
		text,_:=reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if(strings.Contains(text,"/help")){
			help := []string{"/help - Info on all comands","/end - ends program","/user* - Enter a twitter ScreenName after typing /user* | ex: /user*joe","/tweet* - Enter your tweet after typing /tweet* | ex /tweet*Hi from my terminal"}
			fmt.Println("______")
			fmt.Println("|Help|")
			fmt.Println("------")
			for i:=0;i<len(help);i++{
				fmt.Println(help[i])
			}
		}
		if(strings.Contains(text,"/end")){
			os.Exit(3)
		}
		if(strings.Contains(text,"/user*")){
			inputstring := strings.Split(text,"*")
				
				GetTweets(inputstring[1])
		}
		if(strings.Contains(text,"/tweet*")){
			inputstring := strings.Split(text, "*")
			SendTweet(inputstring[1])
		}
	}
}
func SendTweet(message string){
	    config := oauth1.NewConfig("<consumer key>", "<consumer secret>")
    token := oauth1.NewToken("<access token>", "<token secret>")
    httpClient := config.Client(oauth1.NoContext,token)
    path := "https://api.twitter.com/1.1/statuses/update.json"+"?status="+url.QueryEscape(message)
    fmt.Println(path)
    resp, err:= httpClient.Post(path,"application/json",nil)
    defer resp.Body.Close()
    if(err != nil){
    	fmt.Println(err)
    }else{
    	fmt.Println("Sent:",message)
    }
  
    
}
func GetTweets(query string){
	endpoint,_:=url.Parse("https://api.twitter.com/1.1/statuses/user_timeline.json")
		params := url.Values{}
		params.Add("screen_name",query)
		endpoint.RawQuery = params.Encode()
		parse_endpoint := endpoint.String()
		req,err := http.NewRequest("GET", parse_endpoint,nil)
		req.Header.Add("Authorization",bearertoken)
		resp, err := http.DefaultClient.Do(req)
		if(err != nil){
			log.Fatal(err)
		}
		defer resp.Body.Close()
		var USER Data
		if err := json.NewDecoder(resp.Body).Decode(&USER); err != nil{
			log.Println(err)
		}
		//output
		

		//checks to make sure there is a user with that name
		if(len(USER)!=0){
			for i:=len(USER)-1;0<i;i--{
			
			fmt.Println(USER[i].Text)
			fmt.Println(USER[i].CreatedAt)
			}
		//nice graphics
		name:=USER[0].User.ScreenName
		followercount := strconv.FormatInt(USER[0].User.FollowersCount,10)
		output := "|"+"Username:" +  name + " | Number Of Followers:" + followercount+"|"
		for i:=0;i<len(output);i++{
			fmt.Print("_")
		}
		fmt.Println(" ")
		fmt.Println(output)
		for i:=0;i<len(output);i++{
			fmt.Print("-")
		}
		fmt.Println(" ")

		}
		
		//error message
		if(len(USER)==0){
			fmt.Println("Error - No User With That Name")
		}
		
}
