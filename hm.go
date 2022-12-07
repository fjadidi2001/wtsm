package main

import (
	"errors"
	"fmt"
	"image"
	"os"
)

img,err:=openImage("/home/jadidi/Desktop/Arcface/player1.jpg")
if err!=nil{
return
}
func openImage(path string)(image.Image,error){
	f, err := os.Open(path)
	if err !=nil{
		fmt.Println(err)
		return nil,err
	}
	fi,_:=f.Stat()
	fmt.Println(fi.Name())
	//defer f.Close()sss
	img,format,err:=image.Decode(f)
	if err!=nil{
		fmt.Println("Decoding error:",err.Error())
		return nil,err
	}
	if format != "jpeg"{
		fmt.Println("image format is not jpeg")
		return nil,errors.New("")
	}
	return img,nil
}