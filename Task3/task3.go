package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

//global values for me to see tham if all of functions and changing it if needs
var fileText string
var questions []string
var sentence []string
var exclamation []string
var wordsArray []string
var newQuestionText []string
var newExclamationText []string

func main() {
	openTextFile()
	arrayMaker(fileText)
	replacingWords()
	cratingFile()
}

//here we open a file in falder and saving the text in "fileXext" string
func openTextFile() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("file is not finded")
		return
	}
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("can not a file shape")
		return
	}
	defer file.Close()
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("can not to read the file")
		return
	}
	fileText = string(bs)
}

//here we are cutting a string wich is given in senten, qustions and words
func arrayMaker(s string) {
	var r = regexp.MustCompile("[.?!]")
	var rWords = regexp.MustCompile("[.?!, ]")
	Strings := r.Split(s, -1)
	StringsArg := r.FindAllString(s, -1)
	sentensMaker(Strings, StringsArg)
	wordsSplited := rWords.Split(s, -1)
	for _, value := range wordsSplited {
		wordsArray = append(wordsArray, value)
	}
}

//making a new array of main atribute
func sentensMaker(text []string, simbol []string) {

	for i := 0; i < len(simbol); i++ {
		if simbol[i] == "." {
			sentence = append(sentence, text[i]+simbol[i])
		}
		if simbol[i] == "?" {
			questions = append(questions, text[i]+simbol[i])
		}
		if simbol[i] == "!" {
			exclamation = append(exclamation, text[i]+simbol[i])
		}
	}
}

// creating a file with and writing here an array of things
func cratingFile() {
	file, err := os.Create("newFile.txt")
	if err != nil {
		fmt.Println("can not to creat a file")
		return
	}
	defer file.Close()
	file.WriteString("Sentences\n")
	for _, value := range sentence {
		file.Write([]byte(value))
		file.WriteString("\n")
	}
	file.WriteString("Questions\n")
	for _, value := range questions {
		file.Write([]byte(value))
		file.WriteString("\n")
	}
	file.WriteString("Exclamations\n")
	for _, value := range exclamation {
		file.Write([]byte(value))
		file.WriteString("\n")
	}
	file.WriteString("NewExclamations\n")
	for _, value := range newExclamationText {
		file.Write([]byte(value))
		file.WriteString("\n")
	}
	file.WriteString("NewQuestions\n")
	file.WriteString("I did not undestand the task for this part \n")
	for _, value := range newQuestionText {
		file.Write([]byte(value))
		file.WriteString("\n")
	}
	file.WriteString("All words in use\n")
	for _, value := range wordsArray {
		file.Write([]byte(value))
		file.WriteString("\n")
	}
}

//a logic of task3 which is given

func replacingWords() {

	var rege = regexp.MustCompile("[ !?.]")
	var first string
	var last string
	for _, value := range exclamation {
		words := rege.Split(value, -1)
		first = words[1]
		last = words[len(words)-2]
		words[1] = last
		words[len(words)-2] = first
		wordsStr := strings.Join(words[0:len(words)-1], " ")
		wordsStr = wordsStr + "!"
		newExclamationText = append(newExclamationText, wordsStr)
	}
	for _, value := range questions {
		words := rege.Split(value, -1)
		first = words[1]
		last = words[len(words)-2]
		wordsStr := first + " " + last
		wordsStr = wordsStr + "?"
		newQuestionText = append(newQuestionText, wordsStr)
	}
}
