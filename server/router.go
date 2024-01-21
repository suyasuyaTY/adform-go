package server

import "fmt"

func Init() error {
	r, err := NewRouter()
	if err != nil {
		return err
	} 
	r.Run(fmt.Sprintf(":%d", 8080))
	return nil
}