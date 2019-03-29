all: help


##                   _                   _
##  __ _  ___  _ __ | |__   ___ _ __ ___(_)___  ___  ___
## / _` |/ _ \| '_ \| '_ \ / _ \ '__/ __| / __|/ _ \/ __|
##| (_| | (_) | |_) | | | |  __/ | | (__| \__ \  __/\__ \
## \__, |\___/| .__/|_| |_|\___|_|  \___|_|___/\___||___/
## |___/      |_|
##

##  help:			Help
.PHONY : help
help : Makefile
	@sed -n 's/^##//p' $<

##  quiz-game:		Run Quiz Game (go run 1-quiz-game/part-2/main.go)
.PHONY : quiz-game
quiz-game:
	go run 1-quiz-game/part-2/main.go