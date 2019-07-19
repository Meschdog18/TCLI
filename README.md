# TCLI
A Twitter command line interface, interact with the twitter api via terminal

# Commands

/help - Info on all comands

/end - ends program

/user* - Enter a twitter ScreenName after typing /user* | ex: /user*joe

/tweet* - Enter your tweet after typing /tweet* | ex /tweet*Hi from my terminal

# Authentication
In the current programs state your going to have to generate a bearer token and enter it where its labed var is in the program (How Do that here - https://developer.twitter.com/en/docs/basics/authentication/guides/bearer-tokens.html ),than just put your Consumer Key ,Consumer Secret, Access Token, Token Secret in their respective labed var
# How To Use

Make sure TCLI.go and Structs.go are in the same folder than use go run *.go

# Additional Info

With the current way the program is writen, you will require credentials only available to twitter users with api access, 
so most use cases i'm assuming is just new twitter devs who just want to see an example or want to test and make sure their credentials work. 
