package main

import (
	"fmt"

	"github.com/goravel/framework/contracts/config"
)

type Question struct {

}


func NewQuestion(question Question) *Question {
	return &Question{}
}

func (s *Hello) World() string {
	return fmt.Sprintf("Welcome To Goravel %s", s.config.GetString("hello.name"))
}

func Ask(q *Question , d String ): string
{
    $answer = readline($question.($default ? " ({$default})" : null).': ');

    if (! $answer) {
        return $default;
    }

    return $answer;
}
